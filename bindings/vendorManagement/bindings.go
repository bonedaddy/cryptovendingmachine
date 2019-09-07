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
const VendormanagementABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_location\",\"type\":\"bytes32\"}],\"name\":\"removeProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"soldAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_locations\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"registerProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_location\",\"type\":\"bytes32\"}],\"name\":\"addProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"_locations\",\"type\":\"bytes32[]\"}],\"name\":\"ProductRegistered\",\"type\":\"event\"}]"

// VendormanagementBin is the compiled bytecode used for deploying new contracts.
var VendormanagementBin = "0x608060405234801561001057600080fd5b50604080513360601b602080830191909152825180830360140181526034909201909252805191012060015561040a8061004b6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80639d9497881161005b5780639d94978814610113578063aa14471814610136578063af640d0f146101e2578063c4fe2f93146100825761007d565b8063310f7b911461008257806379054391146100b95780638da5cb5b146100ef575b600080fd5b6100a56004803603604081101561009857600080fd5b50803590602001356101fc565b604080519115158252519081900360200190f35b6100d6600480360360208110156100cf57600080fd5b503561027c565b6040805192835260208301919091528051918290030190f35b6100f7610295565b604080516001600160a01b039092168252519081900360200190f35b6100a56004803603604081101561012957600080fd5b50803590602001356102a4565b6100a56004803603606081101561014c57600080fd5b8135919081019060408101602082013564010000000081111561016e57600080fd5b82018360208201111561018057600080fd5b803590602001918460208302840111640100000000831117156101a257600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506102c4915050565b6101ea6103ac565b60408051918252519081900360200190f35b60006102066103b2565b610251576040805162461bcd60e51b815260206004820152601760248201527618d85b1b195c881b5d5cdd081899481d995b991bdc9959604a1b604482015290519081900360640190fd5b5060009182526003602090815260408084209284529190529020805460ff1916600190811790915590565b6002602052600090815260409020805460019091015482565b6000546001600160a01b031681565b600360209081526000928352604080842090915290825290205460ff1681565b60006102ce6103b2565b610319576040805162461bcd60e51b815260206004820152601760248201527618d85b1b195c881b5d5cdd081899481d995b991bdc9959604a1b604482015290519081900360640190fd5b6040805180820182528581526020808201858152600088815260029092529281209151825591516001909101555b83518110156103a157600085815260036020526040812085516001929087908590811061037057fe5b6020908102919091018101518252810191909152604001600020805460ff1916911515919091179055600101610347565b506001949350505050565b60015481565b600080546001600160a01b03163314156103ce575060016103d2565b5060005b9056fea265627a7a72315820418b8ca00e4f0357d9abec5bd163d9b77fda4927d584da6e8892d4fae79c3e0b64736f6c634300050b0032"

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
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProductRegistered is a free log retrieval operation binding the contract event 0x5557af5bef645f9effaf31614a385de3b8fcb7973118a4e7e28a65e57e163d8a.
//
// Solidity: event ProductRegistered(bytes32 _name, bytes32[] _locations)
func (_Vendormanagement *VendormanagementFilterer) FilterProductRegistered(opts *bind.FilterOpts) (*VendormanagementProductRegisteredIterator, error) {

	logs, sub, err := _Vendormanagement.contract.FilterLogs(opts, "ProductRegistered")
	if err != nil {
		return nil, err
	}
	return &VendormanagementProductRegisteredIterator{contract: _Vendormanagement.contract, event: "ProductRegistered", logs: logs, sub: sub}, nil
}

// WatchProductRegistered is a free log subscription operation binding the contract event 0x5557af5bef645f9effaf31614a385de3b8fcb7973118a4e7e28a65e57e163d8a.
//
// Solidity: event ProductRegistered(bytes32 _name, bytes32[] _locations)
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

// ParseProductRegistered is a log parse operation binding the contract event 0x5557af5bef645f9effaf31614a385de3b8fcb7973118a4e7e28a65e57e163d8a.
//
// Solidity: event ProductRegistered(bytes32 _name, bytes32[] _locations)
func (_Vendormanagement *VendormanagementFilterer) ParseProductRegistered(log types.Log) (*VendormanagementProductRegistered, error) {
	event := new(VendormanagementProductRegistered)
	if err := _Vendormanagement.contract.UnpackLog(event, "ProductRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
