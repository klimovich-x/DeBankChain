// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DebankL2RegisterMetaData contains all meta data concerning the DebankL2Register contract.
var DebankL2RegisterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l2Account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"registerCnt\",\"type\":\"uint256\"}],\"name\":\"Register\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"l2Accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Account\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b5060016080819052600060a081905260c0829052816108db61004a833960006104c80152600061049f0152600061047601526108db6000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c806354fd4d501161005057806354fd4d50146100ff5780637ecebe00146101145780638da5cb5b1461014257600080fd5b806313af4035146100775780632a71d08d1461008c5780634420e486146100ec575b600080fd5b61008a61008536600461064f565b610162565b005b6100c261009a36600461064f565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61008a6100fa36600461064f565b610398565b61010761046f565b6040516100e391906106bc565b61013461012236600461064f565b60016020526000908152604090205481565b6040519081526020016100e3565b6000546100c29073ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff163314610234576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604b60248201527f446562616e6b4c3252656769737465723a2066756e6374696f6e2063616e206f60448201527f6e6c792062652063616c6c656420627920746865206f776e6572206f6620746860648201527f697320636f6e7472616374000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166102fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f446562616e6b4c3252656769737465723a2063616e206f6e6c7920626520646960448201527f7361626c65642076696120656e61626c65417262697472617279436f6e74726160648201527f63744465706c6f796d656e740000000000000000000000000000000000000000608482015260a40161022b565b6000546040805173ffffffffffffffffffffffffffffffffffffffff928316815291831660208301527fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c910160405180910390a1600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b33600090815260026020908152604080832080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8616179055600191829052822080549192909161040190849061073c565b90915550503360008181526001602090815260409182902054825193845273ffffffffffffffffffffffffffffffffffffffff85169184019190915282820152517fcc0bec1447060c88cdc5a739cf29cfa26c453574dd3f5b9e4dcc317d6401cb1c9181900360600190a150565b606061049a7f0000000000000000000000000000000000000000000000000000000000000000610512565b6104c37f0000000000000000000000000000000000000000000000000000000000000000610512565b6104ec7f0000000000000000000000000000000000000000000000000000000000000000610512565b6040516020016104fe93929190610754565b604051602081830303815290604052905090565b60608160000361055557505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561057f5780610569816107ca565b91506105789050600a83610831565b9150610559565b60008167ffffffffffffffff81111561059a5761059a610845565b6040519080825280601f01601f1916602001820160405280156105c4576020820181803683370190505b5090505b8415610647576105d9600183610874565b91506105e6600a8661088b565b6105f190603061073c565b60f81b8183815181106106065761060661089f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610640600a86610831565b94506105c8565b949350505050565b60006020828403121561066157600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461068557600080fd5b9392505050565b60005b838110156106a757818101518382015260200161068f565b838111156106b6576000848401525b50505050565b60208152600082518060208401526106db81604085016020870161068c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561074f5761074f61070d565b500190565b6000845161076681846020890161068c565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516107a2816001850160208a0161068c565b600192019182015283516107bd81600284016020880161068c565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107fb576107fb61070d565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261084057610840610802565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156108865761088661070d565b500390565b60008261089a5761089a610802565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a",
}

// DebankL2RegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use DebankL2RegisterMetaData.ABI instead.
var DebankL2RegisterABI = DebankL2RegisterMetaData.ABI

// DebankL2RegisterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DebankL2RegisterMetaData.Bin instead.
var DebankL2RegisterBin = DebankL2RegisterMetaData.Bin

// DeployDebankL2Register deploys a new Ethereum contract, binding an instance of DebankL2Register to it.
func DeployDebankL2Register(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebankL2Register, error) {
	parsed, err := DebankL2RegisterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DebankL2RegisterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebankL2Register{DebankL2RegisterCaller: DebankL2RegisterCaller{contract: contract}, DebankL2RegisterTransactor: DebankL2RegisterTransactor{contract: contract}, DebankL2RegisterFilterer: DebankL2RegisterFilterer{contract: contract}}, nil
}

// DebankL2Register is an auto generated Go binding around an Ethereum contract.
type DebankL2Register struct {
	DebankL2RegisterCaller     // Read-only binding to the contract
	DebankL2RegisterTransactor // Write-only binding to the contract
	DebankL2RegisterFilterer   // Log filterer for contract events
}

// DebankL2RegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebankL2RegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebankL2RegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebankL2RegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebankL2RegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebankL2RegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebankL2RegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebankL2RegisterSession struct {
	Contract     *DebankL2Register // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebankL2RegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebankL2RegisterCallerSession struct {
	Contract *DebankL2RegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DebankL2RegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebankL2RegisterTransactorSession struct {
	Contract     *DebankL2RegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DebankL2RegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebankL2RegisterRaw struct {
	Contract *DebankL2Register // Generic contract binding to access the raw methods on
}

// DebankL2RegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebankL2RegisterCallerRaw struct {
	Contract *DebankL2RegisterCaller // Generic read-only contract binding to access the raw methods on
}

// DebankL2RegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebankL2RegisterTransactorRaw struct {
	Contract *DebankL2RegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebankL2Register creates a new instance of DebankL2Register, bound to a specific deployed contract.
func NewDebankL2Register(address common.Address, backend bind.ContractBackend) (*DebankL2Register, error) {
	contract, err := bindDebankL2Register(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebankL2Register{DebankL2RegisterCaller: DebankL2RegisterCaller{contract: contract}, DebankL2RegisterTransactor: DebankL2RegisterTransactor{contract: contract}, DebankL2RegisterFilterer: DebankL2RegisterFilterer{contract: contract}}, nil
}

// NewDebankL2RegisterCaller creates a new read-only instance of DebankL2Register, bound to a specific deployed contract.
func NewDebankL2RegisterCaller(address common.Address, caller bind.ContractCaller) (*DebankL2RegisterCaller, error) {
	contract, err := bindDebankL2Register(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebankL2RegisterCaller{contract: contract}, nil
}

// NewDebankL2RegisterTransactor creates a new write-only instance of DebankL2Register, bound to a specific deployed contract.
func NewDebankL2RegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*DebankL2RegisterTransactor, error) {
	contract, err := bindDebankL2Register(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebankL2RegisterTransactor{contract: contract}, nil
}

// NewDebankL2RegisterFilterer creates a new log filterer instance of DebankL2Register, bound to a specific deployed contract.
func NewDebankL2RegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*DebankL2RegisterFilterer, error) {
	contract, err := bindDebankL2Register(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebankL2RegisterFilterer{contract: contract}, nil
}

// bindDebankL2Register binds a generic wrapper to an already deployed contract.
func bindDebankL2Register(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebankL2RegisterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebankL2Register *DebankL2RegisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebankL2Register.Contract.DebankL2RegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebankL2Register *DebankL2RegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebankL2Register.Contract.DebankL2RegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebankL2Register *DebankL2RegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebankL2Register.Contract.DebankL2RegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebankL2Register *DebankL2RegisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebankL2Register.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebankL2Register *DebankL2RegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebankL2Register.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebankL2Register *DebankL2RegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebankL2Register.Contract.contract.Transact(opts, method, params...)
}

// L2Accounts is a free data retrieval call binding the contract method 0x2a71d08d.
//
// Solidity: function l2Accounts(address ) view returns(address)
func (_DebankL2Register *DebankL2RegisterCaller) L2Accounts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DebankL2Register.contract.Call(opts, &out, "l2Accounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Accounts is a free data retrieval call binding the contract method 0x2a71d08d.
//
// Solidity: function l2Accounts(address ) view returns(address)
func (_DebankL2Register *DebankL2RegisterSession) L2Accounts(arg0 common.Address) (common.Address, error) {
	return _DebankL2Register.Contract.L2Accounts(&_DebankL2Register.CallOpts, arg0)
}

// L2Accounts is a free data retrieval call binding the contract method 0x2a71d08d.
//
// Solidity: function l2Accounts(address ) view returns(address)
func (_DebankL2Register *DebankL2RegisterCallerSession) L2Accounts(arg0 common.Address) (common.Address, error) {
	return _DebankL2Register.Contract.L2Accounts(&_DebankL2Register.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DebankL2Register *DebankL2RegisterCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DebankL2Register.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DebankL2Register *DebankL2RegisterSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DebankL2Register.Contract.Nonces(&_DebankL2Register.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_DebankL2Register *DebankL2RegisterCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _DebankL2Register.Contract.Nonces(&_DebankL2Register.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DebankL2Register *DebankL2RegisterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DebankL2Register.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DebankL2Register *DebankL2RegisterSession) Owner() (common.Address, error) {
	return _DebankL2Register.Contract.Owner(&_DebankL2Register.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DebankL2Register *DebankL2RegisterCallerSession) Owner() (common.Address, error) {
	return _DebankL2Register.Contract.Owner(&_DebankL2Register.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DebankL2Register *DebankL2RegisterCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DebankL2Register.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DebankL2Register *DebankL2RegisterSession) Version() (string, error) {
	return _DebankL2Register.Contract.Version(&_DebankL2Register.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DebankL2Register *DebankL2RegisterCallerSession) Version() (string, error) {
	return _DebankL2Register.Contract.Version(&_DebankL2Register.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address l2Account) returns()
func (_DebankL2Register *DebankL2RegisterTransactor) Register(opts *bind.TransactOpts, l2Account common.Address) (*types.Transaction, error) {
	return _DebankL2Register.contract.Transact(opts, "register", l2Account)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address l2Account) returns()
func (_DebankL2Register *DebankL2RegisterSession) Register(l2Account common.Address) (*types.Transaction, error) {
	return _DebankL2Register.Contract.Register(&_DebankL2Register.TransactOpts, l2Account)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address l2Account) returns()
func (_DebankL2Register *DebankL2RegisterTransactorSession) Register(l2Account common.Address) (*types.Transaction, error) {
	return _DebankL2Register.Contract.Register(&_DebankL2Register.TransactOpts, l2Account)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_DebankL2Register *DebankL2RegisterTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _DebankL2Register.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_DebankL2Register *DebankL2RegisterSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _DebankL2Register.Contract.SetOwner(&_DebankL2Register.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_DebankL2Register *DebankL2RegisterTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _DebankL2Register.Contract.SetOwner(&_DebankL2Register.TransactOpts, _owner)
}

// DebankL2RegisterOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the DebankL2Register contract.
type DebankL2RegisterOwnerChangedIterator struct {
	Event *DebankL2RegisterOwnerChanged // Event containing the contract specifics and raw log

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
func (it *DebankL2RegisterOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebankL2RegisterOwnerChanged)
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
		it.Event = new(DebankL2RegisterOwnerChanged)
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
func (it *DebankL2RegisterOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebankL2RegisterOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebankL2RegisterOwnerChanged represents a OwnerChanged event raised by the DebankL2Register contract.
type DebankL2RegisterOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_DebankL2Register *DebankL2RegisterFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*DebankL2RegisterOwnerChangedIterator, error) {

	logs, sub, err := _DebankL2Register.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &DebankL2RegisterOwnerChangedIterator{contract: _DebankL2Register.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_DebankL2Register *DebankL2RegisterFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *DebankL2RegisterOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _DebankL2Register.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebankL2RegisterOwnerChanged)
				if err := _DebankL2Register.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_DebankL2Register *DebankL2RegisterFilterer) ParseOwnerChanged(log types.Log) (*DebankL2RegisterOwnerChanged, error) {
	event := new(DebankL2RegisterOwnerChanged)
	if err := _DebankL2Register.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebankL2RegisterRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the DebankL2Register contract.
type DebankL2RegisterRegisterIterator struct {
	Event *DebankL2RegisterRegister // Event containing the contract specifics and raw log

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
func (it *DebankL2RegisterRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebankL2RegisterRegister)
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
		it.Event = new(DebankL2RegisterRegister)
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
func (it *DebankL2RegisterRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebankL2RegisterRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebankL2RegisterRegister represents a Register event raised by the DebankL2Register contract.
type DebankL2RegisterRegister struct {
	User        common.Address
	L2Account   common.Address
	RegisterCnt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0xcc0bec1447060c88cdc5a739cf29cfa26c453574dd3f5b9e4dcc317d6401cb1c.
//
// Solidity: event Register(address user, address l2Account, uint256 registerCnt)
func (_DebankL2Register *DebankL2RegisterFilterer) FilterRegister(opts *bind.FilterOpts) (*DebankL2RegisterRegisterIterator, error) {

	logs, sub, err := _DebankL2Register.contract.FilterLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return &DebankL2RegisterRegisterIterator{contract: _DebankL2Register.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0xcc0bec1447060c88cdc5a739cf29cfa26c453574dd3f5b9e4dcc317d6401cb1c.
//
// Solidity: event Register(address user, address l2Account, uint256 registerCnt)
func (_DebankL2Register *DebankL2RegisterFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *DebankL2RegisterRegister) (event.Subscription, error) {

	logs, sub, err := _DebankL2Register.contract.WatchLogs(opts, "Register")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebankL2RegisterRegister)
				if err := _DebankL2Register.contract.UnpackLog(event, "Register", log); err != nil {
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

// ParseRegister is a log parse operation binding the contract event 0xcc0bec1447060c88cdc5a739cf29cfa26c453574dd3f5b9e4dcc317d6401cb1c.
//
// Solidity: event Register(address user, address l2Account, uint256 registerCnt)
func (_DebankL2Register *DebankL2RegisterFilterer) ParseRegister(log types.Log) (*DebankL2RegisterRegister, error) {
	event := new(DebankL2RegisterRegister)
	if err := _DebankL2Register.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
