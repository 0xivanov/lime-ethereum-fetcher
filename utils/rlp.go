package utils

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
)

func DecodeRlphex(rlphex string) ([]string, error) {
	// Decode the input hex string
	bytes, err := hex.DecodeString(rlphex)
	if err != nil {
		return nil, err
	}

	// Decode the RLP-encoded list
	var list []string
	if err := rlp.DecodeBytes(bytes, &list); err != nil {
		return nil, err
	}

	return list, nil
}
