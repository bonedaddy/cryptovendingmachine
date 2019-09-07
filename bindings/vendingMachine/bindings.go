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
const VendingmachineABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"backend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vendors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vendorContract\",\"type\":\"address\"}],\"name\":\"addVendor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"location\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vendorContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_backend\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VendingmachineBin is the compiled bytecode used for deploying new contracts.
var VendingmachineBin = "0x608060405234801561001057600080fd5b506040516105cf3803806105cf8339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825164010000000081118282018810171561008257600080fd5b82525081516020918201929091019080838360005b838110156100af578181015183820152602001610097565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b5060405260209081015184519093506100fb9250600091850190610122565b50600180546001600160a01b0319166001600160a01b0392909216919091179055506101bd565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061016357805160ff1916838001178555610190565b82800160010185558215610190579182015b82811115610190578251825591602001919060010190610175565b5061019c9291506101a0565b5090565b6101ba91905b8082111561019c57600081556001016101a6565b90565b610403806101cc6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063099e41331461005c57806327e30c361461008057806345e0f753146100ba578063516f279e146100e0578063a46ead241461015d575b600080fd5b610064610183565b604080516001600160a01b039092168252519081900360200190f35b6100a66004803603602081101561009657600080fd5b50356001600160a01b0316610192565b604080519115158252519081900360200190f35b6100a6600480360360208110156100d057600080fd5b50356001600160a01b03166101a7565b6100e8610325565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561012257818101518382015260200161010a565b50505050905090810190601f16801561014f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100646004803603602081101561017357600080fd5b50356001600160a01b03166103b3565b6001546001600160a01b031681565b60026020526000908152604090205460ff1681565b3360009081526002602052604081205460ff161561020c576040805162461bcd60e51b815260206004820152601d60248201527f76656e646f72206d757374206e6f742062652072656769737465726564000000604482015290519081900360640190fd5b6000829050336001600160a01b0316816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561025457600080fd5b505afa158015610268573d6000803e3d6000fd5b505050506040513d602081101561027e57600080fd5b50516001600160a01b0316146102db576040805162461bcd60e51b815260206004820152601d60248201527f63616c6c6572206d7573742062652076656e646f72206d616e61676572000000604482015290519081900360640190fd5b5050336000908152600260209081526040808320805460ff19166001908117909155600390925290912080546001600160a01b0384166001600160a01b0319909116179055919050565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103ab5780601f10610380576101008083540402835291602001916103ab565b820191906000526020600020905b81548152906001019060200180831161038e57829003601f168201915b505050505081565b6003602052600090815260409020546001600160a01b03168156fea265627a7a723158206e5ccb20d264b7f4bbb556ca87ebc3d3f2db53299989ff9cf92342f8eeee313664736f6c634300050b0032"

// DeployVendingmachine deploys a new Ethereum contract, binding an instance of Vendingmachine to it.
func DeployVendingmachine(auth *bind.TransactOpts, backend bind.ContractBackend, _location string, _backend common.Address) (common.Address, *types.Transaction, *Vendingmachine, error) {
	parsed, err := abi.JSON(strings.NewReader(VendingmachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VendingmachineBin), backend, _location, _backend)
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

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() constant returns(address)
func (_Vendingmachine *VendingmachineCaller) Backend(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendingmachine.contract.Call(opts, out, "backend")
	return *ret0, err
}

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() constant returns(address)
func (_Vendingmachine *VendingmachineSession) Backend() (common.Address, error) {
	return _Vendingmachine.Contract.Backend(&_Vendingmachine.CallOpts)
}

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() constant returns(address)
func (_Vendingmachine *VendingmachineCallerSession) Backend() (common.Address, error) {
	return _Vendingmachine.Contract.Backend(&_Vendingmachine.CallOpts)
}

// Location is a free data retrieval call binding the contract method 0x516f279e.
//
// Solidity: function location() constant returns(string)
func (_Vendingmachine *VendingmachineCaller) Location(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Vendingmachine.contract.Call(opts, out, "location")
	return *ret0, err
}

// Location is a free data retrieval call binding the contract method 0x516f279e.
//
// Solidity: function location() constant returns(string)
func (_Vendingmachine *VendingmachineSession) Location() (string, error) {
	return _Vendingmachine.Contract.Location(&_Vendingmachine.CallOpts)
}

// Location is a free data retrieval call binding the contract method 0x516f279e.
//
// Solidity: function location() constant returns(string)
func (_Vendingmachine *VendingmachineCallerSession) Location() (string, error) {
	return _Vendingmachine.Contract.Location(&_Vendingmachine.CallOpts)
}

// VendorContracts is a free data retrieval call binding the contract method 0xa46ead24.
//
// Solidity: function vendorContracts(address ) constant returns(address)
func (_Vendingmachine *VendingmachineCaller) VendorContracts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendingmachine.contract.Call(opts, out, "vendorContracts", arg0)
	return *ret0, err
}

// VendorContracts is a free data retrieval call binding the contract method 0xa46ead24.
//
// Solidity: function vendorContracts(address ) constant returns(address)
func (_Vendingmachine *VendingmachineSession) VendorContracts(arg0 common.Address) (common.Address, error) {
	return _Vendingmachine.Contract.VendorContracts(&_Vendingmachine.CallOpts, arg0)
}

// VendorContracts is a free data retrieval call binding the contract method 0xa46ead24.
//
// Solidity: function vendorContracts(address ) constant returns(address)
func (_Vendingmachine *VendingmachineCallerSession) VendorContracts(arg0 common.Address) (common.Address, error) {
	return _Vendingmachine.Contract.VendorContracts(&_Vendingmachine.CallOpts, arg0)
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
