package testutils

import (
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

// SumKeccak256 is used to hash input data against
// keccak256 and turn it into a [32]byte suitable
// for gas efficient smart contract storage
func SumKeccak256(data []byte) [32]byte {
	hashedData := crypto.Keccak256Hash(data)
	var hash [32]byte
	for i := 0; i < len(hashedData.Bytes()); i++ {
		hash[i] = hashedData.Bytes()[i]
	}
	return hash
}

// FormatCID is used to format a CID to be
// store on a smart contract in a manner that
// allows for taking up exactly 2 slots of
// smart contract stroage on ethereum
func FormatCID(cid string) ([2][32]byte, error) {
	if len(cid) != 64 {
		return [2][32]byte{}, errors.New("invalid cid length")
	}
	var fcid [2][32]byte
	for i := 0; i < 32; i++ {
		fcid[0][i] = cid[i]
	}
	for i := 0; i < 32; i++ {
		var ii = i + 32
		fcid[1][i] = cid[ii]
	}
	return fcid, nil
}
