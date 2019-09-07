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

	bindingsm "github.com/postables/cryptovendingmachine/bindings/vendingMachine"
	bindingsf "github.com/postables/cryptovendingmachine/bindings/vendorFactory"
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

// RegisterProduct is used to register a product
func RegisterProduct(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
	contract *bindingsvm.Vendormanagement,
) {
	tx, err := contract.RegisterProduct(
		auth,
		"lays chip",
		[]string{"1", "da hood"},
		oneEthInWei,
	)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
}

// AddProductLocation is used to add a product location
func AddProductLocation(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
	contract *bindingsvm.Vendormanagement,
) {
	tx, err := contract.AddProductLocation(
		auth,
		"lays chip",
		"3",
	)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
}

// RemoveProductLocation is used to add a product location
func RemoveProductLocation(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
	contract *bindingsvm.Vendormanagement,
) {
	tx, err := contract.RemoveProductLocation(
		auth,
		"lays chip",
		"3",
	)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
}

// DeployVendorFactory is used to deploy the vendor factory
func DeployVendorFactory(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
) (*bindingsf.Vendorfactory, common.Address) {
	addr, tx, contract, err := bindingsf.DeployVendorfactory(auth, sim)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
	printGasUsed(t, sim, tx.Hash(), "vendor-factory")
	return contract, addr
}

// DeployVendingMachine is used to deploy the vending machine contract
func DeployVendingMachine(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
) (*bindingsm.Vendingmachine, common.Address) {
	addr, tx, contract, err := bindingsm.DeployVendingmachine(auth, sim, "da hood", auth.From)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
	printGasUsed(t, sim, tx.Hash(), "vending-machine")
	return contract, addr
}

// NewVendor is used to deploy a vendor management contract through the factory
func NewVendor(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
	contract *bindingsf.Vendorfactory,
) error {
	tx, err := contract.NewVendor(auth)
	if err != nil {
		return err
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		return err
	}
	return nil
}

// AddVendor is used to add a vendor to a vending machine
func AddVendor(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
	contract *bindingsm.Vendingmachine,
	contractAddress common.Address,
) error {
	tx, err := contract.AddVendor(auth, "lays", contractAddress)
	if err != nil {
		return err
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		return err
	}
	return nil
}

// PurchaseProduct is used to purchase a product from the vending machine
func PurchaseProduct(
	t *testing.T,
	sim *backends.SimulatedBackend,
	auth *bind.TransactOpts,
	contract *bindingsm.Vendingmachine,
) error {
	auth.Value = oneEthInWei
	// defer reset of transaction value
	defer func() {
		auth.Value = nil
	}()
	tx, err := contract.PurchaseProduct(auth, "lays", "lays chip")
	if err != nil {
		return err
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		return err
	}
	return nil
}
