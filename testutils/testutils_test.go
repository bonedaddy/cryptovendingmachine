package testutils

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	bindingsm "github.com/postables/cryptovendingmachine/bindings/vendingMachine"
	bindingsf "github.com/postables/cryptovendingmachine/bindings/vendorFactory"
	bindingsvm "github.com/postables/cryptovendingmachine/bindings/vendorManagement"
)

var (
	// blake2b-264 with cidv1
	hash1 = "bafy2dzaceeylnokurfh2lhglows4kyrhua4iw2wb7h423qc66vozap26pivgr7i"
	// bke2b-264 with cidv1
	hash2 = "bafy2dzaceegxix47qjggdvh66kduvov677rncikxbgukyk47rngou733miejgmq"
)

func getHash(t *testing.T, hash string) [2][32]byte {
	cid, err := FormatCID(hash)
	if err != nil {
		t.Fatal(err)
	}
	return cid
}

func Test_SumKeccak256(t *testing.T) {
	hash := SumKeccak256([]byte(hash1))
	if len(hash) != 32 {
		t.Fatal("bad hash")
	}
}

func Test_FormatCID(t *testing.T) {
	cid := getHash(t, hash1)
	fmt.Println(len(hash1))
	if len(cid) != 2 {
		t.Fatal("bad formatted hash")
	}
}

func Test_NewBlockchain(t *testing.T) {
	_, _ = NewBlockchain(t)
}

func Test_DeployVendorManagement(t *testing.T) {
	auth, sim := NewBlockchain(t)
	_, addr := DeployVendorManagement(t, sim, auth)
	if _, err := bindingsvm.NewVendormanagement(addr, sim); err != nil {
		t.Fatal(err)
	}
}

func Test_VendorManagement(t *testing.T) {
	auth, sim := NewBlockchain(t)
	contract, _ := DeployVendorManagement(t, sim, auth)
	retID, err := contract.Id(nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(retID, SumKeccak256(auth.From.Bytes())) {
		t.Fatal("bad id")
	}
	RegisterProduct(t, sim, auth, contract)
	isSold, err := contract.SoldAt(
		nil,
		"lays chip",
		"1",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isSold {
		t.Fatal("product should be sold at location")
	}
	isSold, err = contract.SoldAt(
		nil,
		"lays chip",
		"da hood",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isSold {
		t.Fatal("product should be sold at location")
	}
	isSold, err = contract.SoldAt(
		nil,
		"lays chip",
		"3",
	)
	if err != nil {
		t.Fatal(err)
	}
	if isSold {
		t.Fatal("product should not be sold at location")
	}
	AddProductLocation(t, sim, auth, contract)
	isSold, err = contract.SoldAt(
		nil,
		"lays chip",
		"3",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isSold {
		t.Fatal("product should be sold at location")
	}
	RemoveProductLocation(t, sim, auth, contract)
	isSold, err = contract.SoldAt(
		nil,
		"lays chip",
		"3",
	)
	if err != nil {
		t.Fatal(err)
	}
	if isSold {
		t.Fatal("product should not be sold at location")
	}
	isSold, err = contract.SoldAt(
		nil,
		"lays chip",
		"da hood",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isSold {
		t.Fatal("product should be sold at location")
	}
}

func Test_VendorFactory(t *testing.T) {
	auth, sim := NewBlockchain(t)
	contract, addr := DeployVendorFactory(t, sim, auth)
	if _, err := bindingsf.NewVendorfactory(addr, sim); err != nil {
		t.Fatal(err)
	}
	// deploy the vendor management contract
	if err := NewVendor(t, sim, auth, contract); err != nil {
		t.Fatal(err)
	}
	// test another deploy, this should fail
	if err := NewVendor(t, sim, auth, contract); err == nil {
		t.Fatal("error expected")
	}
	// test that they're registered
	isRegistered, err := contract.RegisteredVendors(nil, auth.From)
	if err != nil {
		t.Fatal(err)
	}
	if !isRegistered {
		t.Fatal("vendor should be registered")
	}
	// get the management contract address from the vendor address
	managementAddr, err := contract.VendorContract(nil, auth.From)
	if err != nil {
		t.Fatal(err)
	}
	// construct management contract wrapping
	management, err := bindingsvm.NewVendormanagement(managementAddr, sim)
	if err != nil {
		t.Fatal(err)
	}
	// conduct all vendor management tests
	retID, err := management.Id(nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(retID, SumKeccak256(auth.From.Bytes())) {
		t.Fatal("bad id")
	}
	RegisterProduct(t, sim, auth, management)
	isSold, err := management.SoldAt(
		nil,
		"lays chip",
		"1",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isSold {
		t.Fatal("product should be sold at location")
	}
	isSold, err = management.SoldAt(
		nil,
		"lays chip",
		"da hood",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isSold {
		t.Fatal("product should be sold at location")
	}
	isSold, err = management.SoldAt(
		nil,
		"lays chip",
		"3",
	)
	if err != nil {
		t.Fatal(err)
	}
	if isSold {
		t.Fatal("product should not be sold at location")
	}
	AddProductLocation(t, sim, auth, management)
	isSold, err = management.SoldAt(
		nil,
		"lays chip",
		"3",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isSold {
		t.Fatal("product should be sold at location")
	}
	RemoveProductLocation(t, sim, auth, management)
	isSold, err = management.SoldAt(
		nil,
		"lays chip",
		"3",
	)
	if err != nil {
		t.Fatal(err)
	}
	if isSold {
		t.Fatal("product should not be sold at location")
	}
	isSold, err = management.SoldAt(
		nil,
		"lays chip",
		"da hood",
	)
	if err != nil {
		t.Fatal(err)
	}
	if !isSold {
		t.Fatal("product should be sold at location")
	}
}

func Test_VendingMachine(t *testing.T) {
	auth, sim := NewBlockchain(t)
	manageContract, managementAddr := DeployVendorManagement(t, sim, auth)
	contract, addr := DeployVendingMachine(t, sim, auth)
	if _, err := bindingsm.NewVendingmachine(addr, sim); err != nil {
		t.Fatal(err)
	}
	location, err := contract.Location(nil)
	if err != nil {
		t.Fatal(err)
	}
	if location != "da hood" {
		t.Fatal("bad location recovered")
	}
	backend, err := contract.Backend(nil)
	if err != nil {
		t.Fatal(err)
	}
	if backend != auth.From {
		t.Fatal("bad backend address recovered")
	}
	// test registration
	if err := AddVendor(t, sim, auth, contract, managementAddr); err != nil {
		t.Fatal(err)
	}
	// test duplicate registration
	if err := AddVendor(t, sim, auth, contract, managementAddr); err == nil {
		t.Fatal("error expected")
	}
	// test vendor registration
	isRegistered, err := contract.Vendors(nil, auth.From)
	if err != nil {
		t.Fatal(err)
	}
	if !isRegistered {
		t.Fatal("vendor should be registered")
	}
	retVendorContract, err := contract.VendorContracts(nil, auth.From)
	if err != nil {
		t.Fatal(err)
	}
	if retVendorContract != managementAddr {
		t.Fatal("bad vendor contract returned")
	}
	// test vendor name to vendor contract lookup
	retManageContract, err := contract.VendorNames(nil, "lays")
	if err != nil {
		t.Fatal(err)
	}
	if retManageContract != managementAddr {
		t.Fatal("bad vendor contract returned")
	}
	// TODO(postables): add purchase tests

	// test a bad withdraw because there are no funds
	if _, err = manageContract.WithdrawFunds(auth); err == nil {
		t.Fatal("error expected")
	}
	// register products for sale
	RegisterProduct(t, sim, auth, manageContract)
	if err := PurchaseProduct(t, sim, auth, contract); err != nil {
		t.Fatal(err)
	}
	tx, err := manageContract.WithdrawFunds(auth)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	if _, err := bind.WaitMined(context.Background(), sim, tx); err != nil {
		t.Fatal(err)
	}
}
