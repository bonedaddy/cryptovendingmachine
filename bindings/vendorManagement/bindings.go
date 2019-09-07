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
const VendormanagementABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"soldAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"removeProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"addProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_locations\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"registerProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_locations\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"ProductRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"ProductLocationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"ProductLocationRemoved\",\"type\":\"event\"}]"

// VendormanagementBin is the compiled bytecode used for deploying new contracts.
var VendormanagementBin = "0x608060405234801561001057600080fd5b5032604051602001610022919061006b565b60408051601f198184030181529190528051602090910120600155600080546001600160a01b031916321790556100af565b61006561006082610080565b61009d565b82525050565b60006100778284610054565b50601401919050565b600061008b82610091565b92915050565b6001600160a01b031690565b600061008b82600061008b8260601b90565b610b2c806100be6000396000f3fe6080604052600436106100705760003560e01c80636c8e745c1161004e5780636c8e745c146101315780638da5cb5b14610151578063af640d0f14610173578063c43df6aa1461019557610070565b80630186a423146100ad578063338a6d10146100e45780635d85ed1314610111575b600080546040516001600160a01b03909116913480156108fc02929091818181858888f193505050501580156100aa573d6000803e3d6000fd5b50005b3480156100b957600080fd5b506100cd6100c83660046106cc565b6101b5565b6040516100db9291906109b8565b60405180910390f35b3480156100f057600080fd5b506101046100ff366004610785565b610263565b6040516100db9190610943565b34801561011d57600080fd5b5061010461012c366004610785565b6102a0565b34801561013d57600080fd5b5061010461014c366004610785565b610353565b34801561015d57600080fd5b506101666103f7565b6040516100db9190610935565b34801561017f57600080fd5b50610188610406565b6040516100db9190610951565b3480156101a157600080fd5b506101046101b0366004610709565b61040c565b805180820160209081018051600280835293830194830194909420939052825460408051601f6000196101006001861615020190931694909404918201839004830284018301905280835283918301828280156102535780601f1061022857610100808354040283529160200191610253565b820191906000526020600020905b81548152906001019060200180831161023657829003601f168201915b5050505050908060010154905082565b8151602081840181018051600382529282019482019490942091909352815180830184018051928152908401929093019190912091525460ff1681565b60006102aa61054b565b6102cf5760405162461bcd60e51b81526004016102c6906109d8565b60405180910390fd5b6003836040516102df9190610929565b9081526020016040518091039020826040516102fb9190610929565b908152604051908190036020018120805460ff191690557f6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f906103419085908590610993565b60405180910390a15060015b92915050565b600061035d61054b565b6103795760405162461bcd60e51b81526004016102c6906109d8565b600160038460405161038b9190610929565b9081526020016040518091039020836040516103a79190610929565b908152604051908190036020018120805492151560ff19909316929092179091557f20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc906103419085908590610993565b6000546001600160a01b031681565b60015481565b600061041661054b565b6104325760405162461bcd60e51b81526004016102c6906109d8565b6040518060400160405280858152602001838152506002856040516104579190610929565b9081526020016040518091039020600082015181600001908051906020019061048192919061056e565b506020919091015160019091015560005b83518110156105045760016003866040516104ad9190610929565b90815260200160405180910390208583815181106104c757fe5b60200260200101516040516104dc9190610929565b908152604051908190036020019020805491151560ff19909216919091179055600101610492565b507f3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d8484846040516105389392919061095f565b60405180910390a15060015b9392505050565b600080546001600160a01b03163314156105675750600161056b565b5060005b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105af57805160ff19168380011785556105dc565b828001600101855582156105dc579182015b828111156105dc5782518255916020019190600101906105c1565b506105e89291506105ec565b5090565b61056b91905b808211156105e857600081556001016105f2565b600082601f83011261061757600080fd5b813561062a61062582610a0f565b6109e8565b81815260209384019390925082018360005b8381101561066857813586016106528882610672565b845250602092830192919091019060010161063c565b5050505092915050565b600082601f83011261068357600080fd5b813561069161062582610a30565b915080825260208301602083018583830111156106ad57600080fd5b6106b8838284610a8c565b50505092915050565b803561034d81610ad2565b6000602082840312156106de57600080fd5b813567ffffffffffffffff8111156106f557600080fd5b61070184828501610672565b949350505050565b60008060006060848603121561071e57600080fd5b833567ffffffffffffffff81111561073557600080fd5b61074186828701610672565b935050602084013567ffffffffffffffff81111561075e57600080fd5b61076a86828701610606565b925050604061077b868287016106c1565b9150509250925092565b6000806040838503121561079857600080fd5b823567ffffffffffffffff8111156107af57600080fd5b6107bb85828601610672565b925050602083013567ffffffffffffffff8111156107d857600080fd5b6107e485828601610672565b9150509250929050565b60006105448383610889565b61080381610a70565b82525050565b600061081482610a5e565b61081e8185610a62565b93508360208202850161083085610a58565b8060005b8581101561086a578484038952815161084d85826107ee565b945061085883610a58565b60209a909a0199925050600101610834565b5091979650505050505050565b61080381610a7b565b6108038161056b565b600061089482610a5e565b61089e8185610a62565b93506108ae818560208601610a98565b6108b781610ac8565b9093019392505050565b60006108cc82610a5e565b6108d68185610a6b565b93506108e6818560208601610a98565b9290920192915050565b60006108fd601783610a62565b7f63616c6c6572206d7573742062652076656e646f726564000000000000000000815260200192915050565b600061054482846108c1565b6020810161034d82846107fa565b6020810161034d8284610877565b6020810161034d8284610880565b606080825281016109708186610889565b905081810360208301526109848185610809565b90506107016040830184610880565b604080825281016109a48185610889565b905081810360208301526107018184610889565b604080825281016109c98185610889565b90506105446020830184610880565b6020808252810161034d816108f0565b60405181810167ffffffffffffffff81118282101715610a0757600080fd5b604052919050565b600067ffffffffffffffff821115610a2657600080fd5b5060209081020190565b600067ffffffffffffffff821115610a4757600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b600061034d82610a80565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015610ab3578181015183820152602001610a9b565b83811115610ac2576000848401525b50505050565b601f01601f191690565b610adb8161056b565b8114610ae657600080fd5b5056fea365627a7a7231582068574bf963d1750dd69316317ebed21e0ff0999b71de9e499cb4ad2108e32d726c6578706572696d656e74616cf564736f6c634300050b0040"

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
