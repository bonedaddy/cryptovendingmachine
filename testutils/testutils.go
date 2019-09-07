package testutils

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"

	"io/ioutil"

	"os"
)

var (
	oneEthInWei = big.NewInt(1000000000000000000)
)

// NewBlockchain is used to generate a simulated blockchain
func NewBlockchain(t *testing.T) (*bind.TransactOpts, *backends.SimulatedBackend) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(key)
	// https://medium.com/coinmonks/unit-testing-solidity-contracts-on-ethereum-with-go-3cc924091281
	balance := new(big.Int).Mul(big.NewInt(999999999999999), big.NewInt(999999999999999))
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: balance},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 8000000)
	return auth, sim
}

func printGasUsed(t *testing.T, sim *backends.SimulatedBackend, hash common.Hash, contract string) {
	rcpt, err := sim.TransactionReceipt(context.Background(), hash)
	if err != nil {
		t.Fatal(err)
	}
	msg := fmt.Sprint("gas used: ", rcpt.CumulativeGasUsed)
	fmt.Println(msg)
	if err := ioutil.WriteFile(fmt.Sprintf("%s.gas.log", contract), []byte(msg), os.FileMode(0640)); err != nil {
		t.Fatal(err)
	}
}
