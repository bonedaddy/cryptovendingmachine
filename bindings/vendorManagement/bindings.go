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
const VendormanagementABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_location\",\"type\":\"bytes32\"}],\"name\":\"removeProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_locations\",\"type\":\"bytes32[]\"}],\"name\":\"registerProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"vendors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumVendorManagement.State\",\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_location\",\"type\":\"bytes32\"}],\"name\":\"addProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerVendor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vendor\",\"type\":\"address\"}],\"name\":\"VendorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"_locations\",\"type\":\"bytes32[]\"}],\"name\":\"ProductRegistered\",\"type\":\"event\"}]"

// VendormanagementBin is the compiled bytecode used for deploying new contracts.
var VendormanagementBin = "0x608060405234801561001057600080fd5b506105c6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063310f7b911461005c5780639bee671614610093578063a61a809c1461013d578063c4fe2f931461018e578063e66f3b4f146101b1575b600080fd5b61007f6004803603604081101561007257600080fd5b50803590602001356101b9565b604080519115158252519081900360200190f35b61007f600480360360408110156100a957600080fd5b813591908101906040810160208201356401000000008111156100cb57600080fd5b8201836020820111156100dd57600080fd5b803590602001918460208302840111640100000000831117156100ff57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610250945050505050565b61015a6004803603602081101561015357600080fd5b503561034d565b6040516001600160a01b03831681526020810182600181111561017957fe5b60ff1681526020019250505060405180910390f35b61007f600480360360408110156101a457600080fd5b5080359060200135610374565b61007f61040d565b6000806101c461051e565b90506101cf81610549565b61021c576040805162461bcd60e51b81526020600482015260196024820152781d995b991bdc881b5d5cdd081899481c9959da5cdd195c9959603a1b604482015290519081900360640190fd5b600090815260016020818152604080842087855282528084208685528301909152909120805460ff19169055905092915050565b60008061025b61051e565b905061026681610549565b6102b3576040805162461bcd60e51b81526020600482015260196024820152781d995b991bdc881b5d5cdd081899481c9959da5cdd195c9959603a1b604482015290519081900360640190fd5b6040805160208082018352868252600084815260018252838120888252909152918220905190555b83518110156103425760008281526001602081815260408084208985529091528220865191929083019187908590811061031157fe5b6020908102919091018101518252810191909152604001600020805460ff19169115159190911790556001016102db565b506001949350505050565b6000602081905290815260409020546001600160a01b03811690600160a01b900460ff1682565b60008061037f61051e565b905061038a81610549565b6103d7576040805162461bcd60e51b81526020600482015260196024820152781d995b991bdc881b5d5cdd081899481c9959da5cdd195c9959603a1b604482015290519081900360640190fd5b600090815260016020818152604080842087855282528084208685528301909152909120805460ff191682179055905092915050565b60008061041861051e565b905061042381610589565b610474576040805162461bcd60e51b815260206004820152601d60248201527f76656e646f72206d757374206e6f742062652072656967737465726564000000604482015290519081900360640190fd5b604080518082019091523381526020810160019052600082815260208181526040909120825181546001600160a01b0319166001600160a01b0390911617808255918301519091829060ff60a01b1916600160a01b8360018111156104d557fe5b0217905550506040805183815233602082015281517f5aaceb50d5042e96b113fa728a9e3471cbe751a187e513a021ef75f68003c75c93509081900390910190a1600191505090565b604080513360601b602080830191909152825180830360140181526034909201909252805191012090565b600060015b600083815260208190526040902054600160a01b900460ff16600181111561057257fe5b141561058057506001610584565b5060005b919050565b60008061054e56fea265627a7a723158204f6de89cbe91e9cf81ecbf799202f2a177303824210d168f10c06aaf7e133f3564736f6c634300050b0032"

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

// Vendors is a free data retrieval call binding the contract method 0xa61a809c.
//
// Solidity: function vendors(bytes32 ) constant returns(address owner, uint8 state)
func (_Vendormanagement *VendormanagementCaller) Vendors(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner common.Address
	State uint8
}, error) {
	ret := new(struct {
		Owner common.Address
		State uint8
	})
	out := ret
	err := _Vendormanagement.contract.Call(opts, out, "vendors", arg0)
	return *ret, err
}

// Vendors is a free data retrieval call binding the contract method 0xa61a809c.
//
// Solidity: function vendors(bytes32 ) constant returns(address owner, uint8 state)
func (_Vendormanagement *VendormanagementSession) Vendors(arg0 [32]byte) (struct {
	Owner common.Address
	State uint8
}, error) {
	return _Vendormanagement.Contract.Vendors(&_Vendormanagement.CallOpts, arg0)
}

// Vendors is a free data retrieval call binding the contract method 0xa61a809c.
//
// Solidity: function vendors(bytes32 ) constant returns(address owner, uint8 state)
func (_Vendormanagement *VendormanagementCallerSession) Vendors(arg0 [32]byte) (struct {
	Owner common.Address
	State uint8
}, error) {
	return _Vendormanagement.Contract.Vendors(&_Vendormanagement.CallOpts, arg0)
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

// RegisterProduct is a paid mutator transaction binding the contract method 0x9bee6716.
//
// Solidity: function registerProduct(bytes32 _name, bytes32[] _locations) returns(bool)
func (_Vendormanagement *VendormanagementTransactor) RegisterProduct(opts *bind.TransactOpts, _name [32]byte, _locations [][32]byte) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "registerProduct", _name, _locations)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0x9bee6716.
//
// Solidity: function registerProduct(bytes32 _name, bytes32[] _locations) returns(bool)
func (_Vendormanagement *VendormanagementSession) RegisterProduct(_name [32]byte, _locations [][32]byte) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RegisterProduct(&_Vendormanagement.TransactOpts, _name, _locations)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0x9bee6716.
//
// Solidity: function registerProduct(bytes32 _name, bytes32[] _locations) returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) RegisterProduct(_name [32]byte, _locations [][32]byte) (*types.Transaction, error) {
	return _Vendormanagement.Contract.RegisterProduct(&_Vendormanagement.TransactOpts, _name, _locations)
}

// RegisterVendor is a paid mutator transaction binding the contract method 0xe66f3b4f.
//
// Solidity: function registerVendor() returns(bool)
func (_Vendormanagement *VendormanagementTransactor) RegisterVendor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendormanagement.contract.Transact(opts, "registerVendor")
}

// RegisterVendor is a paid mutator transaction binding the contract method 0xe66f3b4f.
//
// Solidity: function registerVendor() returns(bool)
func (_Vendormanagement *VendormanagementSession) RegisterVendor() (*types.Transaction, error) {
	return _Vendormanagement.Contract.RegisterVendor(&_Vendormanagement.TransactOpts)
}

// RegisterVendor is a paid mutator transaction binding the contract method 0xe66f3b4f.
//
// Solidity: function registerVendor() returns(bool)
func (_Vendormanagement *VendormanagementTransactorSession) RegisterVendor() (*types.Transaction, error) {
	return _Vendormanagement.Contract.RegisterVendor(&_Vendormanagement.TransactOpts)
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

// VendormanagementVendorRegisteredIterator is returned from FilterVendorRegistered and is used to iterate over the raw logs and unpacked data for VendorRegistered events raised by the Vendormanagement contract.
type VendormanagementVendorRegisteredIterator struct {
	Event *VendormanagementVendorRegistered // Event containing the contract specifics and raw log

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
func (it *VendormanagementVendorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendormanagementVendorRegistered)
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
		it.Event = new(VendormanagementVendorRegistered)
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
func (it *VendormanagementVendorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendormanagementVendorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendormanagementVendorRegistered represents a VendorRegistered event raised by the Vendormanagement contract.
type VendormanagementVendorRegistered struct {
	Id     [32]byte
	Vendor common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVendorRegistered is a free log retrieval operation binding the contract event 0x5aaceb50d5042e96b113fa728a9e3471cbe751a187e513a021ef75f68003c75c.
//
// Solidity: event VendorRegistered(bytes32 _id, address _vendor)
func (_Vendormanagement *VendormanagementFilterer) FilterVendorRegistered(opts *bind.FilterOpts) (*VendormanagementVendorRegisteredIterator, error) {

	logs, sub, err := _Vendormanagement.contract.FilterLogs(opts, "VendorRegistered")
	if err != nil {
		return nil, err
	}
	return &VendormanagementVendorRegisteredIterator{contract: _Vendormanagement.contract, event: "VendorRegistered", logs: logs, sub: sub}, nil
}

// WatchVendorRegistered is a free log subscription operation binding the contract event 0x5aaceb50d5042e96b113fa728a9e3471cbe751a187e513a021ef75f68003c75c.
//
// Solidity: event VendorRegistered(bytes32 _id, address _vendor)
func (_Vendormanagement *VendormanagementFilterer) WatchVendorRegistered(opts *bind.WatchOpts, sink chan<- *VendormanagementVendorRegistered) (event.Subscription, error) {

	logs, sub, err := _Vendormanagement.contract.WatchLogs(opts, "VendorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendormanagementVendorRegistered)
				if err := _Vendormanagement.contract.UnpackLog(event, "VendorRegistered", log); err != nil {
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

// ParseVendorRegistered is a log parse operation binding the contract event 0x5aaceb50d5042e96b113fa728a9e3471cbe751a187e513a021ef75f68003c75c.
//
// Solidity: event VendorRegistered(bytes32 _id, address _vendor)
func (_Vendormanagement *VendormanagementFilterer) ParseVendorRegistered(log types.Log) (*VendormanagementVendorRegistered, error) {
	event := new(VendormanagementVendorRegistered)
	if err := _Vendormanagement.contract.UnpackLog(event, "VendorRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
