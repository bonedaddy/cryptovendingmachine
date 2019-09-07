// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vendingmachine

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VendingmachineABI is the input ABI used to generate the binding from.
const VendingmachineABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vendors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vendorContract\",\"type\":\"address\"}],\"name\":\"addVendor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VendingmachineBin is the compiled bytecode used for deploying new contracts.
var VendingmachineBin = "0x608060405234801561001057600080fd5b506101d9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806327e30c361461003b57806345e0f75314610075575b600080fd5b6100616004803603602081101561005157600080fd5b50356001600160a01b031661009b565b604080519115158252519081900360200190f35b6100616004803603602081101561008b57600080fd5b50356001600160a01b03166100b0565b60006020819052908152604090205460ff1681565b600080829050336001600160a01b0316816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156100f957600080fd5b505afa15801561010d573d6000803e3d6000fd5b505050506040513d602081101561012357600080fd5b50516001600160a01b031614610180576040805162461bcd60e51b815260206004820152601d60248201527f63616c6c6572206d7573742062652076656e646f72206d616e61676572000000604482015290519081900360640190fd5b5050336000908152602081905260409020805460ff1916600190811790915591905056fea265627a7a723158209ef08f15300f2da2b755db727e6244d1211fc70d742d500a155dd643a68fa89864736f6c634300050b0032"

// DeployVendingmachine deploys a new Ethereum contract, binding an instance of Vendingmachine to it.
func DeployVendingmachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vendingmachine, error) {
	parsed, err := abi.JSON(strings.NewReader(VendingmachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VendingmachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vendingmachine{VendingmachineCaller: VendingmachineCaller{contract: contract}, VendingmachineTransactor: VendingmachineTransactor{contract: contract}, VendingmachineFilterer: VendingmachineFilterer{contract: contract}}, nil
}

// Vendingmachine is an auto generated Go binding around an Ethereum contract.
type Vendingmachine struct {
	VendingmachineCaller     // Read-only binding to the contract
	VendingmachineTransactor // Write-only binding to the contract
	VendingmachineFilterer   // Log filterer for contract events
}

// VendingmachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type VendingmachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendingmachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VendingmachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendingmachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VendingmachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendingmachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VendingmachineSession struct {
	Contract     *Vendingmachine   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VendingmachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VendingmachineCallerSession struct {
	Contract *VendingmachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VendingmachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VendingmachineTransactorSession struct {
	Contract     *VendingmachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VendingmachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type VendingmachineRaw struct {
	Contract *Vendingmachine // Generic contract binding to access the raw methods on
}

// VendingmachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VendingmachineCallerRaw struct {
	Contract *VendingmachineCaller // Generic read-only contract binding to access the raw methods on
}

// VendingmachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VendingmachineTransactorRaw struct {
	Contract *VendingmachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVendingmachine creates a new instance of Vendingmachine, bound to a specific deployed contract.
func NewVendingmachine(address common.Address, backend bind.ContractBackend) (*Vendingmachine, error) {
	contract, err := bindVendingmachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vendingmachine{VendingmachineCaller: VendingmachineCaller{contract: contract}, VendingmachineTransactor: VendingmachineTransactor{contract: contract}, VendingmachineFilterer: VendingmachineFilterer{contract: contract}}, nil
}

// NewVendingmachineCaller creates a new read-only instance of Vendingmachine, bound to a specific deployed contract.
func NewVendingmachineCaller(address common.Address, caller bind.ContractCaller) (*VendingmachineCaller, error) {
	contract, err := bindVendingmachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VendingmachineCaller{contract: contract}, nil
}

// NewVendingmachineTransactor creates a new write-only instance of Vendingmachine, bound to a specific deployed contract.
func NewVendingmachineTransactor(address common.Address, transactor bind.ContractTransactor) (*VendingmachineTransactor, error) {
	contract, err := bindVendingmachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VendingmachineTransactor{contract: contract}, nil
}

// NewVendingmachineFilterer creates a new log filterer instance of Vendingmachine, bound to a specific deployed contract.
func NewVendingmachineFilterer(address common.Address, filterer bind.ContractFilterer) (*VendingmachineFilterer, error) {
	contract, err := bindVendingmachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VendingmachineFilterer{contract: contract}, nil
}

// bindVendingmachine binds a generic wrapper to an already deployed contract.
func bindVendingmachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VendingmachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendingmachine *VendingmachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendingmachine.Contract.VendingmachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendingmachine *VendingmachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendingmachine.Contract.VendingmachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendingmachine *VendingmachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendingmachine.Contract.VendingmachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendingmachine *VendingmachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendingmachine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendingmachine *VendingmachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendingmachine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendingmachine *VendingmachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendingmachine.Contract.contract.Transact(opts, method, params...)
}

// Vendors is a free data retrieval call binding the contract method 0x27e30c36.
//
// Solidity: function vendors(address ) constant returns(bool)
func (_Vendingmachine *VendingmachineCaller) Vendors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vendingmachine.contract.Call(opts, out, "vendors", arg0)
	return *ret0, err
}

// Vendors is a free data retrieval call binding the contract method 0x27e30c36.
//
// Solidity: function vendors(address ) constant returns(bool)
func (_Vendingmachine *VendingmachineSession) Vendors(arg0 common.Address) (bool, error) {
	return _Vendingmachine.Contract.Vendors(&_Vendingmachine.CallOpts, arg0)
}

// Vendors is a free data retrieval call binding the contract method 0x27e30c36.
//
// Solidity: function vendors(address ) constant returns(bool)
func (_Vendingmachine *VendingmachineCallerSession) Vendors(arg0 common.Address) (bool, error) {
	return _Vendingmachine.Contract.Vendors(&_Vendingmachine.CallOpts, arg0)
}

// AddVendor is a paid mutator transaction binding the contract method 0x45e0f753.
//
// Solidity: function addVendor(address _vendorContract) returns(bool)
func (_Vendingmachine *VendingmachineTransactor) AddVendor(opts *bind.TransactOpts, _vendorContract common.Address) (*types.Transaction, error) {
	return _Vendingmachine.contract.Transact(opts, "addVendor", _vendorContract)
}

// AddVendor is a paid mutator transaction binding the contract method 0x45e0f753.
//
// Solidity: function addVendor(address _vendorContract) returns(bool)
func (_Vendingmachine *VendingmachineSession) AddVendor(_vendorContract common.Address) (*types.Transaction, error) {
	return _Vendingmachine.Contract.AddVendor(&_Vendingmachine.TransactOpts, _vendorContract)
}

// AddVendor is a paid mutator transaction binding the contract method 0x45e0f753.
//
// Solidity: function addVendor(address _vendorContract) returns(bool)
func (_Vendingmachine *VendingmachineTransactorSession) AddVendor(_vendorContract common.Address) (*types.Transaction, error) {
	return _Vendingmachine.Contract.AddVendor(&_Vendingmachine.TransactOpts, _vendorContract)
}
