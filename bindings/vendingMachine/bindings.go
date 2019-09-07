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
const VendingmachineABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"backend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_vendor\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_product\",\"type\":\"string\"}],\"name\":\"backendPurchaseProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vendors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"vendorNames\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vendorContract\",\"type\":\"address\"}],\"name\":\"addVendor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"location\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_vendor\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_product\",\"type\":\"string\"}],\"name\":\"purchaseProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vendorContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_backend\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_vendor\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_product\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"ProductPurchased\",\"type\":\"event\"}]"

// VendingmachineBin is the compiled bytecode used for deploying new contracts.
var VendingmachineBin = "0x60806040523480156200001157600080fd5b506040516200143d3803806200143d833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b5060405260209081015184519093506200010792506000918501906200012f565b50600180546001600160a01b0319166001600160a01b039290921691909117905550620001d4565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200017257805160ff1916838001178555620001a2565b82800160010185558215620001a2579182015b82811115620001a257825182559160200191906001019062000185565b50620001b0929150620001b4565b5090565b620001d191905b80821115620001b05760008155600101620001bb565b90565b61125980620001e46000396000f3fe60806040526004361061007b5760003560e01c80634fc611981161004e5780634fc61198146102df578063516f279e1461039b578063a0622f4e14610425578063a46ead241461054e5761007b565b8063099e4133146100805780631d3bc57b146100b157806327e30c36146101fb5780632c62204c1461022e575b600080fd5b34801561008c57600080fd5b50610095610581565b604080516001600160a01b039092168252519081900360200190f35b3480156100bd57600080fd5b506101e7600480360360408110156100d457600080fd5b810190602081018135600160201b8111156100ee57600080fd5b82018360208201111561010057600080fd5b803590602001918460018302840111600160201b8311171561012157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561017357600080fd5b82018360208201111561018557600080fd5b803590602001918460018302840111600160201b831117156101a657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610590945050505050565b604080519115158252519081900360200190f35b34801561020757600080fd5b506101e76004803603602081101561021e57600080fd5b50356001600160a01b031661074e565b34801561023a57600080fd5b506100956004803603602081101561025157600080fd5b810190602081018135600160201b81111561026b57600080fd5b82018360208201111561027d57600080fd5b803590602001918460018302840111600160201b8311171561029e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610763945050505050565b3480156102eb57600080fd5b506101e76004803603604081101561030257600080fd5b810190602081018135600160201b81111561031c57600080fd5b82018360208201111561032e57600080fd5b803590602001918460018302840111600160201b8311171561034f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506107899050565b3480156103a757600080fd5b506103b0610a4d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103ea5781810151838201526020016103d2565b50505050905090810190601f1680156104175780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e76004803603604081101561043b57600080fd5b810190602081018135600160201b81111561045557600080fd5b82018360208201111561046757600080fd5b803590602001918460018302840111600160201b8311171561048857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156104da57600080fd5b8201836020820111156104ec57600080fd5b803590602001918460018302840111600160201b8311171561050d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610adb945050505050565b34801561055a57600080fd5b506100956004803603602081101561057157600080fd5b50356001600160a01b0316610f27565b6001546001600160a01b031681565b6001546000906001600160a01b031633146105eb576040805162461bcd60e51b815260206004820152601660248201527518d85b1b195c881b5d5cdd08189948189858dad95b9960521b604482015290519081900360640190fd5b6105f58383610f42565b61063d576040805162461bcd60e51b815260206004820152601460248201527370726f64756374206e6f7420666f722073616c6560601b604482015290519081900360640190fd5b7f3bdb8608a438ca1dc20d36648e2437b914016bb217e835254bcde93444cf3b95838342604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156106a8578181015183820152602001610690565b50505050905090810190601f1680156106d55780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156107085781810151838201526020016106f0565b50505050905090810190601f1680156107355780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a150600192915050565b60026020526000908152604090205460ff1681565b80516020818301810180516004825292820191909301209152546001600160a01b031681565b3360009081526002602052604081205460ff16156107ee576040805162461bcd60e51b815260206004820152601d60248201527f76656e646f72206d757374206e6f742062652072656769737465726564000000604482015290519081900360640190fd5b60006001600160a01b03166004846040518082805190602001908083835b6020831061082b5780518252601f19909201916020918201910161080c565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b03169290921491506108b29050576040805162461bcd60e51b81526020600482015260166024820152753730b6b29036bab9ba103737ba103132903a30b5b2b760511b604482015290519081900360640190fd5b6000829050336001600160a01b0316816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156108fa57600080fd5b505afa15801561090e573d6000803e3d6000fd5b505050506040513d602081101561092457600080fd5b50516001600160a01b031614610981576040805162461bcd60e51b815260206004820152601d60248201527f63616c6c6572206d7573742062652076656e646f72206d616e61676572000000604482015290519081900360640190fd5b336000908152600260209081526040808320805460ff19166001179055600382529182902080546001600160a01b0387166001600160a01b0319909116179055905185518592600492889290918291908401908083835b602083106109f75780518252601f1990920191602091820191016109d8565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080546001600160a01b0319166001600160a01b03949094169390931790925550600195945050505050565b6000805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610ad35780601f10610aa857610100808354040283529160200191610ad3565b820191906000526020600020905b815481529060010190602001808311610ab657829003601f168201915b505050505081565b6000610ae78383610f42565b610b2f576040805162461bcd60e51b815260206004820152601460248201527370726f64756374206e6f7420666f722073616c6560601b604482015290519081900360640190fd5b60006004846040518082805190602001908083835b60208310610b635780518252601f199092019160209182019101610b44565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101842054630186a42360e01b8552600485018281528851602487015288516001600160a01b039092169650600095879550630186a42394508993919283926044909201918501908083838b5b83811015610bef578181015183820152602001610bd7565b50505050905090810190601f168015610c1c5780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b158015610c3957600080fd5b505afa158015610c4d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015610c7657600080fd5b8101908080516040519392919084600160201b821115610c9557600080fd5b908301906020820185811115610caa57600080fd5b8251600160201b811182820188101715610cc357600080fd5b82525081516020918201929091019080838360005b83811015610cf0578181015183820152602001610cd8565b50505050905090810190601f168015610d1d5780820380516001836020036101000a031916815260200191505b50604052602001519350503483149150610d809050576040805162461bcd60e51b815260206004820152601860248201527f696e636f7272656374207061796d656e7420616d6f756e740000000000000000604482015290519081900360640190fd5b6004856040518082805190602001908083835b60208310610db25780518252601f199092019160209182019101610d93565b51815160001960209485036101000a0190811690199190911617905292019485525060405193849003018320546001600160a01b0316923480156108fc02935091506000818181858888f19350505050158015610e13573d6000803e3d6000fd5b507f3bdb8608a438ca1dc20d36648e2437b914016bb217e835254bcde93444cf3b95858542604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015610e7f578181015183820152602001610e67565b50505050905090810190601f168015610eac5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610edf578181015183820152602001610ec7565b50505050905090810190601f168015610f0c5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1506001949350505050565b6003602052600090815260409020546001600160a01b031681565b6000806001600160a01b03166004846040518082805190602001908083835b60208310610f805780518252601f199092019160209182019101610f61565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b031692909214159150610ff990505760405162461bcd60e51b81526004018080602001828103825260228152602001806112036022913960400191505060405180910390fd5b60006004846040518082805190602001908083835b6020831061102d5780518252601f19909201916020918201910161100e565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201852054630338a6d160e41b8652600486019182528851604487015288516001600160a01b039091169650869563338a6d1095508994506000938392602483019260640191870190808383895b838110156110bb5781810151838201526020016110a3565b50505050905090810190601f1680156110e85780820380516001836020036101000a031916815260200191505b5083810382528454600260001961010060018416150201909116048082526020909101908590801561115b5780601f106111305761010080835404028352916020019161115b565b820191906000526020600020905b81548152906001019060200180831161113e57829003601f168201915b505094505050505060206040518083038186803b15801561117b57600080fd5b505afa15801561118f573d6000803e3d6000fd5b505050506040513d60208110156111a557600080fd5b50516111f8576040805162461bcd60e51b815260206004820152601f60248201527f70726f64756374206e6f7420666f722073616c65206174206d616368696e6500604482015290519081900360640190fd5b506001939250505056fe76656e646f72206e6f7420726567697374657265642077697468206d616368696e65a265627a7a7231582087b95483811e89a45313d771dd5e285d86f43bbb8352aa2953b2937a7ac4bbed64736f6c634300050b0032"

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

// VendorNames is a free data retrieval call binding the contract method 0x2c62204c.
//
// Solidity: function vendorNames(string ) constant returns(address)
func (_Vendingmachine *VendingmachineCaller) VendorNames(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendingmachine.contract.Call(opts, out, "vendorNames", arg0)
	return *ret0, err
}

// VendorNames is a free data retrieval call binding the contract method 0x2c62204c.
//
// Solidity: function vendorNames(string ) constant returns(address)
func (_Vendingmachine *VendingmachineSession) VendorNames(arg0 string) (common.Address, error) {
	return _Vendingmachine.Contract.VendorNames(&_Vendingmachine.CallOpts, arg0)
}

// VendorNames is a free data retrieval call binding the contract method 0x2c62204c.
//
// Solidity: function vendorNames(string ) constant returns(address)
func (_Vendingmachine *VendingmachineCallerSession) VendorNames(arg0 string) (common.Address, error) {
	return _Vendingmachine.Contract.VendorNames(&_Vendingmachine.CallOpts, arg0)
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

// AddVendor is a paid mutator transaction binding the contract method 0x4fc61198.
//
// Solidity: function addVendor(string _name, address _vendorContract) returns(bool)
func (_Vendingmachine *VendingmachineTransactor) AddVendor(opts *bind.TransactOpts, _name string, _vendorContract common.Address) (*types.Transaction, error) {
	return _Vendingmachine.contract.Transact(opts, "addVendor", _name, _vendorContract)
}

// AddVendor is a paid mutator transaction binding the contract method 0x4fc61198.
//
// Solidity: function addVendor(string _name, address _vendorContract) returns(bool)
func (_Vendingmachine *VendingmachineSession) AddVendor(_name string, _vendorContract common.Address) (*types.Transaction, error) {
	return _Vendingmachine.Contract.AddVendor(&_Vendingmachine.TransactOpts, _name, _vendorContract)
}

// AddVendor is a paid mutator transaction binding the contract method 0x4fc61198.
//
// Solidity: function addVendor(string _name, address _vendorContract) returns(bool)
func (_Vendingmachine *VendingmachineTransactorSession) AddVendor(_name string, _vendorContract common.Address) (*types.Transaction, error) {
	return _Vendingmachine.Contract.AddVendor(&_Vendingmachine.TransactOpts, _name, _vendorContract)
}

// BackendPurchaseProduct is a paid mutator transaction binding the contract method 0x1d3bc57b.
//
// Solidity: function backendPurchaseProduct(string _vendor, string _product) returns(bool)
func (_Vendingmachine *VendingmachineTransactor) BackendPurchaseProduct(opts *bind.TransactOpts, _vendor string, _product string) (*types.Transaction, error) {
	return _Vendingmachine.contract.Transact(opts, "backendPurchaseProduct", _vendor, _product)
}

// BackendPurchaseProduct is a paid mutator transaction binding the contract method 0x1d3bc57b.
//
// Solidity: function backendPurchaseProduct(string _vendor, string _product) returns(bool)
func (_Vendingmachine *VendingmachineSession) BackendPurchaseProduct(_vendor string, _product string) (*types.Transaction, error) {
	return _Vendingmachine.Contract.BackendPurchaseProduct(&_Vendingmachine.TransactOpts, _vendor, _product)
}

// BackendPurchaseProduct is a paid mutator transaction binding the contract method 0x1d3bc57b.
//
// Solidity: function backendPurchaseProduct(string _vendor, string _product) returns(bool)
func (_Vendingmachine *VendingmachineTransactorSession) BackendPurchaseProduct(_vendor string, _product string) (*types.Transaction, error) {
	return _Vendingmachine.Contract.BackendPurchaseProduct(&_Vendingmachine.TransactOpts, _vendor, _product)
}

// PurchaseProduct is a paid mutator transaction binding the contract method 0xa0622f4e.
//
// Solidity: function purchaseProduct(string _vendor, string _product) returns(bool)
func (_Vendingmachine *VendingmachineTransactor) PurchaseProduct(opts *bind.TransactOpts, _vendor string, _product string) (*types.Transaction, error) {
	return _Vendingmachine.contract.Transact(opts, "purchaseProduct", _vendor, _product)
}

// PurchaseProduct is a paid mutator transaction binding the contract method 0xa0622f4e.
//
// Solidity: function purchaseProduct(string _vendor, string _product) returns(bool)
func (_Vendingmachine *VendingmachineSession) PurchaseProduct(_vendor string, _product string) (*types.Transaction, error) {
	return _Vendingmachine.Contract.PurchaseProduct(&_Vendingmachine.TransactOpts, _vendor, _product)
}

// PurchaseProduct is a paid mutator transaction binding the contract method 0xa0622f4e.
//
// Solidity: function purchaseProduct(string _vendor, string _product) returns(bool)
func (_Vendingmachine *VendingmachineTransactorSession) PurchaseProduct(_vendor string, _product string) (*types.Transaction, error) {
	return _Vendingmachine.Contract.PurchaseProduct(&_Vendingmachine.TransactOpts, _vendor, _product)
}

// VendingmachineProductPurchasedIterator is returned from FilterProductPurchased and is used to iterate over the raw logs and unpacked data for ProductPurchased events raised by the Vendingmachine contract.
type VendingmachineProductPurchasedIterator struct {
	Event *VendingmachineProductPurchased // Event containing the contract specifics and raw log

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
func (it *VendingmachineProductPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VendingmachineProductPurchased)
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
		it.Event = new(VendingmachineProductPurchased)
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
func (it *VendingmachineProductPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VendingmachineProductPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VendingmachineProductPurchased represents a ProductPurchased event raised by the Vendingmachine contract.
type VendingmachineProductPurchased struct {
	Vendor    string
	Product   string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProductPurchased is a free log retrieval operation binding the contract event 0x3bdb8608a438ca1dc20d36648e2437b914016bb217e835254bcde93444cf3b95.
//
// Solidity: event ProductPurchased(string _vendor, string _product, uint256 _timestamp)
func (_Vendingmachine *VendingmachineFilterer) FilterProductPurchased(opts *bind.FilterOpts) (*VendingmachineProductPurchasedIterator, error) {

	logs, sub, err := _Vendingmachine.contract.FilterLogs(opts, "ProductPurchased")
	if err != nil {
		return nil, err
	}
	return &VendingmachineProductPurchasedIterator{contract: _Vendingmachine.contract, event: "ProductPurchased", logs: logs, sub: sub}, nil
}

// WatchProductPurchased is a free log subscription operation binding the contract event 0x3bdb8608a438ca1dc20d36648e2437b914016bb217e835254bcde93444cf3b95.
//
// Solidity: event ProductPurchased(string _vendor, string _product, uint256 _timestamp)
func (_Vendingmachine *VendingmachineFilterer) WatchProductPurchased(opts *bind.WatchOpts, sink chan<- *VendingmachineProductPurchased) (event.Subscription, error) {

	logs, sub, err := _Vendingmachine.contract.WatchLogs(opts, "ProductPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VendingmachineProductPurchased)
				if err := _Vendingmachine.contract.UnpackLog(event, "ProductPurchased", log); err != nil {
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

// ParseProductPurchased is a log parse operation binding the contract event 0x3bdb8608a438ca1dc20d36648e2437b914016bb217e835254bcde93444cf3b95.
//
// Solidity: event ProductPurchased(string _vendor, string _product, uint256 _timestamp)
func (_Vendingmachine *VendingmachineFilterer) ParseProductPurchased(log types.Log) (*VendingmachineProductPurchased, error) {
	event := new(VendingmachineProductPurchased)
	if err := _Vendingmachine.contract.UnpackLog(event, "ProductPurchased", log); err != nil {
		return nil, err
	}
	return event, nil
}
