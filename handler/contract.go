package handler

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/http"
	"os"

	"github.com/0xivanov/lime-ethereum-fetcher-go/api"
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type SmartContract struct {
	logger            hclog.Logger
	client            *ethclient.Client
	contractInterface repo.ContractInterface
}

func NewSmartContract(logger hclog.Logger, client *ethclient.Client, contractIngerface repo.ContractInterface) *SmartContract {
	return &SmartContract{logger, client, contractIngerface}
}

func (sc *SmartContract) SavePerson(ctx *gin.Context) {
	var personInfo model.PersonInfo
	if err := ctx.BindJSON(&personInfo); err != nil {
		sc.logger.Error("invalid input", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Call the setPersonInfo function on the smart contract
	contractAddressStr := os.Getenv("CONTRACT_ADDRESS")
	if contractAddressStr == "" {
		contractAddressStr = "0xCc71C11BfaEC766C86EaFBb2F2F84A4160FdE1b2"
	}
	contractAddress := common.HexToAddress(contractAddressStr)
	contract, err := api.NewApi(contractAddress, sc.client)
	if err != nil {
		sc.logger.Error("failed to instantiate smart contract instance", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to instantiate smart contract instance"})
		return
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	auth, err := getAccountAuth(sc.client, privateKey)
	if err != nil {
		sc.logger.Error("invalid private key", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid private key"})
		return
	}

	tx, err := contract.SetPersonInfo(auth, personInfo.Name, big.NewInt(personInfo.Age))
	if err != nil {
		sc.logger.Error("failed to execute transaction", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to execute transaction"})
		return
	}

	// Return transaction hash and status
	ctx.JSON(http.StatusOK, gin.H{"txHash": tx.Hash().Hex(), "status": "pending"})
}

func (sc *SmartContract) GetPersons(ctx *gin.Context) {
	persons, err := sc.contractInterface.GetPersons(ctx)
	if err != nil {
		sc.logger.Error("failed to retrieve persons from db", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve persons from database"})
		return
	}

	// Return list of persons as JSON response
	ctx.JSON(http.StatusOK, persons)
}

func getAccountAuth(client *ethclient.Client, accountAddress string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid private key")
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("invalid private key")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("invalid private key")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth, nil
}
