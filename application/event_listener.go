package application

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/go-hclog"
)

const personInfoAbi = `[
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "uint256",
                "name": "personIndex",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "newName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "newAge",
                "type": "uint256"
            }
        ],
        "name": "PersonInfoUpdated",
        "type": "event"
    }
]`

type EventListener struct {
	l  hclog.Logger
	cr repo.ContractInterface
}

func (el *EventListener) PersonInfoEventListenerStart(ctx context.Context, client *ethclient.Client) {
	contractAbi, err := abi.JSON(strings.NewReader(personInfoAbi))
	if err != nil {
		fmt.Printf("failed to parse ABI: %v", err)
		panic(err)
	}
	personInfoTopic := crypto.Keccak256Hash([]byte("PersonInfoUpdated(uint256,string,uint256)"))
	log.Println("Subscribed to PersonInfoUpdated events")

	query := ethereum.FilterQuery{
		FromBlock: nil, // Start from latest block
		ToBlock:   nil, // Monitor new blocks
		Addresses: []common.Address{common.HexToAddress("0xCc71C11BfaEC766C86EaFBb2F2F84A4160FdE1b2")},
		Topics:    [][]common.Hash{{personInfoTopic}},
	}

	ch := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)
	if err != nil {
		el.l.Error("cannot subscribe to event", "error", err)
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case log := <-ch:
			el.l.Info("matching log encountered", "log", log)

			var eventData struct {
				NewName     string
				NewAge      *big.Int
				PersonIndex *big.Int
			}

			for _, evInfo := range contractAbi.Events {
				if evInfo.ID.Hex() != log.Topics[0].Hex() {
					continue
				}

				indexed := make([]abi.Argument, 0)
				for _, input := range evInfo.Inputs {
					if input.Indexed {
						indexed = append(indexed, input)
					}
				}

				// parse topics without event name
				if err := abi.ParseTopics(&eventData, indexed, log.Topics[1:]); err != nil {
					el.l.Error("cannot parse indexed topic into event data", "error", err)
					continue
				}

				if err := contractAbi.UnpackIntoInterface(&eventData, evInfo.Name, log.Data); err != nil {
					el.l.Error("cannot parse log data into event data", "error", err)
					continue
				}
				break
			}

			err = el.cr.SavePersonInfoUpdatedEvent(ctx, &model.PersonInfoEvent{
				TxHash: log.TxHash.String(),
				Index:  eventData.PersonIndex.Uint64(),
				Name:   eventData.NewName,
				Age:    eventData.NewAge.Uint64(),
			})
			if err != nil {
				el.l.Error("failed to save event data to db", "error", err)
				continue
			}
			el.l.Info("saved event data to db", "event", eventData)
		case <-ctx.Done():
			sub.Unsubscribe()
			return
		}
	}
}
