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

	bindingsvm "github.com/postables/cryptovendingmachine/bindings/vendorManagement"
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
}

// DeployVendorManagement is used to deploy the vendor management contract
func DeployVendorManagement(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
) (*bindingsvm.Vendormanagement, common.Address) {
	addr, tx, contract, err := bindingsvm.DeployVendormanagement(auth, sim)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
	printGasUsed(t, sim, tx.Hash(), "vendor-management")
	return contract, addr
}

// RegisterVendor is used to register a vendor
func RegisterVendor(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
	contract *bindingsvm.Vendormanagement,
) {
	tx, err := contract.RegisterVendor(auth)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
}

// RegisterProduct is used to register a product
func RegisterProduct(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
	contract *bindingsvm.Vendormanagement,
) {
	tx, err := contract.RegisterProduct(
		auth,
		SumKeccak256([]byte("lays chip")),
		[][32]byte{SumKeccak256([]byte("1")), SumKeccak256([]byte("2"))},
	)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
}
