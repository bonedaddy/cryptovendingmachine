// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vendormanagement

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

// VendormanagementABI is the input ABI used to generate the binding from.
const VendormanagementABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"soldAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"removeProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"addProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_locations\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"registerProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_locations\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"ProductRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"ProductLocationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"ProductLocationRemoved\",\"type\":\"event\"}]"

// VendormanagementBin is the compiled bytecode used for deploying new contracts.
var VendormanagementBin = "0x608060405234801561001057600080fd5b5032604051602001610022919061006b565b60408051601f198184030181529190528051602090910120600155600080546001600160a01b031916321790556100af565b61006561006082610080565b61009d565b82525050565b60006100778284610054565b50601401919050565b600061008b82610091565b92915050565b6001600160a01b031690565b600061008b82600061008b8260601b90565b610aa6806100be6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636c8e745c1161005b5780636c8e745c146100df5780638da5cb5b146100f2578063af640d0f14610107578063c43df6aa1461011c5761007d565b80630186a42314610082578063338a6d10146100ac5780635d85ed13146100cc575b600080fd5b610095610090366004610646565b61012f565b6040516100a3929190610932565b60405180910390f35b6100bf6100ba3660046106ff565b6101dd565b6040516100a391906108bd565b6100bf6100da3660046106ff565b61021a565b6100bf6100ed3660046106ff565b6102cd565b6100fa610371565b6040516100a391906108af565b61010f610380565b6040516100a391906108cb565b6100bf61012a366004610683565b610386565b805180820160209081018051600280835293830194830194909420939052825460408051601f6000196101006001861615020190931694909404918201839004830284018301905280835283918301828280156101cd5780601f106101a2576101008083540402835291602001916101cd565b820191906000526020600020905b8154815290600101906020018083116101b057829003601f168201915b5050505050908060010154905082565b8151602081840181018051600382529282019482019490942091909352815180830184018051928152908401929093019190912091525460ff1681565b60006102246104c5565b6102495760405162461bcd60e51b815260040161024090610952565b60405180910390fd5b60038360405161025991906108a3565b90815260200160405180910390208260405161027591906108a3565b908152604051908190036020018120805460ff191690557f6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f906102bb908590859061090d565b60405180910390a15060015b92915050565b60006102d76104c5565b6102f35760405162461bcd60e51b815260040161024090610952565b600160038460405161030591906108a3565b90815260200160405180910390208360405161032191906108a3565b908152604051908190036020018120805492151560ff19909316929092179091557f20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc906102bb908590859061090d565b6000546001600160a01b031681565b60015481565b60006103906104c5565b6103ac5760405162461bcd60e51b815260040161024090610952565b6040518060400160405280858152602001838152506002856040516103d191906108a3565b908152602001604051809103902060008201518160000190805190602001906103fb9291906104e8565b506020919091015160019091015560005b835181101561047e57600160038660405161042791906108a3565b908152602001604051809103902085838151811061044157fe5b602002602001015160405161045691906108a3565b908152604051908190036020019020805491151560ff1990921691909117905560010161040c565b507f3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d8484846040516104b2939291906108d9565b60405180910390a15060015b9392505050565b600080546001600160a01b03163314156104e1575060016104e5565b5060005b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061052957805160ff1916838001178555610556565b82800160010185558215610556579182015b8281111561055657825182559160200191906001019061053b565b50610562929150610566565b5090565b6104e591905b80821115610562576000815560010161056c565b600082601f83011261059157600080fd5b81356105a461059f82610989565b610962565b81815260209384019390925082018360005b838110156105e257813586016105cc88826105ec565b84525060209283019291909101906001016105b6565b5050505092915050565b600082601f8301126105fd57600080fd5b813561060b61059f826109aa565b9150808252602083016020830185838301111561062757600080fd5b610632838284610a06565b50505092915050565b80356102c781610a4c565b60006020828403121561065857600080fd5b813567ffffffffffffffff81111561066f57600080fd5b61067b848285016105ec565b949350505050565b60008060006060848603121561069857600080fd5b833567ffffffffffffffff8111156106af57600080fd5b6106bb868287016105ec565b935050602084013567ffffffffffffffff8111156106d857600080fd5b6106e486828701610580565b92505060406106f58682870161063b565b9150509250925092565b6000806040838503121561071257600080fd5b823567ffffffffffffffff81111561072957600080fd5b610735858286016105ec565b925050602083013567ffffffffffffffff81111561075257600080fd5b61075e858286016105ec565b9150509250929050565b60006104be8383610803565b61077d816109ea565b82525050565b600061078e826109d8565b61079881856109dc565b9350836020820285016107aa856109d2565b8060005b858110156107e457848403895281516107c78582610768565b94506107d2836109d2565b60209a909a01999250506001016107ae565b5091979650505050505050565b61077d816109f5565b61077d816104e5565b600061080e826109d8565b61081881856109dc565b9350610828818560208601610a12565b61083181610a42565b9093019392505050565b6000610846826109d8565b61085081856109e5565b9350610860818560208601610a12565b9290920192915050565b60006108776017836109dc565b7f63616c6c6572206d7573742062652076656e646f726564000000000000000000815260200192915050565b60006104be828461083b565b602081016102c78284610774565b602081016102c782846107f1565b602081016102c782846107fa565b606080825281016108ea8186610803565b905081810360208301526108fe8185610783565b905061067b60408301846107fa565b6040808252810161091e8185610803565b9050818103602083015261067b8184610803565b604080825281016109438185610803565b90506104be60208301846107fa565b602080825281016102c78161086a565b60405181810167ffffffffffffffff8111828210171561098157600080fd5b604052919050565b600067ffffffffffffffff8211156109a057600080fd5b5060209081020190565b600067ffffffffffffffff8211156109c157600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b60006102c7826109fa565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015610a2d578181015183820152602001610a15565b83811115610a3c576000848401525b50505050565b601f01601f191690565b610a55816104e5565b8114610a6057600080fd5b5056fea365627a7a72315820f44cb1d0f3bb33f59925b29ec621ba46758af58825f48fcc0ec7530d5bb27e686c6578706572696d656e74616cf564736f6c634300050b0040"

// DeployVendormanagement deploys a new Ethereum contract, binding an instance of Vendormanagement to it.
func DeployVendormanagement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vendormanagement, error) {
	parsed, err := abi.JSON(strings.NewReader(VendormanagementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VendormanagementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vendormanagement{VendormanagementCaller: VendormanagementCaller{contract: contract}, VendormanagementTransactor: VendormanagementTransactor{contract: contract}, VendormanagementFilterer: VendormanagementFilterer{contract: contract}}, nil
}

// Vendormanagement is an auto generated Go binding around an Ethereum contract.
type Vendormanagement struct {
	VendormanagementCaller     // Read-only binding to the contract
	VendormanagementTransactor // Write-only binding to the contract
	VendormanagementFilterer   // Log filterer for contract events
}

// VendormanagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type VendormanagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendormanagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VendormanagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendormanagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VendormanagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendormanagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VendormanagementSession struct {
	Contract     *Vendormanagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VendormanagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VendormanagementCallerSession struct {
	Contract *VendormanagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// VendormanagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VendormanagementTransactorSession struct {
	Contract     *VendormanagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// VendormanagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type VendormanagementRaw struct {
	Contract *Vendormanagement // Generic contract binding to access the raw methods on
}

// VendormanagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VendormanagementCallerRaw struct {
	Contract *VendormanagementCaller // Generic read-only contract binding to access the raw methods on
}

// VendormanagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VendormanagementTransactorRaw struct {
	Contract *VendormanagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVendormanagement creates a new instance of Vendormanagement, bound to a specific deployed contract.
func NewVendormanagement(address common.Address, backend bind.ContractBackend) (*Vendormanagement, error) {
	contract, err := bindVendormanagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vendormanagement{VendormanagementCaller: VendormanagementCaller{contract: contract}, VendormanagementTransactor: VendormanagementTransactor{contract: contract}, VendormanagementFilterer: VendormanagementFilterer{contract: contract}}, nil
}

// NewVendormanagementCaller creates a new read-only instance of Vendormanagement, bound to a specific deployed contract.
func NewVendormanagementCaller(address common.Address, caller bind.ContractCaller) (*VendormanagementCaller, error) {
	contract, err := bindVendormanagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VendormanagementCaller{contract: contract}, nil
}

// NewVendormanagementTransactor creates a new write-only instance of Vendormanagement, bound to a specific deployed contract.
func NewVendormanagementTransactor(address common.Address, transactor bind.ContractTransactor) (*VendormanagementTransactor, error) {
	contract, err := bindVendormanagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VendormanagementTransactor{contract: contract}, nil
}

// NewVendormanagementFilterer creates a new log filterer instance of Vendormanagement, bound to a specific deployed contract.
func NewVendormanagementFilterer(address common.Address, filterer bind.ContractFilterer) (*VendormanagementFilterer, error) {
	contract, err := bindVendormanagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VendormanagementFilterer{contract: contract}, nil
}

// bindVendormanagement binds a generic wrapper to an already deployed contract.
func bindVendormanagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VendormanagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendormanagement *VendormanagementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendormanagement.Contract.VendormanagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendormanagement *VendormanagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendormanagement.Contract.VendormanagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendormanagement *VendormanagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendormanagement.Contract.VendormanagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendormanagement *VendormanagementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendormanagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendormanagement *VendormanagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendormanagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendormanagement *VendormanagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendormanagement.Contract.contract.Transact(opts, method, params...)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_Vendormanagement *VendormanagementCaller) Id(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Vendormanagement.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_Vendormanagement *VendormanagementSession) Id() ([32]byte, error) {
	return _Vendormanagement.Contract.Id(&_Vendormanagement.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() constant returns(bytes32)
func (_Vendormanagement *VendormanagementCallerSession) Id() ([32]byte, error) {
	return _Vendormanagement.Contract.Id(&_Vendormanagement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vendormanagement *VendormanagementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendormanagement.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vendormanagement *VendormanagementSession) Owner() (common.Address, error) {
	return _Vendormanagement.Contract.Owner(&_Vendormanagement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vendormanagement *VendormanagementCallerSession) Owner() (common.Address, error) {
	return _Vendormanagement.Contract.Owner(&_Vendormanagement.CallOpts)
}

// Products is a free data retrieval call binding the contract method 0x0186a423.
//
// Solidity: function products(string ) constant returns(string name, uint256 cost)
func (_Vendormanagement *VendormanagementCaller) Products(opts *bind.CallOpts, arg0 string) (struct {
	Name string
	Cost *big.Int
}, error) {
	ret := new(struct {
		Name string
		Cost *big.Int
	})
	out := ret
	err := _Vendormanagement.contract.Call(opts, out, "products", arg0)
	return *ret, err
}

// Products is a free data retrieval call binding the contract method 0x0186a423.
//
// Solidity: function products(string ) constant returns(string name, uint256 cost)
func (_Vendormanagement *VendormanagementSession) Products(arg0 string) (struct {
	Name string
	Cost *big.Int
}, error) {
	return _Vendormanagement.Contract.Products(&_Vendormanagement.CallOpts, arg0)
}

// Products is a free data retrieval call binding the contract method 0x0186a423.
//
// Solidity: function products(string ) constant returns(string name, uint256 cost)
func (_Vendormanagement *VendormanagementCallerSession) Products(arg0 string) (struct {
	Name string
	Cost *big.Int
}, error) {
	return _Vendormanagement.Contract.Products(&_Vendormanagement.CallOpts, arg0)
}

// SoldAt is a free data retrieval call binding the contract method 0x338a6d10.
//
// Solidity: function soldAt(string , string ) constant returns(bool)
func (_Vendormanagement *VendormanagementCaller) SoldAt(opts *bind.CallOpts, arg0 string, arg1 string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vendormanagement.contract.Call(opts, out, "soldAt", arg0, arg1)
	return *ret0, err
}

// SoldAt is a free data retrieval call binding the contract method 0x338a6d10.
//
// Solidity: function soldAt(string , string ) constant returns(bool)
func (_Vendormanagement *VendormanagementSession) SoldAt(arg0 string, arg1 string) (bool, error) {
	return _Vendormanagement.Contract.SoldAt(&_Vendormanagement.CallOpts, arg0, arg1)
}

// SoldAt is a free data retrieval call binding the contract method 0x338a6d10.
//
// Solidity: function soldAt(string , string ) constant returns(bool)
func (_Vendormanagement *VendormanagementCallerSession) SoldAt(arg0 string, arg1 string) (bool, error) {
	return _Vendormanagement.Contract.SoldAt(&_Vendormanagement.CallOpts, arg0, arg1)
}

// AddProductLocation is a paid mutator transaction binding the contract method 0x6c8e745c.
//
// Solidity: function addProductLocation(string _name, string _location) returns(bool)
func (_Vendormanagement *VendormanagementTransactor) AddProductLocation(opts *bind.TransactOpts, _name string, _location string) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "addProductLocation", _name, _location)
}

// AddProductLocation is a paid mutator transaction binding the contract method 0x6c8e745c.
//
// Solidity: function addProductLocation(string _name, string _location) returns(bool)
func (_Vendormanagement *VendormanagementSession) AddProductLocation(_name string, _location string) (*types.Transaction, error) {
	return _Vendormanagement.Contract.AddProductLocation(&_Vendormanagement.TransactOpts, _name, _location)
}

// AddProductLocation is a paid mutator transaction binding the contract method 0x6c8e745c.
//
// Solidity: function addProductLocation(string _name, string _location) returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) AddProductLocation(_name string, _location string) (*types.Transaction, error) {
	return _Vendormanagement.Contract.AddProductLocation(&_Vendormanagement.TransactOpts, _name, _location)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0xc43df6aa.
//
// Solidity: function registerProduct(string _name, string[] _locations, uint256 _cost) returns(bool)
func (_Vendormanagement *VendormanagementTransactor) RegisterProduct(opts *bind.TransactOpts, _name string, _locations []string, _cost *big.Int) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "registerProduct", _name, _locations, _cost)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0xc43df6aa.
//
// Solidity: function registerProduct(string _name, string[] _locations, uint256 _cost) returns(bool)
func (_Vendormanagement *VendormanagementSession) RegisterProduct(_name string, _locations []string, _cost *big.Int) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RegisterProduct(&_Vendormanagement.TransactOpts, _name, _locations, _cost)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0xc43df6aa.
//
// Solidity: function registerProduct(string _name, string[] _locations, uint256 _cost) returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) RegisterProduct(_name string, _locations []string, _cost *big.Int) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RegisterProduct(&_Vendormanagement.TransactOpts, _name, _locations, _cost)
}

// RemoveProductLocation is a paid mutator transaction binding the contract method 0x5d85ed13.
//
// Solidity: function removeProductLocation(string _name, string _location) returns(bool)
func (_Vendormanagement *VendormanagementTransactor) RemoveProductLocation(opts *bind.TransactOpts, _name string, _location string) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "removeProductLocation", _name, _location)
}

// RemoveProductLocation is a paid mutator transaction binding the contract method 0x5d85ed13.
//
// Solidity: function removeProductLocation(string _name, string _location) returns(bool)
func (_Vendormanagement *VendormanagementSession) RemoveProductLocation(_name string, _location string) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RemoveProductLocation(&_Vendormanagement.TransactOpts, _name, _location)
}

// RemoveProductLocation is a paid mutator transaction binding the contract method 0x5d85ed13.
//
// Solidity: function removeProductLocation(string _name, string _location) returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) RemoveProductLocation(_name string, _location string) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RemoveProductLocation(&_Vendormanagement.TransactOpts, _name, _location)
}

// VendormanagementProductLocationAddedIterator is returned from FilterProductLocationAdded and is used to iterate over the raw logs and unpacked data for ProductLocationAdded events raised by the Vendormanagement contract.
type VendormanagementProductLocationAddedIterator struct {
	Event *VendormanagementProductLocationAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VendormanagementProductLocationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendormanagementProductLocationAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VendormanagementProductLocationAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VendormanagementProductLocationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendormanagementProductLocationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendormanagementProductLocationAdded represents a ProductLocationAdded event raised by the Vendormanagement contract.
type VendormanagementProductLocationAdded struct {
	Name     string
	Location string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProductLocationAdded is a free log retrieval operation binding the contract event 0x20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc.
//
// Solidity: event ProductLocationAdded(string _name, string _location)
func (_Vendormanagement *VendormanagementFilterer) FilterProductLocationAdded(opts *bind.FilterOpts) (*VendormanagementProductLocationAddedIterator, error) {

	logs, sub, err := _Vendormanagement.contract.FilterLogs(opts, "ProductLocationAdded")
	if err != nil {
		return nil, err
	}
	return &VendormanagementProductLocationAddedIterator{contract: _Vendormanagement.contract, event: "ProductLocationAdded", logs: logs, sub: sub}, nil
}

// WatchProductLocationAdded is a free log subscription operation binding the contract event 0x20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc.
//
// Solidity: event ProductLocationAdded(string _name, string _location)
func (_Vendormanagement *VendormanagementFilterer) WatchProductLocationAdded(opts *bind.WatchOpts, sink chan<- *VendormanagementProductLocationAdded) (event.Subscription, error) {

	logs, sub, err := _Vendormanagement.contract.WatchLogs(opts, "ProductLocationAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendormanagementProductLocationAdded)
				if err := _Vendormanagement.contract.UnpackLog(event, "ProductLocationAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProductLocationAdded is a log parse operation binding the contract event 0x20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc.
//
// Solidity: event ProductLocationAdded(string _name, string _location)
func (_Vendormanagement *VendormanagementFilterer) ParseProductLocationAdded(log types.Log) (*VendormanagementProductLocationAdded, error) {
	event := new(VendormanagementProductLocationAdded)
	if err := _Vendormanagement.contract.UnpackLog(event, "ProductLocationAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VendormanagementProductLocationRemovedIterator is returned from FilterProductLocationRemoved and is used to iterate over the raw logs and unpacked data for ProductLocationRemoved events raised by the Vendormanagement contract.
type VendormanagementProductLocationRemovedIterator struct {
	Event *VendormanagementProductLocationRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VendormanagementProductLocationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendormanagementProductLocationRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VendormanagementProductLocationRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VendormanagementProductLocationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendormanagementProductLocationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendormanagementProductLocationRemoved represents a ProductLocationRemoved event raised by the Vendormanagement contract.
type VendormanagementProductLocationRemoved struct {
	Name     string
	Location string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProductLocationRemoved is a free log retrieval operation binding the contract event 0x6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f.
//
// Solidity: event ProductLocationRemoved(string _name, string _location)
func (_Vendormanagement *VendormanagementFilterer) FilterProductLocationRemoved(opts *bind.FilterOpts) (*VendormanagementProductLocationRemovedIterator, error) {

	logs, sub, err := _Vendormanagement.contract.FilterLogs(opts, "ProductLocationRemoved")
	if err != nil {
		return nil, err
	}
	return &VendormanagementProductLocationRemovedIterator{contract: _Vendormanagement.contract, event: "ProductLocationRemoved", logs: logs, sub: sub}, nil
}

// WatchProductLocationRemoved is a free log subscription operation binding the contract event 0x6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f.
//
// Solidity: event ProductLocationRemoved(string _name, string _location)
func (_Vendormanagement *VendormanagementFilterer) WatchProductLocationRemoved(opts *bind.WatchOpts, sink chan<- *VendormanagementProductLocationRemoved) (event.Subscription, error) {

	logs, sub, err := _Vendormanagement.contract.WatchLogs(opts, "ProductLocationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendormanagementProductLocationRemoved)
				if err := _Vendormanagement.contract.UnpackLog(event, "ProductLocationRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProductLocationRemoved is a log parse operation binding the contract event 0x6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f.
//
// Solidity: event ProductLocationRemoved(string _name, string _location)
func (_Vendormanagement *VendormanagementFilterer) ParseProductLocationRemoved(log types.Log) (*VendormanagementProductLocationRemoved, error) {
	event := new(VendormanagementProductLocationRemoved)
	if err := _Vendormanagement.contract.UnpackLog(event, "ProductLocationRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VendormanagementProductRegisteredIterator is returned from FilterProductRegistered and is used to iterate over the raw logs and unpacked data for ProductRegistered events raised by the Vendormanagement contract.
type VendormanagementProductRegisteredIterator struct {
	Event *VendormanagementProductRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VendormanagementProductRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendormanagementProductRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VendormanagementProductRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VendormanagementProductRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendormanagementProductRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendormanagementProductRegistered represents a ProductRegistered event raised by the Vendormanagement contract.
type VendormanagementProductRegistered struct {
	Name      string
	Locations []string
	Cost      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProductRegistered is a free log retrieval operation binding the contract event 0x3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d.
//
// Solidity: event ProductRegistered(string _name, string[] _locations, uint256 _cost)
func (_Vendormanagement *VendormanagementFilterer) FilterProductRegistered(opts *bind.FilterOpts) (*VendormanagementProductRegisteredIterator, error) {

	logs, sub, err := _Vendormanagement.contract.FilterLogs(opts, "ProductRegistered")
	if err != nil {
		return nil, err
	}
	return &VendormanagementProductRegisteredIterator{contract: _Vendormanagement.contract, event: "ProductRegistered", logs: logs, sub: sub}, nil
}

// WatchProductRegistered is a free log subscription operation binding the contract event 0x3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d.
//
// Solidity: event ProductRegistered(string _name, string[] _locations, uint256 _cost)
func (_Vendormanagement *VendormanagementFilterer) WatchProductRegistered(opts *bind.WatchOpts, sink chan<- *VendormanagementProductRegistered) (event.Subscription, error) {

	logs, sub, err := _Vendormanagement.contract.WatchLogs(opts, "ProductRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendormanagementProductRegistered)
				if err := _Vendormanagement.contract.UnpackLog(event, "ProductRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProductRegistered is a log parse operation binding the contract event 0x3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d.
//
// Solidity: event ProductRegistered(string _name, string[] _locations, uint256 _cost)
func (_Vendormanagement *VendormanagementFilterer) ParseProductRegistered(log types.Log) (*VendormanagementProductRegistered, error) {
	event := new(VendormanagementProductRegistered)
	if err := _Vendormanagement.contract.UnpackLog(event, "ProductRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
