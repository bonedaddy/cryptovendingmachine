// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vendorfactory

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

// VendorfactoryABI is the input ABI used to generate the binding from.
const VendorfactoryABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vendorContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vendorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newVendor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredVendors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VendorfactoryBin is the compiled bytecode used for deploying new contracts.
var VendorfactoryBin = "0x608060405234801561001057600080fd5b50610e7b806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631e6765c71461005157806340b90f4b1461008a578063c27f2112146100b0578063e7b8c718146100cc575b600080fd5b61006e6004803603602081101561006757600080fd5b50356100f2565b604080516001600160a01b039092168252519081900360200190f35b61006e600480360360208110156100a057600080fd5b50356001600160a01b0316610119565b6100b8610134565b604080519115158252519081900360200190f35b6100b8600480360360208110156100e257600080fd5b50356001600160a01b031661023a565b600081815481106100ff57fe5b6000918252602090912001546001600160a01b0316905081565b6002602052600090815260409020546001600160a01b031681565b3360009081526001602052604081205460ff1615610192576040805162461bcd60e51b81526020600482015260166024820152751b5d5cdd081b9bdd081899481c9959da5cdd195c995960521b604482015290519081900360640190fd5b60006040516101a09061024f565b604051809103906000f0801580156101bc573d6000803e3d6000fd5b5060008054600181810183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390910180546001600160a01b039094166001600160a01b0319948516811790915533835260208281526040808520805460ff1916851790556002909152909220805490931690911790915591505090565b60016020526000908152604090205460ff1681565b610bea8061025d8339019056fe608060405234801561001057600080fd5b5032604051602001610022919061006b565b60408051601f198184030181529190528051602090910120600155600080546001600160a01b031916321790556100af565b61006561006082610080565b61009d565b82525050565b60006100778284610054565b50601401919050565b600061008b82610091565b92915050565b6001600160a01b031690565b600061008b82600061008b8260601b90565b610b2c806100be6000396000f3fe6080604052600436106100705760003560e01c80636c8e745c1161004e5780636c8e745c146101315780638da5cb5b14610151578063af640d0f14610173578063c43df6aa1461019557610070565b80630186a423146100ad578063338a6d10146100e45780635d85ed1314610111575b600080546040516001600160a01b03909116913480156108fc02929091818181858888f193505050501580156100aa573d6000803e3d6000fd5b50005b3480156100b957600080fd5b506100cd6100c83660046106cc565b6101b5565b6040516100db9291906109b8565b60405180910390f35b3480156100f057600080fd5b506101046100ff366004610785565b610263565b6040516100db9190610943565b34801561011d57600080fd5b5061010461012c366004610785565b6102a0565b34801561013d57600080fd5b5061010461014c366004610785565b610353565b34801561015d57600080fd5b506101666103f7565b6040516100db9190610935565b34801561017f57600080fd5b50610188610406565b6040516100db9190610951565b3480156101a157600080fd5b506101046101b0366004610709565b61040c565b805180820160209081018051600280835293830194830194909420939052825460408051601f6000196101006001861615020190931694909404918201839004830284018301905280835283918301828280156102535780601f1061022857610100808354040283529160200191610253565b820191906000526020600020905b81548152906001019060200180831161023657829003601f168201915b5050505050908060010154905082565b8151602081840181018051600382529282019482019490942091909352815180830184018051928152908401929093019190912091525460ff1681565b60006102aa61054b565b6102cf5760405162461bcd60e51b81526004016102c6906109d8565b60405180910390fd5b6003836040516102df9190610929565b9081526020016040518091039020826040516102fb9190610929565b908152604051908190036020018120805460ff191690557f6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f906103419085908590610993565b60405180910390a15060015b92915050565b600061035d61054b565b6103795760405162461bcd60e51b81526004016102c6906109d8565b600160038460405161038b9190610929565b9081526020016040518091039020836040516103a79190610929565b908152604051908190036020018120805492151560ff19909316929092179091557f20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc906103419085908590610993565b6000546001600160a01b031681565b60015481565b600061041661054b565b6104325760405162461bcd60e51b81526004016102c6906109d8565b6040518060400160405280858152602001838152506002856040516104579190610929565b9081526020016040518091039020600082015181600001908051906020019061048192919061056e565b506020919091015160019091015560005b83518110156105045760016003866040516104ad9190610929565b90815260200160405180910390208583815181106104c757fe5b60200260200101516040516104dc9190610929565b908152604051908190036020019020805491151560ff19909216919091179055600101610492565b507f3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d8484846040516105389392919061095f565b60405180910390a15060015b9392505050565b600080546001600160a01b03163314156105675750600161056b565b5060005b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105af57805160ff19168380011785556105dc565b828001600101855582156105dc579182015b828111156105dc5782518255916020019190600101906105c1565b506105e89291506105ec565b5090565b61056b91905b808211156105e857600081556001016105f2565b600082601f83011261061757600080fd5b813561062a61062582610a0f565b6109e8565b81815260209384019390925082018360005b8381101561066857813586016106528882610672565b845250602092830192919091019060010161063c565b5050505092915050565b600082601f83011261068357600080fd5b813561069161062582610a30565b915080825260208301602083018583830111156106ad57600080fd5b6106b8838284610a8c565b50505092915050565b803561034d81610ad2565b6000602082840312156106de57600080fd5b813567ffffffffffffffff8111156106f557600080fd5b61070184828501610672565b949350505050565b60008060006060848603121561071e57600080fd5b833567ffffffffffffffff81111561073557600080fd5b61074186828701610672565b935050602084013567ffffffffffffffff81111561075e57600080fd5b61076a86828701610606565b925050604061077b868287016106c1565b9150509250925092565b6000806040838503121561079857600080fd5b823567ffffffffffffffff8111156107af57600080fd5b6107bb85828601610672565b925050602083013567ffffffffffffffff8111156107d857600080fd5b6107e485828601610672565b9150509250929050565b60006105448383610889565b61080381610a70565b82525050565b600061081482610a5e565b61081e8185610a62565b93508360208202850161083085610a58565b8060005b8581101561086a578484038952815161084d85826107ee565b945061085883610a58565b60209a909a0199925050600101610834565b5091979650505050505050565b61080381610a7b565b6108038161056b565b600061089482610a5e565b61089e8185610a62565b93506108ae818560208601610a98565b6108b781610ac8565b9093019392505050565b60006108cc82610a5e565b6108d68185610a6b565b93506108e6818560208601610a98565b9290920192915050565b60006108fd601783610a62565b7f63616c6c6572206d7573742062652076656e646f726564000000000000000000815260200192915050565b600061054482846108c1565b6020810161034d82846107fa565b6020810161034d8284610877565b6020810161034d8284610880565b606080825281016109708186610889565b905081810360208301526109848185610809565b90506107016040830184610880565b604080825281016109a48185610889565b905081810360208301526107018184610889565b604080825281016109c98185610889565b90506105446020830184610880565b6020808252810161034d816108f0565b60405181810167ffffffffffffffff81118282101715610a0757600080fd5b604052919050565b600067ffffffffffffffff821115610a2657600080fd5b5060209081020190565b600067ffffffffffffffff821115610a4757600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b600061034d82610a80565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015610ab3578181015183820152602001610a9b565b83811115610ac2576000848401525b50505050565b601f01601f191690565b610adb8161056b565b8114610ae657600080fd5b5056fea365627a7a7231582068574bf963d1750dd69316317ebed21e0ff0999b71de9e499cb4ad2108e32d726c6578706572696d656e74616cf564736f6c634300050b0040a265627a7a7231582077da1c9a4610384b5775bce9ad2a4940b74ab4d52f4e1703d2468aa4c6e3663864736f6c634300050b0032"

// DeployVendorfactory deploys a new Ethereum contract, binding an instance of Vendorfactory to it.
func DeployVendorfactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vendorfactory, error) {
	parsed, err := abi.JSON(strings.NewReader(VendorfactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VendorfactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vendorfactory{VendorfactoryCaller: VendorfactoryCaller{contract: contract}, VendorfactoryTransactor: VendorfactoryTransactor{contract: contract}, VendorfactoryFilterer: VendorfactoryFilterer{contract: contract}}, nil
}

// Vendorfactory is an auto generated Go binding around an Ethereum contract.
type Vendorfactory struct {
	VendorfactoryCaller     // Read-only binding to the contract
	VendorfactoryTransactor // Write-only binding to the contract
	VendorfactoryFilterer   // Log filterer for contract events
}

// VendorfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VendorfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VendorfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VendorfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VendorfactorySession struct {
	Contract     *Vendorfactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VendorfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VendorfactoryCallerSession struct {
	Contract *VendorfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VendorfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VendorfactoryTransactorSession struct {
	Contract     *VendorfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VendorfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VendorfactoryRaw struct {
	Contract *Vendorfactory // Generic contract binding to access the raw methods on
}

// VendorfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VendorfactoryCallerRaw struct {
	Contract *VendorfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VendorfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VendorfactoryTransactorRaw struct {
	Contract *VendorfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVendorfactory creates a new instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactory(address common.Address, backend bind.ContractBackend) (*Vendorfactory, error) {
	contract, err := bindVendorfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vendorfactory{VendorfactoryCaller: VendorfactoryCaller{contract: contract}, VendorfactoryTransactor: VendorfactoryTransactor{contract: contract}, VendorfactoryFilterer: VendorfactoryFilterer{contract: contract}}, nil
}

// NewVendorfactoryCaller creates a new read-only instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryCaller(address common.Address, caller bind.ContractCaller) (*VendorfactoryCaller, error) {
	contract, err := bindVendorfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryCaller{contract: contract}, nil
}

// NewVendorfactoryTransactor creates a new write-only instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VendorfactoryTransactor, error) {
	contract, err := bindVendorfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryTransactor{contract: contract}, nil
}

// NewVendorfactoryFilterer creates a new log filterer instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VendorfactoryFilterer, error) {
	contract, err := bindVendorfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryFilterer{contract: contract}, nil
}

// bindVendorfactory binds a generic wrapper to an already deployed contract.
func bindVendorfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VendorfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendorfactory *VendorfactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendorfactory.Contract.VendorfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendorfactory *VendorfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.Contract.VendorfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendorfactory *VendorfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendorfactory.Contract.VendorfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendorfactory *VendorfactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendorfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendorfactory *VendorfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendorfactory *VendorfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendorfactory.Contract.contract.Transact(opts, method, params...)
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactoryCaller) RegisteredVendors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "registeredVendors", arg0)
	return *ret0, err
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactorySession) RegisteredVendors(arg0 common.Address) (bool, error) {
	return _Vendorfactory.Contract.RegisteredVendors(&_Vendorfactory.CallOpts, arg0)
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactoryCallerSession) RegisteredVendors(arg0 common.Address) (bool, error) {
	return _Vendorfactory.Contract.RegisteredVendors(&_Vendorfactory.CallOpts, arg0)
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactoryCaller) VendorContract(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "vendorContract", arg0)
	return *ret0, err
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactorySession) VendorContract(arg0 common.Address) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContract(&_Vendorfactory.CallOpts, arg0)
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactoryCallerSession) VendorContract(arg0 common.Address) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContract(&_Vendorfactory.CallOpts, arg0)
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactoryCaller) VendorContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "vendorContracts", arg0)
	return *ret0, err
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactorySession) VendorContracts(arg0 *big.Int) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContracts(&_Vendorfactory.CallOpts, arg0)
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactoryCallerSession) VendorContracts(arg0 *big.Int) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContracts(&_Vendorfactory.CallOpts, arg0)
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactoryTransactor) NewVendor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.contract.Transact(opts, "newVendor")
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactorySession) NewVendor() (*types.Transaction, error) {
	return _Vendorfactory.Contract.NewVendor(&_Vendorfactory.TransactOpts)
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactoryTransactorSession) NewVendor() (*types.Transaction, error) {
	return _Vendorfactory.Contract.NewVendor(&_Vendorfactory.TransactOpts)
}
