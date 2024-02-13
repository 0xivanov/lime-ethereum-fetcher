package utils

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
)

func DecodeRlphex(rlphex string) ([]string, error) {
	// decode the input hex string
	bytes, err := hex.DecodeString(rlphex)
	if err != nil {
		return nil, err
	}

	// decode the RLP-encoded list
	var list []string
	if err := rlp.DecodeBytes(bytes, &list); err != nil {
		return nil, err
	}

	return list, nil
}
