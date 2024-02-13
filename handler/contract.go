package handler

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/http"
	"os"

	"github.com/0xivanov/lime-ethereum-fetcher-go/api"
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type SmartContract struct {
	l      hclog.Logger
	client *ethclient.Client
}

func NewSmartContract(l hclog.Logger, client *ethclient.Client) *SmartContract {
	return &SmartContract{l, client}
}

func (sc *SmartContract) SavePerson(ctx *gin.Context) {
	var personInfo model.PersonInfo
	if err := ctx.BindJSON(&personInfo); err != nil {
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to instantiate smart contract instance"})
		return
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	auth := getAccountAuth(sc.client, privateKey)
	tx, err := contract.SetPersonInfo(auth, personInfo.Name, big.NewInt(personInfo.Age))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to execute transaction"})
		return
	}

	// Return transaction hash and status
	ctx.JSON(http.StatusOK, gin.H{"txHash": tx.Hash().Hex(), "status": "pending"})
}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
