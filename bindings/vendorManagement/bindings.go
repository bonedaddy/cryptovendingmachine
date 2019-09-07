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
const VendormanagementABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_location\",\"type\":\"bytes32\"}],\"name\":\"removeProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"soldAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_locations\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"registerProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_location\",\"type\":\"bytes32\"}],\"name\":\"addProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"_locations\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"ProductRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_location\",\"type\":\"bytes32\"}],\"name\":\"ProductLocationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_location\",\"type\":\"bytes32\"}],\"name\":\"ProductLocationRemoved\",\"type\":\"event\"}]"

// VendormanagementBin is the compiled bytecode used for deploying new contracts.
var VendormanagementBin = "0x608060405234801561001057600080fd5b506040805133606081901b60208084019190915283518084036014018152603490930190935281519190920120600155600080546001600160a01b03191690911790556105a5806100626000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80639d9497881161005b5780639d94978814610113578063aa14471814610136578063af640d0f146101e2578063c4fe2f93146101fc5761007d565b8063310f7b911461008257806379054391146100b95780638da5cb5b146100ef575b600080fd5b6100a56004803603604081101561009857600080fd5b508035906020013561021f565b604080519115158252519081900360200190f35b6100d6600480360360208110156100cf57600080fd5b50356102d7565b6040805192835260208301919091528051918290030190f35b6100f76102f0565b604080516001600160a01b039092168252519081900360200190f35b6100a56004803603604081101561012957600080fd5b50803590602001356102ff565b6100a56004803603606081101561014c57600080fd5b8135919081019060408101602082013564010000000081111561016e57600080fd5b82018360208201111561018057600080fd5b803590602001918460208302840111640100000000831117156101a257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550509135925061031f915050565b6101ea61048c565b60408051918252519081900360200190f35b6100a56004803603604081101561021257600080fd5b5080359060200135610492565b600061022961054d565b610274576040805162461bcd60e51b815260206004820152601760248201527618d85b1b195c881b5d5cdd081899481d995b991bdc9959604a1b604482015290519081900360640190fd5b6000838152600360209081526040808320858452825291829020805460ff19169055815185815290810184905281517fa16e19a51eb50f485a1110b0c514e093a9a1676f77216aa019505467af0bee48929181900390910190a150600192915050565b6002602052600090815260409020805460019091015482565b6000546001600160a01b031681565b600360209081526000928352604080842090915290825290205460ff1681565b600061032961054d565b610374576040805162461bcd60e51b815260206004820152601760248201527618d85b1b195c881b5d5cdd081899481d995b991bdc9959604a1b604482015290519081900360640190fd5b6040805180820182528581526020808201858152600088815260029092529281209151825591516001909101555b83518110156103fc5760008581526003602052604081208551600192908790859081106103cb57fe5b6020908102919091018101518252810191909152604001600020805460ff19169115159190911790556001016103a2565b507f86a0e56ff60b92d5010a66e5f314a5433028ab1db0ef9bdcdb19db16b35654f18484846040518084815260200180602001838152602001828103825284818151815260200191508051906020019060200280838360005b8381101561046d578181015183820152602001610455565b5050505090500194505050505060405180910390a15060019392505050565b60015481565b600061049c61054d565b6104e7576040805162461bcd60e51b815260206004820152601760248201527618d85b1b195c881b5d5cdd081899481d995b991bdc9959604a1b604482015290519081900360640190fd5b6000838152600360209081526040808320858452825291829020805460ff19166001179055815185815290810184905281517fc16ed24d108fc30d7daeb46fc1ead0b3251c5a6cd18a0496071802e55c9f8373929181900390910190a150600192915050565b600080546001600160a01b03163314156105695750600161056d565b5060005b9056fea265627a7a72315820251b64ede4655a61a8074af9bba91daaa97f01bf022a21b4674aa7d7d2756fea64736f6c634300050b0032"

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

// Products is a free data retrieval call binding the contract method 0x79054391.
//
// Solidity: function products(bytes32 ) constant returns(bytes32 name, uint256 cost)
func (_Vendormanagement *VendormanagementCaller) Products(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Name [32]byte
	Cost *big.Int
}, error) {
	ret := new(struct {
		Name [32]byte
		Cost *big.Int
	})
	out := ret
	err := _Vendormanagement.contract.Call(opts, out, "products", arg0)
	return *ret, err
}

// Products is a free data retrieval call binding the contract method 0x79054391.
//
// Solidity: function products(bytes32 ) constant returns(bytes32 name, uint256 cost)
func (_Vendormanagement *VendormanagementSession) Products(arg0 [32]byte) (struct {
	Name [32]byte
	Cost *big.Int
}, error) {
	return _Vendormanagement.Contract.Products(&_Vendormanagement.CallOpts, arg0)
}

// Products is a free data retrieval call binding the contract method 0x79054391.
//
// Solidity: function products(bytes32 ) constant returns(bytes32 name, uint256 cost)
func (_Vendormanagement *VendormanagementCallerSession) Products(arg0 [32]byte) (struct {
	Name [32]byte
	Cost *big.Int
}, error) {
	return _Vendormanagement.Contract.Products(&_Vendormanagement.CallOpts, arg0)
}

// SoldAt is a free data retrieval call binding the contract method 0x9d949788.
//
// Solidity: function soldAt(bytes32 , bytes32 ) constant returns(bool)
func (_Vendormanagement *VendormanagementCaller) SoldAt(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vendormanagement.contract.Call(opts, out, "soldAt", arg0, arg1)
	return *ret0, err
}

// SoldAt is a free data retrieval call binding the contract method 0x9d949788.
//
// Solidity: function soldAt(bytes32 , bytes32 ) constant returns(bool)
func (_Vendormanagement *VendormanagementSession) SoldAt(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _Vendormanagement.Contract.SoldAt(&_Vendormanagement.CallOpts, arg0, arg1)
}

// SoldAt is a free data retrieval call binding the contract method 0x9d949788.
//
// Solidity: function soldAt(bytes32 , bytes32 ) constant returns(bool)
func (_Vendormanagement *VendormanagementCallerSession) SoldAt(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _Vendormanagement.Contract.SoldAt(&_Vendormanagement.CallOpts, arg0, arg1)
}

// AddProductLocation is a paid mutator transaction binding the contract method 0xc4fe2f93.
//
// Solidity: function addProductLocation(bytes32 _name, bytes32 _location) returns(bool)
func (_Vendormanagement *VendormanagementTransactor) AddProductLocation(opts *bind.TransactOpts, _name [32]byte, _location [32]byte) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "addProductLocation", _name, _location)
}

// AddProductLocation is a paid mutator transaction binding the contract method 0xc4fe2f93.
//
// Solidity: function addProductLocation(bytes32 _name, bytes32 _location) returns(bool)
func (_Vendormanagement *VendormanagementSession) AddProductLocation(_name [32]byte, _location [32]byte) (*types.Transaction, error) {
	return _Vendormanagement.Contract.AddProductLocation(&_Vendormanagement.TransactOpts, _name, _location)
}

// AddProductLocation is a paid mutator transaction binding the contract method 0xc4fe2f93.
//
// Solidity: function addProductLocation(bytes32 _name, bytes32 _location) returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) AddProductLocation(_name [32]byte, _location [32]byte) (*types.Transaction, error) {
	return _Vendormanagement.Contract.AddProductLocation(&_Vendormanagement.TransactOpts, _name, _location)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0xaa144718.
//
// Solidity: function registerProduct(bytes32 _name, bytes32[] _locations, uint256 _cost) returns(bool)
func (_Vendormanagement *VendormanagementTransactor) RegisterProduct(opts *bind.TransactOpts, _name [32]byte, _locations [][32]byte, _cost *big.Int) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "registerProduct", _name, _locations, _cost)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0xaa144718.
//
// Solidity: function registerProduct(bytes32 _name, bytes32[] _locations, uint256 _cost) returns(bool)
func (_Vendormanagement *VendormanagementSession) RegisterProduct(_name [32]byte, _locations [][32]byte, _cost *big.Int) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RegisterProduct(&_Vendormanagement.TransactOpts, _name, _locations, _cost)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0xaa144718.
//
// Solidity: function registerProduct(bytes32 _name, bytes32[] _locations, uint256 _cost) returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) RegisterProduct(_name [32]byte, _locations [][32]byte, _cost *big.Int) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RegisterProduct(&_Vendormanagement.TransactOpts, _name, _locations, _cost)
}

// RemoveProductLocation is a paid mutator transaction binding the contract method 0x310f7b91.
//
// Solidity: function removeProductLocation(bytes32 _name, bytes32 _location) returns(bool)
func (_Vendormanagement *VendormanagementTransactor) RemoveProductLocation(opts *bind.TransactOpts, _name [32]byte, _location [32]byte) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "removeProductLocation", _name, _location)
}

// RemoveProductLocation is a paid mutator transaction binding the contract method 0x310f7b91.
//
// Solidity: function removeProductLocation(bytes32 _name, bytes32 _location) returns(bool)
func (_Vendormanagement *VendormanagementSession) RemoveProductLocation(_name [32]byte, _location [32]byte) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RemoveProductLocation(&_Vendormanagement.TransactOpts, _name, _location)
}

// RemoveProductLocation is a paid mutator transaction binding the contract method 0x310f7b91.
//
// Solidity: function removeProductLocation(bytes32 _name, bytes32 _location) returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) RemoveProductLocation(_name [32]byte, _location [32]byte) (*types.Transaction, error) {
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
	Name     [32]byte
	Location [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProductLocationAdded is a free log retrieval operation binding the contract event 0xc16ed24d108fc30d7daeb46fc1ead0b3251c5a6cd18a0496071802e55c9f8373.
//
// Solidity: event ProductLocationAdded(bytes32 _name, bytes32 _location)
func (_Vendormanagement *VendormanagementFilterer) FilterProductLocationAdded(opts *bind.FilterOpts) (*VendormanagementProductLocationAddedIterator, error) {

	logs, sub, err := _Vendormanagement.contract.FilterLogs(opts, "ProductLocationAdded")
	if err != nil {
		return nil, err
	}
	return &VendormanagementProductLocationAddedIterator{contract: _Vendormanagement.contract, event: "ProductLocationAdded", logs: logs, sub: sub}, nil
}

// WatchProductLocationAdded is a free log subscription operation binding the contract event 0xc16ed24d108fc30d7daeb46fc1ead0b3251c5a6cd18a0496071802e55c9f8373.
//
// Solidity: event ProductLocationAdded(bytes32 _name, bytes32 _location)
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

// ParseProductLocationAdded is a log parse operation binding the contract event 0xc16ed24d108fc30d7daeb46fc1ead0b3251c5a6cd18a0496071802e55c9f8373.
//
// Solidity: event ProductLocationAdded(bytes32 _name, bytes32 _location)
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
	Name     [32]byte
	Location [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProductLocationRemoved is a free log retrieval operation binding the contract event 0xa16e19a51eb50f485a1110b0c514e093a9a1676f77216aa019505467af0bee48.
//
// Solidity: event ProductLocationRemoved(bytes32 _name, bytes32 _location)
func (_Vendormanagement *VendormanagementFilterer) FilterProductLocationRemoved(opts *bind.FilterOpts) (*VendormanagementProductLocationRemovedIterator, error) {

	logs, sub, err := _Vendormanagement.contract.FilterLogs(opts, "ProductLocationRemoved")
	if err != nil {
		return nil, err
	}
	return &VendormanagementProductLocationRemovedIterator{contract: _Vendormanagement.contract, event: "ProductLocationRemoved", logs: logs, sub: sub}, nil
}

// WatchProductLocationRemoved is a free log subscription operation binding the contract event 0xa16e19a51eb50f485a1110b0c514e093a9a1676f77216aa019505467af0bee48.
//
// Solidity: event ProductLocationRemoved(bytes32 _name, bytes32 _location)
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

// ParseProductLocationRemoved is a log parse operation binding the contract event 0xa16e19a51eb50f485a1110b0c514e093a9a1676f77216aa019505467af0bee48.
//
// Solidity: event ProductLocationRemoved(bytes32 _name, bytes32 _location)
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
	Name      [32]byte
	Locations [][32]byte
	Cost      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProductRegistered is a free log retrieval operation binding the contract event 0x86a0e56ff60b92d5010a66e5f314a5433028ab1db0ef9bdcdb19db16b35654f1.
//
// Solidity: event ProductRegistered(bytes32 _name, bytes32[] _locations, uint256 _cost)
func (_Vendormanagement *VendormanagementFilterer) FilterProductRegistered(opts *bind.FilterOpts) (*VendormanagementProductRegisteredIterator, error) {

	logs, sub, err := _Vendormanagement.contract.FilterLogs(opts, "ProductRegistered")
	if err != nil {
		return nil, err
	}
	return &VendormanagementProductRegisteredIterator{contract: _Vendormanagement.contract, event: "ProductRegistered", logs: logs, sub: sub}, nil
}

// WatchProductRegistered is a free log subscription operation binding the contract event 0x86a0e56ff60b92d5010a66e5f314a5433028ab1db0ef9bdcdb19db16b35654f1.
//
// Solidity: event ProductRegistered(bytes32 _name, bytes32[] _locations, uint256 _cost)
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

// ParseProductRegistered is a log parse operation binding the contract event 0x86a0e56ff60b92d5010a66e5f314a5433028ab1db0ef9bdcdb19db16b35654f1.
//
// Solidity: event ProductRegistered(bytes32 _name, bytes32[] _locations, uint256 _cost)
func (_Vendormanagement *VendormanagementFilterer) ParseProductRegistered(log types.Log) (*VendormanagementProductRegistered, error) {
	event := new(VendormanagementProductRegistered)
	if err := _Vendormanagement.contract.UnpackLog(event, "ProductRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
