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
const VendormanagementABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"soldAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"removeProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"addProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_locations\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"registerProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_locations\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"ProductRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"ProductLocationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"ProductLocationRemoved\",\"type\":\"event\"}]"

// VendormanagementBin is the compiled bytecode used for deploying new contracts.
var VendormanagementBin = "0x608060405234801561001057600080fd5b5032604051602001610022919061006b565b60408051601f198184030181529190528051602090910120600155600080546001600160a01b031916321790556100af565b61006561006082610080565b61009d565b82525050565b60006100778284610054565b50601401919050565b600061008b82610091565b92915050565b6001600160a01b031690565b600061008b82600061008b8260601b90565b610cbd806100be6000396000f3fe6080604052600436106100865760003560e01c80635d85ed13116100595780635d85ed13146101165780636c8e745c146101365780638da5cb5b14610156578063af640d0f14610178578063c43df6aa1461019a57610086565b80630186a4231461008857806324600fc3146100bf578063338a6d10146100e15780635c388ca614610101575b005b34801561009457600080fd5b506100a86100a3366004610791565b6101ba565b6040516100b6929190610b19565b60405180910390f35b3480156100cb57600080fd5b506100d461026d565b6040516100b69190610aa4565b3480156100ed57600080fd5b506100d46100fc36600461084a565b610329565b34801561010d57600080fd5b506100d4610366565b34801561012257600080fd5b506100d461013136600461084a565b61036f565b34801561014257600080fd5b506100d461015136600461084a565b610419565b34801561016257600080fd5b5061016b6104bd565b6040516100b69190610a96565b34801561018457600080fd5b5061018d6104cc565b6040516100b69190610ab2565b3480156101a657600080fd5b506100d46101b53660046107ce565b6104d2565b805160208183018101805160038252928201938201939093209190925280546040805160026001841615610100026000190190931692909204601f810185900485028301850190915280825291929091839183018282801561025d5780601f106102325761010080835404028352916020019161025d565b820191906000526020600020905b81548152906001019060200180831161024057829003601f168201915b5050505050908060010154905082565b6000610277610611565b61029c5760405162461bcd60e51b815260040161029390610b69565b60405180910390fd5b30316102ba5760405162461bcd60e51b815260040161029390610b39565b60025460ff16156102dd5760405162461bcd60e51b815260040161029390610b59565b6002805460ff191660011790556040513390303180156108fc02916000818181858888f19350505050158015610317573d6000803e3d6000fd5b50506002805460ff1916905560015b90565b8151602081840181018051600482529282019482019490942091909352815180830184018051928152908401929093019190912091525460ff1681565b60025460ff1681565b6000610379610611565b6103955760405162461bcd60e51b815260040161029390610b49565b6004836040516103a59190610a8a565b9081526020016040518091039020826040516103c19190610a8a565b908152604051908190036020018120805460ff191690557f6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f906104079085908590610af4565b60405180910390a15060015b92915050565b6000610423610611565b61043f5760405162461bcd60e51b815260040161029390610b49565b60016004846040516104519190610a8a565b90815260200160405180910390208360405161046d9190610a8a565b908152604051908190036020018120805492151560ff19909316929092179091557f20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc906104079085908590610af4565b6000546001600160a01b031681565b60015481565b60006104dc610611565b6104f85760405162461bcd60e51b815260040161029390610b49565b60405180604001604052808581526020018381525060038560405161051d9190610a8a565b90815260200160405180910390206000820151816000019080519060200190610547929190610633565b506020919091015160019091015560005b83518110156105ca5760016004866040516105739190610a8a565b908152602001604051809103902085838151811061058d57fe5b60200260200101516040516105a29190610a8a565b908152604051908190036020019020805491151560ff19909216919091179055600101610558565b507f3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d8484846040516105fe93929190610ac0565b60405180910390a15060015b9392505050565b600080546001600160a01b031633141561062d57506001610326565b50600090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061067457805160ff19168380011785556106a1565b828001600101855582156106a1579182015b828111156106a1578251825591602001919060010190610686565b506106ad9291506106b1565b5090565b61032691905b808211156106ad57600081556001016106b7565b600082601f8301126106dc57600080fd5b81356106ef6106ea82610ba0565b610b79565b81815260209384019390925082018360005b8381101561072d57813586016107178882610737565b8452506020928301929190910190600101610701565b5050505092915050565b600082601f83011261074857600080fd5b81356107566106ea82610bc1565b9150808252602083016020830185838301111561077257600080fd5b61077d838284610c1d565b50505092915050565b803561041381610c63565b6000602082840312156107a357600080fd5b813567ffffffffffffffff8111156107ba57600080fd5b6107c684828501610737565b949350505050565b6000806000606084860312156107e357600080fd5b833567ffffffffffffffff8111156107fa57600080fd5b61080686828701610737565b935050602084013567ffffffffffffffff81111561082357600080fd5b61082f868287016106cb565b925050604061084086828701610786565b9150509250925092565b6000806040838503121561085d57600080fd5b823567ffffffffffffffff81111561087457600080fd5b61088085828601610737565b925050602083013567ffffffffffffffff81111561089d57600080fd5b6108a985828601610737565b9150509250929050565b600061060a838361094e565b6108c881610c01565b82525050565b60006108d982610bef565b6108e38185610bf3565b9350836020820285016108f585610be9565b8060005b8581101561092f578484038952815161091285826108b3565b945061091d83610be9565b60209a909a01999250506001016108f9565b5091979650505050505050565b6108c881610c0c565b6108c881610326565b600061095982610bef565b6109638185610bf3565b9350610973818560208601610c29565b61097c81610c59565b9093019392505050565b600061099182610bef565b61099b8185610bfc565b93506109ab818560208601610c29565b9290920192915050565b60006109c2601683610bf3565b751b9bc818985b185b98d9481a5b8818dbdb9d1c9858dd60521b815260200192915050565b60006109f4601783610bf3565b7f63616c6c6572206d7573742062652076656e646f726564000000000000000000815260200192915050565b6000610a2d602083610bf3565b7f636f6e7472616374206d757374206e6f74206265207769746864726177696e67815260200192915050565b6000610a66601583610bf3565b7431b0b63632b91036bab9ba103132903b32b73237b960591b815260200192915050565b600061060a8284610986565b6020810161041382846108bf565b60208101610413828461093c565b602081016104138284610945565b60608082528101610ad1818661094e565b90508181036020830152610ae581856108ce565b90506107c66040830184610945565b60408082528101610b05818561094e565b905081810360208301526107c6818461094e565b60408082528101610b2a818561094e565b905061060a6020830184610945565b60208082528101610413816109b5565b60208082528101610413816109e7565b6020808252810161041381610a20565b6020808252810161041381610a59565b60405181810167ffffffffffffffff81118282101715610b9857600080fd5b604052919050565b600067ffffffffffffffff821115610bb757600080fd5b5060209081020190565b600067ffffffffffffffff821115610bd857600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b600061041382610c11565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015610c44578181015183820152602001610c2c565b83811115610c53576000848401525b50505050565b601f01601f191690565b610c6c81610326565b8114610c7757600080fd5b5056fea365627a7a72315820a9a2c9094e1a0cb1bcc7077e51dc09400585ef41c2da7265cb78aec80889d5e76c6578706572696d656e74616cf564736f6c634300050b0040"

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

// WithdrawLock is a free data retrieval call binding the contract method 0x5c388ca6.
//
// Solidity: function withdrawLock() constant returns(bool)
func (_Vendormanagement *VendormanagementCaller) WithdrawLock(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vendormanagement.contract.Call(opts, out, "withdrawLock")
	return *ret0, err
}

// WithdrawLock is a free data retrieval call binding the contract method 0x5c388ca6.
//
// Solidity: function withdrawLock() constant returns(bool)
func (_Vendormanagement *VendormanagementSession) WithdrawLock() (bool, error) {
	return _Vendormanagement.Contract.WithdrawLock(&_Vendormanagement.CallOpts)
}

// WithdrawLock is a free data retrieval call binding the contract method 0x5c388ca6.
//
// Solidity: function withdrawLock() constant returns(bool)
func (_Vendormanagement *VendormanagementCallerSession) WithdrawLock() (bool, error) {
	return _Vendormanagement.Contract.WithdrawLock(&_Vendormanagement.CallOpts)
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

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns(bool)
func (_Vendormanagement *VendormanagementTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "withdrawFunds")
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns(bool)
func (_Vendormanagement *VendormanagementSession) WithdrawFunds() (*types.Transaction, error) {
	return _Vendormanagement.Contract.WithdrawFunds(&_Vendormanagement.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _Vendormanagement.Contract.WithdrawFunds(&_Vendormanagement.TransactOpts)
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
