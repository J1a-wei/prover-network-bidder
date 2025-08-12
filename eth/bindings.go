// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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
	_ = abi.ConvertType
)

// FeeParams is an auto generated low-level Go binding around an user-defined struct.
type FeeParams struct {
	MaxFee   *big.Int
	MinStake *big.Int
	Deadline uint64
}

// ProofRequest is an auto generated low-level Go binding around an user-defined struct.
type ProofRequest struct {
	Nonce              uint64
	Vk                 [32]byte
	PublicValuesDigest [32]byte
	ImgURL             string
	InputData          [][]byte
	InputURL           string
	Fee                FeeParams
}

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"numRoleAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roleAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"accounts\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAccounts is a free data retrieval call binding the contract method 0xe1ed0a82.
//
// Solidity: function getRoleAccounts(bytes32 role) view returns(address[] accounts)
func (_AccessControl *AccessControlCaller) GetRoleAccounts(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleAccounts", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleAccounts is a free data retrieval call binding the contract method 0xe1ed0a82.
//
// Solidity: function getRoleAccounts(bytes32 role) view returns(address[] accounts)
func (_AccessControl *AccessControlSession) GetRoleAccounts(role [32]byte) ([]common.Address, error) {
	return _AccessControl.Contract.GetRoleAccounts(&_AccessControl.CallOpts, role)
}

// GetRoleAccounts is a free data retrieval call binding the contract method 0xe1ed0a82.
//
// Solidity: function getRoleAccounts(bytes32 role) view returns(address[] accounts)
func (_AccessControl *AccessControlCallerSession) GetRoleAccounts(role [32]byte) ([]common.Address, error) {
	return _AccessControl.Contract.GetRoleAccounts(&_AccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// NumRoleAccounts is a free data retrieval call binding the contract method 0xa1b5ff08.
//
// Solidity: function numRoleAccounts(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCaller) NumRoleAccounts(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "numRoleAccounts", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRoleAccounts is a free data retrieval call binding the contract method 0xa1b5ff08.
//
// Solidity: function numRoleAccounts(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlSession) NumRoleAccounts(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.NumRoleAccounts(&_AccessControl.CallOpts, role)
}

// NumRoleAccounts is a free data retrieval call binding the contract method 0xa1b5ff08.
//
// Solidity: function numRoleAccounts(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCallerSession) NumRoleAccounts(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.NumRoleAccounts(&_AccessControl.CallOpts, role)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControl *AccessControlCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControl *AccessControlSession) Owner() (common.Address, error) {
	return _AccessControl.Contract.Owner(&_AccessControl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControl *AccessControlCallerSession) Owner() (common.Address, error) {
	return _AccessControl.Contract.Owner(&_AccessControl.CallOpts)
}

// RoleAccounts is a free data retrieval call binding the contract method 0x0ec693ea.
//
// Solidity: function roleAccounts(bytes32 role, uint256 ) view returns(address accounts)
func (_AccessControl *AccessControlCaller) RoleAccounts(opts *bind.CallOpts, role [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "roleAccounts", role, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleAccounts is a free data retrieval call binding the contract method 0x0ec693ea.
//
// Solidity: function roleAccounts(bytes32 role, uint256 ) view returns(address accounts)
func (_AccessControl *AccessControlSession) RoleAccounts(role [32]byte, arg1 *big.Int) (common.Address, error) {
	return _AccessControl.Contract.RoleAccounts(&_AccessControl.CallOpts, role, arg1)
}

// RoleAccounts is a free data retrieval call binding the contract method 0x0ec693ea.
//
// Solidity: function roleAccounts(bytes32 role, uint256 ) view returns(address accounts)
func (_AccessControl *AccessControlCallerSession) RoleAccounts(role [32]byte, arg1 *big.Int) (common.Address, error) {
	return _AccessControl.Contract.RoleAccounts(&_AccessControl.CallOpts, role, arg1)
}

// Roles is a free data retrieval call binding the contract method 0xf8fc08b9.
//
// Solidity: function roles(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) Roles(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "roles", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0xf8fc08b9.
//
// Solidity: function roles(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) Roles(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.Roles(&_AccessControl.CallOpts, role, account)
}

// Roles is a free data retrieval call binding the contract method 0xf8fc08b9.
//
// Solidity: function roles(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) Roles(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.Roles(&_AccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlTransactor) GrantRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRoles", role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRoles(&_AccessControl.TransactOpts, role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRoles(&_AccessControl.TransactOpts, role, accounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlTransactor) RevokeRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRoles", role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRoles(&_AccessControl.TransactOpts, role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRoles(&_AccessControl.TransactOpts, role, accounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccessControl *AccessControlTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccessControl *AccessControlSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.TransferOwnership(&_AccessControl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccessControl *AccessControlTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.TransferOwnership(&_AccessControl.TransactOpts, newOwner)
}

// AccessControlOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccessControl contract.
type AccessControlOwnershipTransferredIterator struct {
	Event *AccessControlOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccessControlOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlOwnershipTransferred)
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
		it.Event = new(AccessControlOwnershipTransferred)
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
func (it *AccessControlOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlOwnershipTransferred represents a OwnershipTransferred event raised by the AccessControl contract.
type AccessControlOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccessControl *AccessControlFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccessControlOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlOwnershipTransferredIterator{contract: _AccessControl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccessControl *AccessControlFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlOwnershipTransferred)
				if err := _AccessControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccessControl *AccessControlFilterer) ParseOwnershipTransferred(log types.Log) (*AccessControlOwnershipTransferred, error) {
	event := new(AccessControlOwnershipTransferred)
	if err := _AccessControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 role, address account)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts) (*AccessControlRoleGrantedIterator, error) {

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted")
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 role, address account)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted) (event.Subscription, error) {

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 role, address account)
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 role, address account)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts) (*AccessControlRoleRevokedIterator, error) {

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 role, address account)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 role, address account)
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketMetaData contains all meta data concerning the BrevisMarket contract.
var BrevisMarketMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BidRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldDuration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"BiddingPhaseDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"imgURL\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"inputData\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"inputURL\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structFeeParams\",\"name\":\"fee\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structProofRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"NewRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"PicoVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"ProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldDuration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"RevealPhaseDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"biddingPhaseDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIPicoVerifier\",\"name\":\"_picoVerifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_biddingPhaseDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_revealPhaseDuration\",\"type\":\"uint64\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"numRoleAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"picoVerifier\",\"outputs\":[{\"internalType\":\"contractIPicoVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"imgURL\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"inputData\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"inputURL\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structFeeParams\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structProofRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealPhaseDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roleAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"accounts\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"setBiddingPhaseDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPicoVerifier\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"setPicoVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"setRevealPhaseDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"submitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461005b5760008054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a361234990816100618239f35b600080fdfe6080604052600436101561001257600080fd5b6000803560e01c80630ec693ea146118e3578063196f0f62146118835780632cf5d279146114af5780632f2ff15d14611468578063434f967c146113765780635950f9f1146111d65780637249fbb6146110685780637d9b715814610fb85780637e88b1a014610c275780638894a09714610bf35780638bb9c5bf14610bd55780638da5cb5b14610ba257806391d1485414610b715780639559196614610a5357806396f6fe9114610a2b578063a1b5ff0814610a01578063c22d3155146104db578063d547741f14610494578063deb9a3a21461042b578063dfc7537214610400578063e1ed0a8214610330578063eaf57ad714610255578063f2fde38b146101815763f8fc08b91461012557600080fd5b3461017e57604060031936011261017e5773ffffffffffffffffffffffffffffffffffffffff60406101556119e6565b9260043581526001602052209116600052602052602060ff604060002054166040519015158152f35b80fd5b503461017e57602060031936011261017e5761019b611a0e565b73ffffffffffffffffffffffffffffffffffffffff6101be818454163314611d7e565b8116156101d1576101ce906122a6565b80f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b503461017e57602060031936011261017e577ffbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda67ffffffffffffffff610299611ad6565b6102bb73ffffffffffffffffffffffffffffffffffffffff8554163314611d7e565b61032a600354916fffffffffffffffff00000000000000008160401b167fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff84161760035560405193849360401c168390602090939293604083019467ffffffffffffffff809216845216910152565b0390a180f35b503461017e576020806003193601126103fc576004358252600281526040822060405192838383549182815201908193835284832090835b8181106103d2575050508461037e9103856119a5565b60405193838594850191818652518092526040850193925b8281106103a557505050500390f35b835173ffffffffffffffffffffffffffffffffffffffff1685528695509381019392810192600101610396565b825473ffffffffffffffffffffffffffffffffffffffff1684529286019260019283019201610368565b5080fd5b503461017e578060031936011261017e57602067ffffffffffffffff60035460401c16604051908152f35b503461017e5761043a36611a31565b73ffffffffffffffffffffffffffffffffffffffff92919290610461828454163314611d7e565b825b815181101561049057806104868461047e61048b9486611e10565b511687611e24565b611de3565b610463565b8380f35b503461017e57604060031936011261017e576101ce6104b16119e6565b6104d373ffffffffffffffffffffffffffffffffffffffff8454163314611d7e565b600435611fdf565b506003196020813601126103fc5767ffffffffffffffff600435116103fc576101206004353603918201126103fc5760c4600435013534106109a35761052661010460043501611b02565b67ffffffffffffffff429116111561094557610546600435600401611b02565b6040517fffffffffffffffff000000000000000000000000000000000000000000000000602082019260c01b168252602460043501356028820152604460043501356048820152604881526080810181811067ffffffffffffffff8211176109165760405251902090818352600560205260408320805467ffffffffffffffff8160081c166108b8577fffffff0000000000000000000000000000000000000000000000000000000000164260081b68ffffffffffffffff0016173360481b7cffffffffffffffffffffffffffffffffffffffff0000000000000000001617815560043560c4810135600183015560e48101356002830155600382019067ffffffffffffffff9061065a9061010401611b02565b167fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000825416179055602460043501356004820155600560446004350135910155604051906020825267ffffffffffffffff6106b9600435600401611aed565b1660208301526024600435013560408301526044600435013560608301526107016106ee606460043501600435600401611b17565b6101206080860152610140850191611b67565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd6084600435013591018112156108b457600435019067ffffffffffffffff6004830135116108b457600482013560051b80360360248401136108b057907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08482030160a08501526004830135815260208082019282010192602481019287915b600481013583106108595788887f7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123898061081b8b6107eb60a460043501600435600401611b17565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c0860152611b67565b60c4600435013560e083015260e4600435013561010083015267ffffffffffffffff61084c61010460043501611aed565b166101208301520390a280f35b90919293946020806108a1837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0896001960301875261089b8a60248801611b17565b90611b67565b970193019301919392906107a3565b8580fd5b8480fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7265717565737420616c726561647920657869737473000000000000000000006044820152fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f646561646c696e65206d75737420626520696e206675747572650000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f696e73756666696369656e7420666565000000000000000000000000000000006044820152fd5b503461017e57602060031936011261017e5760406020916004358152600283522054604051908152f35b503461017e578060031936011261017e57602067ffffffffffffffff60035416604051908152f35b503461017e57602060031936011261017e57610a6d611a0e565b73ffffffffffffffffffffffffffffffffffffffff8091610a92828554163314611d7e565b16908115610aed57600454827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600455167f289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a808380a380f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f6e65772076657269666965722063616e6e6f74206265207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152fd5b503461017e57604060031936011261017e5773ffffffffffffffffffffffffffffffffffffffff60406101556119e6565b503461017e578060031936011261017e5773ffffffffffffffffffffffffffffffffffffffff6020915416604051908152f35b503461017e57602060031936011261017e576101ce33600435611fdf565b503461017e578060031936011261017e57602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b503461017e5761012060031936011261017e5760043536610124116103fc57808252602090600582526040832067ffffffffffffffff90816003820154164211610f5a57600781019073ffffffffffffffffffffffffffffffffffffffff808354163303610efc5760ff8254166003811015610ecf57610ca79015611c56565b8060045416956004830154600584015495883b15610ecb57604051917ff39751a8000000000000000000000000000000000000000000000000000000008352600483015289826101448160249c8d9b8c830152610100809c60448401375afa8015610ec057610e80575b505086885b60088110610e6b57505060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0084541617835587600a840154948360098601541615610e5f575b81808786829454165af1610d6f611cbb565b5015610e025750600190825460481c16910154918203918211610dd457858080807ff50908f290de42a8782b2026370f8d61e1d1de8f3018a20ae36ee019abb3707b9695610dc7955af1610dc1611cbb565b50611d19565b806040519485373393a380f35b847f4e487b710000000000000000000000000000000000000000000000000000000060005260116004526000fd5b606490601988604051927f08c379a000000000000000000000000000000000000000000000000000000000845260048401528201527f73656e642066656520746f2070726f766572206661696c6564000000000000006044820152fd5b60088501549550610d5d565b8135858201600b015590820190600101610d16565b8199929911610e9457604052963880610d11565b87827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6040513d8c823e3d90fd5b8980fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b606486604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152601360248201527f6e6f742065787065637465642070726f766572000000000000000000000000006044820152fd5b606484604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152600f60248201527f646561646c696e652070617373656400000000000000000000000000000000006044820152fd5b503461017e57602060031936011261017e577fae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e610ff3611ad6565b61101573ffffffffffffffffffffffffffffffffffffffff8454163314611d7e565b6003805467ffffffffffffffff9283167fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000821681179092556040805193909116835260208301919091528190810161032a565b503461017e57602060031936011261017e5760043580825260056020526040822090600182019167ffffffffffffffff60038201541642111561117857805460ff8116600381101561114b579160027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006020936111067ff552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe1839615611c56565b16179081815561113a878080808a5473ffffffffffffffffffffffffffffffffffffffff809860481c165af1610dc1611cbb565b5460481c169354604051908152a380f35b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6265666f726520646561646c696e6500000000000000000000000000000000006044820152fd5b503461017e57608060031936011261017e576111f0611a0e565b6111f86119e6565b906044359067ffffffffffffffff928383168093036108b45760643593841684036108b45773ffffffffffffffffffffffffffffffffffffffff918286541661131857828116156112ba5761124c906122a6565b167fffffffffffffffffffffffff000000000000000000000000000000000000000060045416176004557fffffffffffffffffffffffffffffffff000000000000000000000000000000006fffffffffffffffff00000000000000006003549360401b169216171760035580f35b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6e6577206f776e657220697320746865207a65726f20616464726573730000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6f776e657220616c7265616479207365740000000000000000000000000000006044820152fd5b503461017e57604060031936011261017e576004356024359080835260056020526040832067ffffffffffffffff6113c481835460081c166113b9811515611ba6565b826003541690611c0b565b16421161140a57600690338552016020528160408420556040519182527fe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e60203393a380f35b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f62696464696e6720706861736520656e646564000000000000000000000000006044820152fd5b503461017e57604060031936011261017e576101ce6114856119e6565b6114a773ffffffffffffffffffffffffffffffffffffffff8454163314611d7e565b600435611e24565b503461017e57606060031936011261017e57600435602435908083526020600581526040842067ffffffffffffffff80825460081c166114f0811515611ba6565b6003549082821690836115038383611c0b565b1642111561182557918361151b611526938295611c0b565b9160401c1690611c0b565b1642116117c75760405183810186815260443560408301526040825260608201928284109084111761179a578260405281519020338852600684018552604088205403611740575050600181015484116116e25790838260077f9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f94019073ffffffffffffffffffffffffffffffffffffffff91828154161580156116d5575b15611660576008926009830190828203611626575b505083856040516115ea81611989565b3381520152337fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905501555b6040519384523393a380f35b8254167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905582820154600a83015538806115da565b506009810180549283161580156116c8575b611680575b5050505061161a565b600a92848660405161169181611989565b33815201527fffffffffffffffffffffffff0000000000000000000000000000000000000000339116179055015583388080611677565b50600a8201548410611672565b50600882015484106115c5565b606482604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152601360248201527f6665652065786365656473206d6178696d756d000000000000000000000000006044820152fd5b907f6d69736d61746368206269642072657665616c0000000000000000000000000060a4606493867f08c379a000000000000000000000000000000000000000000000000000000000855285820152601360848201520152fd5b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b606483604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152601260248201527f72657665616c20706861736520656e64656400000000000000000000000000006044820152fd5b606486604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152601760248201527f62696464696e67207068617365206e6f7420656e6465640000000000000000006044820152fd5b503461017e5761189236611a31565b73ffffffffffffffffffffffffffffffffffffffff929192906118b9828454163314611d7e565b825b81518110156104905780610486846118d66118de9486611e10565b511687611fdf565b6118bb565b503461017e57604060031936011261017e5760243590600435815260026020526040812090815483101561017e57602073ffffffffffffffffffffffffffffffffffffffff6119328585611942565b9190546040519260031b1c168152f35b805482101561195a5760005260206000200190600090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6040810190811067ffffffffffffffff82111761091657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761091657604052565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203611a0957565b600080fd5b6004359073ffffffffffffffffffffffffffffffffffffffff82168203611a0957565b906040600319830112611a0957600435916024359067ffffffffffffffff90818311611a095780602384011215611a09578260040135918211610916578160051b60405193602093611a85858401876119a5565b8552602484860192820101928311611a0957602401905b828210611aaa575050505090565b813573ffffffffffffffffffffffffffffffffffffffff81168103611a09578152908301908301611a9c565b6004359067ffffffffffffffff82168203611a0957565b359067ffffffffffffffff82168203611a0957565b3567ffffffffffffffff81168103611a095790565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215611a0957016020813591019167ffffffffffffffff8211611a09578136038313611a0957565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b15611bad57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7265717565737420646f6573206e6f74206578697374000000000000000000006044820152fd5b91909167ffffffffffffffff80809416911601918211611c2757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b15611c5d57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964207265712073746174757300000000000000000000000000006044820152fd5b3d15611d14573d9067ffffffffffffffff82116109165760405191611d0860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601846119a5565b82523d6000602084013e565b606090565b15611d2057565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f726566756e6420666565206661696c65640000000000000000000000000000006044820152fd5b15611d8557565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611c275760010190565b805182101561195a5760209160051b010190565b91906000928084526020600181526040908186209073ffffffffffffffffffffffffffffffffffffffff851691828852815260ff8388205416611f82578387526002815282872080549068010000000000000000821015611f555792849392611eeb88611ebe85611f50999760017f2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f39e9f98018155611942565b90919082549060031b9173ffffffffffffffffffffffffffffffffffffffff809116831b921b1916179055565b86835260018152838320918352522060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055519283928390929173ffffffffffffffffffffffffffffffffffffffff6020916040840195845216910152565b0390a1565b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6064908351907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152601060248201527f616c72656164792068617320726f6c65000000000000000000000000000000006044820152fd5b919060008381526001906020938285526040938483209673ffffffffffffffffffffffffffffffffffffffff9788831690818652885260ff8786205416156122495781855260028852868520988954997fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808c019b8c1161221c5787895b6120bf575b60648c8c51907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152601660248201527f726f6c65206163636f756e74206e6f7420666f756e64000000000000000000006044820152fd5b82548110156122175783856120d48386611942565b929054600393841b1c16146120fa5750906120f18a949392611de3565b9091929361205d565b94979b969a989c908185949592829c989c106121f0575b505050825480156121c357927f155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a529b9c959289989592611f509a989501926121588484611942565b8154921b9290921b19169055558784528152838320918352522080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555191825273ffffffffffffffffffffffffffffffffffffffff90921660208201529081906040820190565b60248d7f4e487b710000000000000000000000000000000000000000000000000000000081526031600452fd5b61220f92612201611ebe9288611942565b905490891b1c169186611942565b813880612111565b612062565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b6064888851907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152600c60248201527f6e6f742068617320726f6c6500000000000000000000000000000000000000006044820152fd5b6000549073ffffffffffffffffffffffffffffffffffffffff80911691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a356fea264697066735822122001e2aadd8af3d6a42ec06152b45446d108283948f0b18b7bf16806c53ed282d864736f6c63430008140033",
}

// BrevisMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use BrevisMarketMetaData.ABI instead.
var BrevisMarketABI = BrevisMarketMetaData.ABI

// BrevisMarketBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrevisMarketMetaData.Bin instead.
var BrevisMarketBin = BrevisMarketMetaData.Bin

// DeployBrevisMarket deploys a new Ethereum contract, binding an instance of BrevisMarket to it.
func DeployBrevisMarket(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BrevisMarket, error) {
	parsed, err := BrevisMarketMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrevisMarketBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BrevisMarket{BrevisMarketCaller: BrevisMarketCaller{contract: contract}, BrevisMarketTransactor: BrevisMarketTransactor{contract: contract}, BrevisMarketFilterer: BrevisMarketFilterer{contract: contract}}, nil
}

// BrevisMarket is an auto generated Go binding around an Ethereum contract.
type BrevisMarket struct {
	BrevisMarketCaller     // Read-only binding to the contract
	BrevisMarketTransactor // Write-only binding to the contract
	BrevisMarketFilterer   // Log filterer for contract events
}

// BrevisMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrevisMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrevisMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrevisMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrevisMarketSession struct {
	Contract     *BrevisMarket     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrevisMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrevisMarketCallerSession struct {
	Contract *BrevisMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BrevisMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrevisMarketTransactorSession struct {
	Contract     *BrevisMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BrevisMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrevisMarketRaw struct {
	Contract *BrevisMarket // Generic contract binding to access the raw methods on
}

// BrevisMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrevisMarketCallerRaw struct {
	Contract *BrevisMarketCaller // Generic read-only contract binding to access the raw methods on
}

// BrevisMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrevisMarketTransactorRaw struct {
	Contract *BrevisMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrevisMarket creates a new instance of BrevisMarket, bound to a specific deployed contract.
func NewBrevisMarket(address common.Address, backend bind.ContractBackend) (*BrevisMarket, error) {
	contract, err := bindBrevisMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BrevisMarket{BrevisMarketCaller: BrevisMarketCaller{contract: contract}, BrevisMarketTransactor: BrevisMarketTransactor{contract: contract}, BrevisMarketFilterer: BrevisMarketFilterer{contract: contract}}, nil
}

// NewBrevisMarketCaller creates a new read-only instance of BrevisMarket, bound to a specific deployed contract.
func NewBrevisMarketCaller(address common.Address, caller bind.ContractCaller) (*BrevisMarketCaller, error) {
	contract, err := bindBrevisMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketCaller{contract: contract}, nil
}

// NewBrevisMarketTransactor creates a new write-only instance of BrevisMarket, bound to a specific deployed contract.
func NewBrevisMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*BrevisMarketTransactor, error) {
	contract, err := bindBrevisMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketTransactor{contract: contract}, nil
}

// NewBrevisMarketFilterer creates a new log filterer instance of BrevisMarket, bound to a specific deployed contract.
func NewBrevisMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*BrevisMarketFilterer, error) {
	contract, err := bindBrevisMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketFilterer{contract: contract}, nil
}

// bindBrevisMarket binds a generic wrapper to an already deployed contract.
func bindBrevisMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BrevisMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrevisMarket *BrevisMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrevisMarket.Contract.BrevisMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrevisMarket *BrevisMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevisMarket.Contract.BrevisMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrevisMarket *BrevisMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrevisMarket.Contract.BrevisMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrevisMarket *BrevisMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrevisMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrevisMarket *BrevisMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevisMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrevisMarket *BrevisMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrevisMarket.Contract.contract.Transact(opts, method, params...)
}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketCaller) BiddingPhaseDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "biddingPhaseDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketSession) BiddingPhaseDuration() (uint64, error) {
	return _BrevisMarket.Contract.BiddingPhaseDuration(&_BrevisMarket.CallOpts)
}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketCallerSession) BiddingPhaseDuration() (uint64, error) {
	return _BrevisMarket.Contract.BiddingPhaseDuration(&_BrevisMarket.CallOpts)
}

// GetRoleAccounts is a free data retrieval call binding the contract method 0xe1ed0a82.
//
// Solidity: function getRoleAccounts(bytes32 role) view returns(address[] accounts)
func (_BrevisMarket *BrevisMarketCaller) GetRoleAccounts(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "getRoleAccounts", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleAccounts is a free data retrieval call binding the contract method 0xe1ed0a82.
//
// Solidity: function getRoleAccounts(bytes32 role) view returns(address[] accounts)
func (_BrevisMarket *BrevisMarketSession) GetRoleAccounts(role [32]byte) ([]common.Address, error) {
	return _BrevisMarket.Contract.GetRoleAccounts(&_BrevisMarket.CallOpts, role)
}

// GetRoleAccounts is a free data retrieval call binding the contract method 0xe1ed0a82.
//
// Solidity: function getRoleAccounts(bytes32 role) view returns(address[] accounts)
func (_BrevisMarket *BrevisMarketCallerSession) GetRoleAccounts(role [32]byte) ([]common.Address, error) {
	return _BrevisMarket.Contract.GetRoleAccounts(&_BrevisMarket.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BrevisMarket.Contract.HasRole(&_BrevisMarket.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BrevisMarket.Contract.HasRole(&_BrevisMarket.CallOpts, role, account)
}

// NumRoleAccounts is a free data retrieval call binding the contract method 0xa1b5ff08.
//
// Solidity: function numRoleAccounts(bytes32 role) view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) NumRoleAccounts(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "numRoleAccounts", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRoleAccounts is a free data retrieval call binding the contract method 0xa1b5ff08.
//
// Solidity: function numRoleAccounts(bytes32 role) view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) NumRoleAccounts(role [32]byte) (*big.Int, error) {
	return _BrevisMarket.Contract.NumRoleAccounts(&_BrevisMarket.CallOpts, role)
}

// NumRoleAccounts is a free data retrieval call binding the contract method 0xa1b5ff08.
//
// Solidity: function numRoleAccounts(bytes32 role) view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) NumRoleAccounts(role [32]byte) (*big.Int, error) {
	return _BrevisMarket.Contract.NumRoleAccounts(&_BrevisMarket.CallOpts, role)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisMarket *BrevisMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisMarket *BrevisMarketSession) Owner() (common.Address, error) {
	return _BrevisMarket.Contract.Owner(&_BrevisMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisMarket *BrevisMarketCallerSession) Owner() (common.Address, error) {
	return _BrevisMarket.Contract.Owner(&_BrevisMarket.CallOpts)
}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address)
func (_BrevisMarket *BrevisMarketCaller) PicoVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "picoVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address)
func (_BrevisMarket *BrevisMarketSession) PicoVerifier() (common.Address, error) {
	return _BrevisMarket.Contract.PicoVerifier(&_BrevisMarket.CallOpts)
}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address)
func (_BrevisMarket *BrevisMarketCallerSession) PicoVerifier() (common.Address, error) {
	return _BrevisMarket.Contract.PicoVerifier(&_BrevisMarket.CallOpts)
}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketCaller) RevealPhaseDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "revealPhaseDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketSession) RevealPhaseDuration() (uint64, error) {
	return _BrevisMarket.Contract.RevealPhaseDuration(&_BrevisMarket.CallOpts)
}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketCallerSession) RevealPhaseDuration() (uint64, error) {
	return _BrevisMarket.Contract.RevealPhaseDuration(&_BrevisMarket.CallOpts)
}

// RoleAccounts is a free data retrieval call binding the contract method 0x0ec693ea.
//
// Solidity: function roleAccounts(bytes32 role, uint256 ) view returns(address accounts)
func (_BrevisMarket *BrevisMarketCaller) RoleAccounts(opts *bind.CallOpts, role [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "roleAccounts", role, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleAccounts is a free data retrieval call binding the contract method 0x0ec693ea.
//
// Solidity: function roleAccounts(bytes32 role, uint256 ) view returns(address accounts)
func (_BrevisMarket *BrevisMarketSession) RoleAccounts(role [32]byte, arg1 *big.Int) (common.Address, error) {
	return _BrevisMarket.Contract.RoleAccounts(&_BrevisMarket.CallOpts, role, arg1)
}

// RoleAccounts is a free data retrieval call binding the contract method 0x0ec693ea.
//
// Solidity: function roleAccounts(bytes32 role, uint256 ) view returns(address accounts)
func (_BrevisMarket *BrevisMarketCallerSession) RoleAccounts(role [32]byte, arg1 *big.Int) (common.Address, error) {
	return _BrevisMarket.Contract.RoleAccounts(&_BrevisMarket.CallOpts, role, arg1)
}

// Roles is a free data retrieval call binding the contract method 0xf8fc08b9.
//
// Solidity: function roles(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketCaller) Roles(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "roles", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0xf8fc08b9.
//
// Solidity: function roles(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketSession) Roles(role [32]byte, account common.Address) (bool, error) {
	return _BrevisMarket.Contract.Roles(&_BrevisMarket.CallOpts, role, account)
}

// Roles is a free data retrieval call binding the contract method 0xf8fc08b9.
//
// Solidity: function roles(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketCallerSession) Roles(role [32]byte, account common.Address) (bool, error) {
	return _BrevisMarket.Contract.Roles(&_BrevisMarket.CallOpts, role, account)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_BrevisMarket *BrevisMarketTransactor) Bid(opts *bind.TransactOpts, reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "bid", reqid, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_BrevisMarket *BrevisMarketSession) Bid(reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Bid(&_BrevisMarket.TransactOpts, reqid, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Bid(reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Bid(&_BrevisMarket.TransactOpts, reqid, bidHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.GrantRole(&_BrevisMarket.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.GrantRole(&_BrevisMarket.TransactOpts, role, account)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketTransactor) GrantRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "grantRoles", role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.GrantRoles(&_BrevisMarket.TransactOpts, role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.GrantRoles(&_BrevisMarket.TransactOpts, role, accounts)
}

// Init is a paid mutator transaction binding the contract method 0x5950f9f1.
//
// Solidity: function init(address _owner, address _picoVerifier, uint64 _biddingPhaseDuration, uint64 _revealPhaseDuration) returns()
func (_BrevisMarket *BrevisMarketTransactor) Init(opts *bind.TransactOpts, _owner common.Address, _picoVerifier common.Address, _biddingPhaseDuration uint64, _revealPhaseDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "init", _owner, _picoVerifier, _biddingPhaseDuration, _revealPhaseDuration)
}

// Init is a paid mutator transaction binding the contract method 0x5950f9f1.
//
// Solidity: function init(address _owner, address _picoVerifier, uint64 _biddingPhaseDuration, uint64 _revealPhaseDuration) returns()
func (_BrevisMarket *BrevisMarketSession) Init(_owner common.Address, _picoVerifier common.Address, _biddingPhaseDuration uint64, _revealPhaseDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Init(&_BrevisMarket.TransactOpts, _owner, _picoVerifier, _biddingPhaseDuration, _revealPhaseDuration)
}

// Init is a paid mutator transaction binding the contract method 0x5950f9f1.
//
// Solidity: function init(address _owner, address _picoVerifier, uint64 _biddingPhaseDuration, uint64 _revealPhaseDuration) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Init(_owner common.Address, _picoVerifier common.Address, _biddingPhaseDuration uint64, _revealPhaseDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Init(&_BrevisMarket.TransactOpts, _owner, _picoVerifier, _biddingPhaseDuration, _revealPhaseDuration)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketTransactor) Refund(opts *bind.TransactOpts, reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "refund", reqid)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketSession) Refund(reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Refund(&_BrevisMarket.TransactOpts, reqid)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Refund(reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Refund(&_BrevisMarket.TransactOpts, reqid)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_BrevisMarket *BrevisMarketTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_BrevisMarket *BrevisMarketSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RenounceRole(&_BrevisMarket.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RenounceRole(&_BrevisMarket.TransactOpts, role)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) payable returns()
func (_BrevisMarket *BrevisMarketTransactor) RequestProof(opts *bind.TransactOpts, req ProofRequest) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "requestProof", req)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) payable returns()
func (_BrevisMarket *BrevisMarketSession) RequestProof(req ProofRequest) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RequestProof(&_BrevisMarket.TransactOpts, req)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) payable returns()
func (_BrevisMarket *BrevisMarketTransactorSession) RequestProof(req ProofRequest) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RequestProof(&_BrevisMarket.TransactOpts, req)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_BrevisMarket *BrevisMarketTransactor) Reveal(opts *bind.TransactOpts, reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "reveal", reqid, fee, nonce)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_BrevisMarket *BrevisMarketSession) Reveal(reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Reveal(&_BrevisMarket.TransactOpts, reqid, fee, nonce)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Reveal(reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Reveal(&_BrevisMarket.TransactOpts, reqid, fee, nonce)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RevokeRole(&_BrevisMarket.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RevokeRole(&_BrevisMarket.TransactOpts, role, account)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketTransactor) RevokeRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "revokeRoles", role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RevokeRoles(&_BrevisMarket.TransactOpts, role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RevokeRoles(&_BrevisMarket.TransactOpts, role, accounts)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetBiddingPhaseDuration(opts *bind.TransactOpts, newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setBiddingPhaseDuration", newDuration)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketSession) SetBiddingPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetBiddingPhaseDuration(&_BrevisMarket.TransactOpts, newDuration)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetBiddingPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetBiddingPhaseDuration(&_BrevisMarket.TransactOpts, newDuration)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetPicoVerifier(opts *bind.TransactOpts, newVerifier common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setPicoVerifier", newVerifier)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_BrevisMarket *BrevisMarketSession) SetPicoVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetPicoVerifier(&_BrevisMarket.TransactOpts, newVerifier)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetPicoVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetPicoVerifier(&_BrevisMarket.TransactOpts, newVerifier)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetRevealPhaseDuration(opts *bind.TransactOpts, newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setRevealPhaseDuration", newDuration)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketSession) SetRevealPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetRevealPhaseDuration(&_BrevisMarket.TransactOpts, newDuration)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetRevealPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetRevealPhaseDuration(&_BrevisMarket.TransactOpts, newDuration)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_BrevisMarket *BrevisMarketTransactor) SubmitProof(opts *bind.TransactOpts, reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "submitProof", reqid, proof)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_BrevisMarket *BrevisMarketSession) SubmitProof(reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SubmitProof(&_BrevisMarket.TransactOpts, reqid, proof)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SubmitProof(reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SubmitProof(&_BrevisMarket.TransactOpts, reqid, proof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisMarket *BrevisMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisMarket *BrevisMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.TransferOwnership(&_BrevisMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.TransferOwnership(&_BrevisMarket.TransactOpts, newOwner)
}

// BrevisMarketBidRevealedIterator is returned from FilterBidRevealed and is used to iterate over the raw logs and unpacked data for BidRevealed events raised by the BrevisMarket contract.
type BrevisMarketBidRevealedIterator struct {
	Event *BrevisMarketBidRevealed // Event containing the contract specifics and raw log

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
func (it *BrevisMarketBidRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketBidRevealed)
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
		it.Event = new(BrevisMarketBidRevealed)
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
func (it *BrevisMarketBidRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketBidRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketBidRevealed represents a BidRevealed event raised by the BrevisMarket contract.
type BrevisMarketBidRevealed struct {
	Reqid  [32]byte
	Prover common.Address
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidRevealed is a free log retrieval operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_BrevisMarket *BrevisMarketFilterer) FilterBidRevealed(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*BrevisMarketBidRevealedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "BidRevealed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketBidRevealedIterator{contract: _BrevisMarket.contract, event: "BidRevealed", logs: logs, sub: sub}, nil
}

// WatchBidRevealed is a free log subscription operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_BrevisMarket *BrevisMarketFilterer) WatchBidRevealed(opts *bind.WatchOpts, sink chan<- *BrevisMarketBidRevealed, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "BidRevealed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketBidRevealed)
				if err := _BrevisMarket.contract.UnpackLog(event, "BidRevealed", log); err != nil {
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

// ParseBidRevealed is a log parse operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_BrevisMarket *BrevisMarketFilterer) ParseBidRevealed(log types.Log) (*BrevisMarketBidRevealed, error) {
	event := new(BrevisMarketBidRevealed)
	if err := _BrevisMarket.contract.UnpackLog(event, "BidRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketBiddingPhaseDurationUpdatedIterator is returned from FilterBiddingPhaseDurationUpdated and is used to iterate over the raw logs and unpacked data for BiddingPhaseDurationUpdated events raised by the BrevisMarket contract.
type BrevisMarketBiddingPhaseDurationUpdatedIterator struct {
	Event *BrevisMarketBiddingPhaseDurationUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketBiddingPhaseDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketBiddingPhaseDurationUpdated)
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
		it.Event = new(BrevisMarketBiddingPhaseDurationUpdated)
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
func (it *BrevisMarketBiddingPhaseDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketBiddingPhaseDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketBiddingPhaseDurationUpdated represents a BiddingPhaseDurationUpdated event raised by the BrevisMarket contract.
type BrevisMarketBiddingPhaseDurationUpdated struct {
	OldDuration uint64
	NewDuration uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBiddingPhaseDurationUpdated is a free log retrieval operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) FilterBiddingPhaseDurationUpdated(opts *bind.FilterOpts) (*BrevisMarketBiddingPhaseDurationUpdatedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "BiddingPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketBiddingPhaseDurationUpdatedIterator{contract: _BrevisMarket.contract, event: "BiddingPhaseDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchBiddingPhaseDurationUpdated is a free log subscription operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) WatchBiddingPhaseDurationUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketBiddingPhaseDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "BiddingPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketBiddingPhaseDurationUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "BiddingPhaseDurationUpdated", log); err != nil {
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

// ParseBiddingPhaseDurationUpdated is a log parse operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) ParseBiddingPhaseDurationUpdated(log types.Log) (*BrevisMarketBiddingPhaseDurationUpdated, error) {
	event := new(BrevisMarketBiddingPhaseDurationUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "BiddingPhaseDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the BrevisMarket contract.
type BrevisMarketNewBidIterator struct {
	Event *BrevisMarketNewBid // Event containing the contract specifics and raw log

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
func (it *BrevisMarketNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketNewBid)
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
		it.Event = new(BrevisMarketNewBid)
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
func (it *BrevisMarketNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketNewBid represents a NewBid event raised by the BrevisMarket contract.
type BrevisMarketNewBid struct {
	Reqid   [32]byte
	Prover  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_BrevisMarket *BrevisMarketFilterer) FilterNewBid(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*BrevisMarketNewBidIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "NewBid", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketNewBidIterator{contract: _BrevisMarket.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_BrevisMarket *BrevisMarketFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *BrevisMarketNewBid, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "NewBid", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketNewBid)
				if err := _BrevisMarket.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_BrevisMarket *BrevisMarketFilterer) ParseNewBid(log types.Log) (*BrevisMarketNewBid, error) {
	event := new(BrevisMarketNewBid)
	if err := _BrevisMarket.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketNewRequestIterator is returned from FilterNewRequest and is used to iterate over the raw logs and unpacked data for NewRequest events raised by the BrevisMarket contract.
type BrevisMarketNewRequestIterator struct {
	Event *BrevisMarketNewRequest // Event containing the contract specifics and raw log

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
func (it *BrevisMarketNewRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketNewRequest)
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
		it.Event = new(BrevisMarketNewRequest)
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
func (it *BrevisMarketNewRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketNewRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketNewRequest represents a NewRequest event raised by the BrevisMarket contract.
type BrevisMarketNewRequest struct {
	Reqid [32]byte
	Req   ProofRequest
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewRequest is a free log retrieval operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_BrevisMarket *BrevisMarketFilterer) FilterNewRequest(opts *bind.FilterOpts, reqid [][32]byte) (*BrevisMarketNewRequestIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "NewRequest", reqidRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketNewRequestIterator{contract: _BrevisMarket.contract, event: "NewRequest", logs: logs, sub: sub}, nil
}

// WatchNewRequest is a free log subscription operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_BrevisMarket *BrevisMarketFilterer) WatchNewRequest(opts *bind.WatchOpts, sink chan<- *BrevisMarketNewRequest, reqid [][32]byte) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "NewRequest", reqidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketNewRequest)
				if err := _BrevisMarket.contract.UnpackLog(event, "NewRequest", log); err != nil {
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

// ParseNewRequest is a log parse operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_BrevisMarket *BrevisMarketFilterer) ParseNewRequest(log types.Log) (*BrevisMarketNewRequest, error) {
	event := new(BrevisMarketNewRequest)
	if err := _BrevisMarket.contract.UnpackLog(event, "NewRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BrevisMarket contract.
type BrevisMarketOwnershipTransferredIterator struct {
	Event *BrevisMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BrevisMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketOwnershipTransferred)
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
		it.Event = new(BrevisMarketOwnershipTransferred)
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
func (it *BrevisMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketOwnershipTransferred represents a OwnershipTransferred event raised by the BrevisMarket contract.
type BrevisMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevisMarket *BrevisMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BrevisMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketOwnershipTransferredIterator{contract: _BrevisMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevisMarket *BrevisMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrevisMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketOwnershipTransferred)
				if err := _BrevisMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevisMarket *BrevisMarketFilterer) ParseOwnershipTransferred(log types.Log) (*BrevisMarketOwnershipTransferred, error) {
	event := new(BrevisMarketOwnershipTransferred)
	if err := _BrevisMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketPicoVerifierUpdatedIterator is returned from FilterPicoVerifierUpdated and is used to iterate over the raw logs and unpacked data for PicoVerifierUpdated events raised by the BrevisMarket contract.
type BrevisMarketPicoVerifierUpdatedIterator struct {
	Event *BrevisMarketPicoVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketPicoVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketPicoVerifierUpdated)
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
		it.Event = new(BrevisMarketPicoVerifierUpdated)
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
func (it *BrevisMarketPicoVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketPicoVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketPicoVerifierUpdated represents a PicoVerifierUpdated event raised by the BrevisMarket contract.
type BrevisMarketPicoVerifierUpdated struct {
	OldVerifier common.Address
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPicoVerifierUpdated is a free log retrieval operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_BrevisMarket *BrevisMarketFilterer) FilterPicoVerifierUpdated(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*BrevisMarketPicoVerifierUpdatedIterator, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "PicoVerifierUpdated", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketPicoVerifierUpdatedIterator{contract: _BrevisMarket.contract, event: "PicoVerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchPicoVerifierUpdated is a free log subscription operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_BrevisMarket *BrevisMarketFilterer) WatchPicoVerifierUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketPicoVerifierUpdated, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "PicoVerifierUpdated", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketPicoVerifierUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "PicoVerifierUpdated", log); err != nil {
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

// ParsePicoVerifierUpdated is a log parse operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_BrevisMarket *BrevisMarketFilterer) ParsePicoVerifierUpdated(log types.Log) (*BrevisMarketPicoVerifierUpdated, error) {
	event := new(BrevisMarketPicoVerifierUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "PicoVerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketProofSubmittedIterator is returned from FilterProofSubmitted and is used to iterate over the raw logs and unpacked data for ProofSubmitted events raised by the BrevisMarket contract.
type BrevisMarketProofSubmittedIterator struct {
	Event *BrevisMarketProofSubmitted // Event containing the contract specifics and raw log

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
func (it *BrevisMarketProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketProofSubmitted)
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
		it.Event = new(BrevisMarketProofSubmitted)
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
func (it *BrevisMarketProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketProofSubmitted represents a ProofSubmitted event raised by the BrevisMarket contract.
type BrevisMarketProofSubmitted struct {
	Reqid  [32]byte
	Prover common.Address
	Proof  [8]*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitted is a free log retrieval operation binding the contract event 0xf50908f290de42a8782b2026370f8d61e1d1de8f3018a20ae36ee019abb3707b.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof)
func (_BrevisMarket *BrevisMarketFilterer) FilterProofSubmitted(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*BrevisMarketProofSubmittedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "ProofSubmitted", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketProofSubmittedIterator{contract: _BrevisMarket.contract, event: "ProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchProofSubmitted is a free log subscription operation binding the contract event 0xf50908f290de42a8782b2026370f8d61e1d1de8f3018a20ae36ee019abb3707b.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof)
func (_BrevisMarket *BrevisMarketFilterer) WatchProofSubmitted(opts *bind.WatchOpts, sink chan<- *BrevisMarketProofSubmitted, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "ProofSubmitted", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketProofSubmitted)
				if err := _BrevisMarket.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
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

// ParseProofSubmitted is a log parse operation binding the contract event 0xf50908f290de42a8782b2026370f8d61e1d1de8f3018a20ae36ee019abb3707b.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof)
func (_BrevisMarket *BrevisMarketFilterer) ParseProofSubmitted(log types.Log) (*BrevisMarketProofSubmitted, error) {
	event := new(BrevisMarketProofSubmitted)
	if err := _BrevisMarket.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the BrevisMarket contract.
type BrevisMarketRefundedIterator struct {
	Event *BrevisMarketRefunded // Event containing the contract specifics and raw log

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
func (it *BrevisMarketRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketRefunded)
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
		it.Event = new(BrevisMarketRefunded)
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
func (it *BrevisMarketRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketRefunded represents a Refunded event raised by the BrevisMarket contract.
type BrevisMarketRefunded struct {
	Reqid     [32]byte
	Requester common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) FilterRefunded(opts *bind.FilterOpts, reqid [][32]byte, requester []common.Address) (*BrevisMarketRefundedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "Refunded", reqidRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketRefundedIterator{contract: _BrevisMarket.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *BrevisMarketRefunded, reqid [][32]byte, requester []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "Refunded", reqidRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketRefunded)
				if err := _BrevisMarket.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) ParseRefunded(log types.Log) (*BrevisMarketRefunded, error) {
	event := new(BrevisMarketRefunded)
	if err := _BrevisMarket.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketRevealPhaseDurationUpdatedIterator is returned from FilterRevealPhaseDurationUpdated and is used to iterate over the raw logs and unpacked data for RevealPhaseDurationUpdated events raised by the BrevisMarket contract.
type BrevisMarketRevealPhaseDurationUpdatedIterator struct {
	Event *BrevisMarketRevealPhaseDurationUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketRevealPhaseDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketRevealPhaseDurationUpdated)
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
		it.Event = new(BrevisMarketRevealPhaseDurationUpdated)
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
func (it *BrevisMarketRevealPhaseDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketRevealPhaseDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketRevealPhaseDurationUpdated represents a RevealPhaseDurationUpdated event raised by the BrevisMarket contract.
type BrevisMarketRevealPhaseDurationUpdated struct {
	OldDuration uint64
	NewDuration uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevealPhaseDurationUpdated is a free log retrieval operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) FilterRevealPhaseDurationUpdated(opts *bind.FilterOpts) (*BrevisMarketRevealPhaseDurationUpdatedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "RevealPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketRevealPhaseDurationUpdatedIterator{contract: _BrevisMarket.contract, event: "RevealPhaseDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRevealPhaseDurationUpdated is a free log subscription operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) WatchRevealPhaseDurationUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketRevealPhaseDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "RevealPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketRevealPhaseDurationUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "RevealPhaseDurationUpdated", log); err != nil {
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

// ParseRevealPhaseDurationUpdated is a log parse operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) ParseRevealPhaseDurationUpdated(log types.Log) (*BrevisMarketRevealPhaseDurationUpdated, error) {
	event := new(BrevisMarketRevealPhaseDurationUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "RevealPhaseDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BrevisMarket contract.
type BrevisMarketRoleGrantedIterator struct {
	Event *BrevisMarketRoleGranted // Event containing the contract specifics and raw log

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
func (it *BrevisMarketRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketRoleGranted)
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
		it.Event = new(BrevisMarketRoleGranted)
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
func (it *BrevisMarketRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketRoleGranted represents a RoleGranted event raised by the BrevisMarket contract.
type BrevisMarketRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 role, address account)
func (_BrevisMarket *BrevisMarketFilterer) FilterRoleGranted(opts *bind.FilterOpts) (*BrevisMarketRoleGrantedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "RoleGranted")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketRoleGrantedIterator{contract: _BrevisMarket.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 role, address account)
func (_BrevisMarket *BrevisMarketFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BrevisMarketRoleGranted) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "RoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketRoleGranted)
				if err := _BrevisMarket.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 role, address account)
func (_BrevisMarket *BrevisMarketFilterer) ParseRoleGranted(log types.Log) (*BrevisMarketRoleGranted, error) {
	event := new(BrevisMarketRoleGranted)
	if err := _BrevisMarket.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BrevisMarket contract.
type BrevisMarketRoleRevokedIterator struct {
	Event *BrevisMarketRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BrevisMarketRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketRoleRevoked)
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
		it.Event = new(BrevisMarketRoleRevoked)
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
func (it *BrevisMarketRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketRoleRevoked represents a RoleRevoked event raised by the BrevisMarket contract.
type BrevisMarketRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 role, address account)
func (_BrevisMarket *BrevisMarketFilterer) FilterRoleRevoked(opts *bind.FilterOpts) (*BrevisMarketRoleRevokedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketRoleRevokedIterator{contract: _BrevisMarket.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 role, address account)
func (_BrevisMarket *BrevisMarketFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BrevisMarketRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketRoleRevoked)
				if err := _BrevisMarket.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 role, address account)
func (_BrevisMarket *BrevisMarketFilterer) ParseRoleRevoked(log types.Log) (*BrevisMarketRoleRevoked, error) {
	event := new(BrevisMarketRoleRevoked)
	if err := _BrevisMarket.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPicoVerifierMetaData contains all meta data concerning the IPicoVerifier contract.
var IPicoVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"riscvVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"verifyPicoProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"riscvVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"verifyPicoProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPicoVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IPicoVerifierMetaData.ABI instead.
var IPicoVerifierABI = IPicoVerifierMetaData.ABI

// IPicoVerifier is an auto generated Go binding around an Ethereum contract.
type IPicoVerifier struct {
	IPicoVerifierCaller     // Read-only binding to the contract
	IPicoVerifierTransactor // Write-only binding to the contract
	IPicoVerifierFilterer   // Log filterer for contract events
}

// IPicoVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPicoVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPicoVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPicoVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPicoVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPicoVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPicoVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPicoVerifierSession struct {
	Contract     *IPicoVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPicoVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPicoVerifierCallerSession struct {
	Contract *IPicoVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPicoVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPicoVerifierTransactorSession struct {
	Contract     *IPicoVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPicoVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPicoVerifierRaw struct {
	Contract *IPicoVerifier // Generic contract binding to access the raw methods on
}

// IPicoVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPicoVerifierCallerRaw struct {
	Contract *IPicoVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IPicoVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPicoVerifierTransactorRaw struct {
	Contract *IPicoVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPicoVerifier creates a new instance of IPicoVerifier, bound to a specific deployed contract.
func NewIPicoVerifier(address common.Address, backend bind.ContractBackend) (*IPicoVerifier, error) {
	contract, err := bindIPicoVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPicoVerifier{IPicoVerifierCaller: IPicoVerifierCaller{contract: contract}, IPicoVerifierTransactor: IPicoVerifierTransactor{contract: contract}, IPicoVerifierFilterer: IPicoVerifierFilterer{contract: contract}}, nil
}

// NewIPicoVerifierCaller creates a new read-only instance of IPicoVerifier, bound to a specific deployed contract.
func NewIPicoVerifierCaller(address common.Address, caller bind.ContractCaller) (*IPicoVerifierCaller, error) {
	contract, err := bindIPicoVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPicoVerifierCaller{contract: contract}, nil
}

// NewIPicoVerifierTransactor creates a new write-only instance of IPicoVerifier, bound to a specific deployed contract.
func NewIPicoVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IPicoVerifierTransactor, error) {
	contract, err := bindIPicoVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPicoVerifierTransactor{contract: contract}, nil
}

// NewIPicoVerifierFilterer creates a new log filterer instance of IPicoVerifier, bound to a specific deployed contract.
func NewIPicoVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IPicoVerifierFilterer, error) {
	contract, err := bindIPicoVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPicoVerifierFilterer{contract: contract}, nil
}

// bindIPicoVerifier binds a generic wrapper to an already deployed contract.
func bindIPicoVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPicoVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPicoVerifier *IPicoVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPicoVerifier.Contract.IPicoVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPicoVerifier *IPicoVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPicoVerifier.Contract.IPicoVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPicoVerifier *IPicoVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPicoVerifier.Contract.IPicoVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPicoVerifier *IPicoVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPicoVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPicoVerifier *IPicoVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPicoVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPicoVerifier *IPicoVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPicoVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierCaller) VerifyPicoProof(opts *bind.CallOpts, riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	var out []interface{}
	err := _IPicoVerifier.contract.Call(opts, &out, "verifyPicoProof", riscvVkey, publicValues, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierSession) VerifyPicoProof(riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	return _IPicoVerifier.Contract.VerifyPicoProof(&_IPicoVerifier.CallOpts, riscvVkey, publicValues, proof)
}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierCallerSession) VerifyPicoProof(riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	return _IPicoVerifier.Contract.VerifyPicoProof(&_IPicoVerifier.CallOpts, riscvVkey, publicValues, proof)
}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesHash, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierCaller) VerifyPicoProof0(opts *bind.CallOpts, riscvVkey [32]byte, publicValuesHash [32]byte, proof [8]*big.Int) error {
	var out []interface{}
	err := _IPicoVerifier.contract.Call(opts, &out, "verifyPicoProof0", riscvVkey, publicValuesHash, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesHash, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierSession) VerifyPicoProof0(riscvVkey [32]byte, publicValuesHash [32]byte, proof [8]*big.Int) error {
	return _IPicoVerifier.Contract.VerifyPicoProof0(&_IPicoVerifier.CallOpts, riscvVkey, publicValuesHash, proof)
}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesHash, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierCallerSession) VerifyPicoProof0(riscvVkey [32]byte, publicValuesHash [32]byte, proof [8]*big.Int) error {
	return _IPicoVerifier.Contract.VerifyPicoProof0(&_IPicoVerifier.CallOpts, riscvVkey, publicValuesHash, proof)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PicoVerifierMetaData contains all meta data concerning the PicoVerifier contract.
var PicoVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"}],\"name\":\"sha256PublicValues\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"riscvVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"verifyPicoProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"riscvVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"verifyPicoProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610016576113f9908161001c8239f35b600080fdfe6080604081815260048036101561001557600080fd5b600092833560e01c90816344f6369214610ba9575080635fe24f2314610798578063dbd7536514610735578063ee87de3714610657578063f11817b2146101335763f39751a81461006557600080fd5b82913461012f576101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261012f57366101441161012f578051916100ac83610d62565b81368437803583526024356020840152303b1561012a576100fa928491835194859283927f5fe24f230000000000000000000000000000000000000000000000000000000084528301610d7e565b0381305afa908115610121575061010e5750f35b61011790610d1f565b61011e5780f35b80fd5b513d84823e3d90fd5b505050fd5b5050fd5b5090346106535760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126106535736608411610653573660c4116106535761017e823561122e565b61018c602435604435611295565b91929061019a60643561122e565b91885194898601947f21479a04cf50085fc3d118540aa4416c0548c633cfb86f697bc20fba235f249b87526001878c60208201987f0e0de325f841aafcf07bf6c6e89c29396362176a1e8ff6c67a4eee690ead0ae68a527f23e3faf88e315d58db00740d43a122ed4423a51351445140655e2763c6e5915f8152818060608501927f0cacfb35b768001e7e9de1cd30202a17d17328a4e1471e14f9ea95445908e48684527f11d4df3cee8e649c7f4f47eab2f87cca5238ca416ccb5f5b00d125b5e9efd79d6084356080880197818952848460608160075afa907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000196816080899360065afa92101616947f198ef1069f1cb584addcc1615c0154cce48a0f0ce946734e20cd0182dc80cd6783525260a43580965260608160075afa928c60808160065afa93101616161696519551961561062b578a5197610300998a8a019b67ffffffffffffffff9c8d8c8210908211176105fc578e528b368c378a5260208a01528b8901526060880152608087015260a086015260c085015260e08401527f10592dad298bacacd320a6d53b0aa6d19634663e74c9d0904d592f0f4c76c23f6101008401527f2e3568a669884dace0f7548ac91acb71f8e8ad34ad9fd989dc345cba33dc3c286101208401527f03bd0098c28f2dd6fcd74b88320067acb7fb385383edbd1585f436f84f348a216101408401527f2318952ecdb575b246b21b9dfe4d570c02afc29e92e969d1212e219f3b51998b6101608401527f1f905760a83cfc22560ae2e55747618102f280cdc3e0cd8a94d8dfbe209d65086101808401527f22aba4b0d677b5afb88a6a097d18a9ea0a583d8b0c0a3913090ab78cb3afde1a6101a08401527f2f41b4486c54a5721dd60b9097cf061e4465bf1b3c9c0a289301c0ace899bc556101c08401527f276d99ea794fdcb1475195104f95f535e6efb16aabdcc254754d5a930945a55a6101e08401527f17405def8568fbfa36af7ab9447a140709e550582489a4929276627df3dd97dd6102008401527f06d1101f5d40d84bab0cbb907522ec60be2d00b38333cd6652cb0070104e57276102208401526102408301526102608201527f234fc0b9aed3a32ea89c2325cf53add49651b855829cfcc9a1368bba013d9b786102808201527f1d1a6cf05f3129136408d8dc1ad2d8e6f60e087f8177f14b2deb7bbec4a37c146102a08201527f13775647cd78947d1d00fd5c997841ab39cb4f7723f375f48806fc067a5464d66102c08201527f12567ee8d519b6290e2c305a440418bc61b6ec775398a77bd9a5bd526904652c6102e08201528351926020840190848210908211176105ce57845260209183918336843760085afa159081156105c1575b5061059b578280f35b517f7fcdd1f4000000000000000000000000000000000000000000000000000000008152fd5b6001915051141538610592565b6041867f4e487b71000000000000000000000000000000000000000000000000000000006000525260246000fd5b5060418f7f4e487b71000000000000000000000000000000000000000000000000000000006000525260246000fd5b8b8b517fa54f8e27000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b509034610653576101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610653578260243567ffffffffffffffff8111610731576106a99036908501610c99565b3661014411610653576106bb91610ccc565b8251906106c782610d62565b83368337843582526020820152303b156107315761071093835194859283927f5fe24f230000000000000000000000000000000000000000000000000000000084528301610d7e565b0381305afa9081156101215750610725575080f35b61072e90610d1f565b80f35b5080fd5b50913461011e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261011e5782359067ffffffffffffffff821161011e575061078b60209361079192369101610c99565b90610ccc565b9051908152f35b50903461065357610140807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610ba557610104368111610ba1573661014411610ba15782519060018483017f21479a04cf50085fc3d118540aa4416c0548c633cfb86f697bc20fba235f249b845260208401927f0e0de325f841aafcf07bf6c6e89c29396362176a1e8ff6c67a4eee690ead0ae684527f23e3faf88e315d58db00740d43a122ed4423a51351445140655e2763c6e5915f825260608501907f0cacfb35b768001e7e9de1cd30202a17d17328a4e1471e14f9ea95445908e486825235908760808701938385527f11d4df3cee8e649c7f4f47eab2f87cca5238ca416ccb5f5b00d125b5e9efd79d828260608160075afa947f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019586858c60808160065afa92101616937f198ef1069f1cb584addcc1615c0154cce48a0f0ce946734e20cd0182dc80cd678352526101243580955260608160075afa91888760808160065afa931016161616915190519115610b78577f03bd0098c28f2dd6fcd74b88320067acb7fb385383edbd1585f436f84f348a218451937f10592dad298bacacd320a6d53b0aa6d19634663e74c9d0904d592f0f4c76c23f610100808988378601527f2e3568a669884dace0f7548ac91acb71f8e8ad34ad9fd989dc345cba33dc3c286101208601528401527f2318952ecdb575b246b21b9dfe4d570c02afc29e92e969d1212e219f3b51998b6101608401527f1f905760a83cfc22560ae2e55747618102f280cdc3e0cd8a94d8dfbe209d65086101808401527f22aba4b0d677b5afb88a6a097d18a9ea0a583d8b0c0a3913090ab78cb3afde1a6101a08401527f2f41b4486c54a5721dd60b9097cf061e4465bf1b3c9c0a289301c0ace899bc556101c08401527f276d99ea794fdcb1475195104f95f535e6efb16aabdcc254754d5a930945a55a6101e08401527f17405def8568fbfa36af7ab9447a140709e550582489a4929276627df3dd97dd6102008401527f06d1101f5d40d84bab0cbb907522ec60be2d00b38333cd6652cb0070104e57276102208401526102408301526102608201527f234fc0b9aed3a32ea89c2325cf53add49651b855829cfcc9a1368bba013d9b786102808201527f1d1a6cf05f3129136408d8dc1ad2d8e6f60e087f8177f14b2deb7bbec4a37c146102a08201527f13775647cd78947d1d00fd5c997841ab39cb4f7723f375f48806fc067a5464d66102c08201527f12567ee8d519b6290e2c305a440418bc61b6ec775398a77bd9a5bd526904652c6102e08201526020816103008160085afa9051161561059b578280f35b505050517fa54f8e27000000000000000000000000000000000000000000000000000000008152fd5b8480fd5b8380fd5b8284863461011e576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261011e57366101041161011e576080840184811067ffffffffffffffff821117610c6d578252926080368237610c116024358435610db9565b8152610c2760843560a435604435606435610f22565b92906020938484015281830152610c4260e43560c435610db9565b6060830152519390845b848310610c5857608086f35b83806001928451815201920192019190610c4c565b6024826041867f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b9181601f84011215610cc75782359167ffffffffffffffff8311610cc75760208381860195010111610cc757565b600080fd5b6020916000918160405192839283378101838152039060025afa15610d13577f1fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6000511690565b6040513d6000823e3d90fd5b67ffffffffffffffff8111610d3357604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117610d3357604052565b919061014083019261010090816044823701906000915b60028310610da257505050565b600190825181526020809101920192019190610d95565b907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47808310801590610e67575b610e2d57821580610e5f575b610e5757610e098160038186818180090908610e71565b828103610e195750505060011b90565b81900681030603610e2d57600190811b1790565b60046040517f7fcdd1f4000000000000000000000000000000000000000000000000000000008152fd5b505050600090565b508115610df2565b5080821015610de6565b90610e7b82610ea7565b917f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4783800903610e2d57565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760a083015260208260c08160055afa91519115610e2d57565b909293917f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4794858310801590611104575b80156110fa575b80156110f0575b610e2d5780828685171717156110e4578561105581808080808981808e9f8f9e9f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448280928709099e8f0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089b09818d8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775080681030695827f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161103c81808c80098187800908610e71565b8408098361104982610ea7565b8009141596879161110e565b9390808214806110db575b1561108d57505050505090600090600014611085575060ff60025b169060021b179190565b60ff9061107b565b828091068103061492836110ca575b505050600014610e2d57600191600090156110c2575060ff60025b169060021b17179190565b60ff906110b7565b81929350068103061438808061109c565b50848414611060565b50600094508493505050565b5085811015610f61565b5085821015610f5a565b5085851015610f53565b9161117e7f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4919492947f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47928380928161116f81808c8009818c800908610e71565b91611222575b50870809610e71565b9381600286096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301528360a083015260208260c08160055afa91519115610e2d5783826001920903610e2d57829082099382808080888009068103068188800908149182159261120f575b5050610e2d57565b8091925084860960020914153880611207565b80910681030681611175565b801561128c578060011c917f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780841015610e2d576001806112788360038189818180090908610e71565b9416146112825750565b8091920681030690565b50600090600090565b9190918015806113bb575b6113ae578060021c9280927f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47918286108015906113a4575b610e2d578584848080858180808080600280827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4461137b9f8f90839109099b16149a8a0981868181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508068103069409818b8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50861110e565b909281936001808297161461138f57505050565b81900681038190069450908190068103069150565b50828110156112d8565b5060009150819081908190565b5082156112a056fea2646970667358221220b946ff44f2c385b79caca448513516b95da8f53fe9086dbc5ec0422415f8f8ed64736f6c63430008140033",
}

// PicoVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use PicoVerifierMetaData.ABI instead.
var PicoVerifierABI = PicoVerifierMetaData.ABI

// PicoVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PicoVerifierMetaData.Bin instead.
var PicoVerifierBin = PicoVerifierMetaData.Bin

// DeployPicoVerifier deploys a new Ethereum contract, binding an instance of PicoVerifier to it.
func DeployPicoVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PicoVerifier, error) {
	parsed, err := PicoVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PicoVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PicoVerifier{PicoVerifierCaller: PicoVerifierCaller{contract: contract}, PicoVerifierTransactor: PicoVerifierTransactor{contract: contract}, PicoVerifierFilterer: PicoVerifierFilterer{contract: contract}}, nil
}

// PicoVerifier is an auto generated Go binding around an Ethereum contract.
type PicoVerifier struct {
	PicoVerifierCaller     // Read-only binding to the contract
	PicoVerifierTransactor // Write-only binding to the contract
	PicoVerifierFilterer   // Log filterer for contract events
}

// PicoVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PicoVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PicoVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PicoVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PicoVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PicoVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PicoVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PicoVerifierSession struct {
	Contract     *PicoVerifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PicoVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PicoVerifierCallerSession struct {
	Contract *PicoVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PicoVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PicoVerifierTransactorSession struct {
	Contract     *PicoVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PicoVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PicoVerifierRaw struct {
	Contract *PicoVerifier // Generic contract binding to access the raw methods on
}

// PicoVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PicoVerifierCallerRaw struct {
	Contract *PicoVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// PicoVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PicoVerifierTransactorRaw struct {
	Contract *PicoVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPicoVerifier creates a new instance of PicoVerifier, bound to a specific deployed contract.
func NewPicoVerifier(address common.Address, backend bind.ContractBackend) (*PicoVerifier, error) {
	contract, err := bindPicoVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PicoVerifier{PicoVerifierCaller: PicoVerifierCaller{contract: contract}, PicoVerifierTransactor: PicoVerifierTransactor{contract: contract}, PicoVerifierFilterer: PicoVerifierFilterer{contract: contract}}, nil
}

// NewPicoVerifierCaller creates a new read-only instance of PicoVerifier, bound to a specific deployed contract.
func NewPicoVerifierCaller(address common.Address, caller bind.ContractCaller) (*PicoVerifierCaller, error) {
	contract, err := bindPicoVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PicoVerifierCaller{contract: contract}, nil
}

// NewPicoVerifierTransactor creates a new write-only instance of PicoVerifier, bound to a specific deployed contract.
func NewPicoVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PicoVerifierTransactor, error) {
	contract, err := bindPicoVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PicoVerifierTransactor{contract: contract}, nil
}

// NewPicoVerifierFilterer creates a new log filterer instance of PicoVerifier, bound to a specific deployed contract.
func NewPicoVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PicoVerifierFilterer, error) {
	contract, err := bindPicoVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PicoVerifierFilterer{contract: contract}, nil
}

// bindPicoVerifier binds a generic wrapper to an already deployed contract.
func bindPicoVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PicoVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PicoVerifier *PicoVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PicoVerifier.Contract.PicoVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PicoVerifier *PicoVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PicoVerifier.Contract.PicoVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PicoVerifier *PicoVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PicoVerifier.Contract.PicoVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PicoVerifier *PicoVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PicoVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PicoVerifier *PicoVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PicoVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PicoVerifier *PicoVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PicoVerifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PicoVerifier *PicoVerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _PicoVerifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PicoVerifier *PicoVerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _PicoVerifier.Contract.CompressProof(&_PicoVerifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_PicoVerifier *PicoVerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _PicoVerifier.Contract.CompressProof(&_PicoVerifier.CallOpts, proof)
}

// Sha256PublicValues is a free data retrieval call binding the contract method 0xdbd75365.
//
// Solidity: function sha256PublicValues(bytes publicValues) pure returns(bytes32)
func (_PicoVerifier *PicoVerifierCaller) Sha256PublicValues(opts *bind.CallOpts, publicValues []byte) ([32]byte, error) {
	var out []interface{}
	err := _PicoVerifier.contract.Call(opts, &out, "sha256PublicValues", publicValues)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sha256PublicValues is a free data retrieval call binding the contract method 0xdbd75365.
//
// Solidity: function sha256PublicValues(bytes publicValues) pure returns(bytes32)
func (_PicoVerifier *PicoVerifierSession) Sha256PublicValues(publicValues []byte) ([32]byte, error) {
	return _PicoVerifier.Contract.Sha256PublicValues(&_PicoVerifier.CallOpts, publicValues)
}

// Sha256PublicValues is a free data retrieval call binding the contract method 0xdbd75365.
//
// Solidity: function sha256PublicValues(bytes publicValues) pure returns(bytes32)
func (_PicoVerifier *PicoVerifierCallerSession) Sha256PublicValues(publicValues []byte) ([32]byte, error) {
	return _PicoVerifier.Contract.Sha256PublicValues(&_PicoVerifier.CallOpts, publicValues)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_PicoVerifier *PicoVerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [2]*big.Int) error {
	var out []interface{}
	err := _PicoVerifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_PicoVerifier *PicoVerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [2]*big.Int) error {
	return _PicoVerifier.Contract.VerifyCompressedProof(&_PicoVerifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_PicoVerifier *PicoVerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [2]*big.Int) error {
	return _PicoVerifier.Contract.VerifyCompressedProof(&_PicoVerifier.CallOpts, compressedProof, input)
}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_PicoVerifier *PicoVerifierCaller) VerifyPicoProof(opts *bind.CallOpts, riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	var out []interface{}
	err := _PicoVerifier.contract.Call(opts, &out, "verifyPicoProof", riscvVkey, publicValues, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_PicoVerifier *PicoVerifierSession) VerifyPicoProof(riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	return _PicoVerifier.Contract.VerifyPicoProof(&_PicoVerifier.CallOpts, riscvVkey, publicValues, proof)
}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_PicoVerifier *PicoVerifierCallerSession) VerifyPicoProof(riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	return _PicoVerifier.Contract.VerifyPicoProof(&_PicoVerifier.CallOpts, riscvVkey, publicValues, proof)
}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesDigest, uint256[8] proof) view returns()
func (_PicoVerifier *PicoVerifierCaller) VerifyPicoProof0(opts *bind.CallOpts, riscvVkey [32]byte, publicValuesDigest [32]byte, proof [8]*big.Int) error {
	var out []interface{}
	err := _PicoVerifier.contract.Call(opts, &out, "verifyPicoProof0", riscvVkey, publicValuesDigest, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesDigest, uint256[8] proof) view returns()
func (_PicoVerifier *PicoVerifierSession) VerifyPicoProof0(riscvVkey [32]byte, publicValuesDigest [32]byte, proof [8]*big.Int) error {
	return _PicoVerifier.Contract.VerifyPicoProof0(&_PicoVerifier.CallOpts, riscvVkey, publicValuesDigest, proof)
}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesDigest, uint256[8] proof) view returns()
func (_PicoVerifier *PicoVerifierCallerSession) VerifyPicoProof0(riscvVkey [32]byte, publicValuesDigest [32]byte, proof [8]*big.Int) error {
	return _PicoVerifier.Contract.VerifyPicoProof0(&_PicoVerifier.CallOpts, riscvVkey, publicValuesDigest, proof)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_PicoVerifier *PicoVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [2]*big.Int) error {
	var out []interface{}
	err := _PicoVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_PicoVerifier *PicoVerifierSession) VerifyProof(proof [8]*big.Int, input [2]*big.Int) error {
	return _PicoVerifier.Contract.VerifyProof(&_PicoVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_PicoVerifier *PicoVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [2]*big.Int) error {
	return _PicoVerifier.Contract.VerifyProof(&_PicoVerifier.CallOpts, proof, input)
}

// VerifierMetaData contains all meta data concerning the Verifier contract.
var VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610016576110a8908161001c8239f35b600080fdfe608060408181526004908136101561001657600080fd5b600092833560e01c90816344f6369214610975575080635fe24f23146105655763f11817b21461004557600080fd5b346105615760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126105615736608411610561573660c4116105615761008e8235610edd565b61009c602435604435610f44565b9192906100aa606435610edd565b91885194898601947f21479a04cf50085fc3d118540aa4416c0548c633cfb86f697bc20fba235f249b87526001878c60208201987f0e0de325f841aafcf07bf6c6e89c29396362176a1e8ff6c67a4eee690ead0ae68a527f23e3faf88e315d58db00740d43a122ed4423a51351445140655e2763c6e5915f8152818060608501927f0cacfb35b768001e7e9de1cd30202a17d17328a4e1471e14f9ea95445908e48684527f11d4df3cee8e649c7f4f47eab2f87cca5238ca416ccb5f5b00d125b5e9efd79d6084356080880197818952848460608160075afa907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000196816080899360065afa92101616947f198ef1069f1cb584addcc1615c0154cce48a0f0ce946734e20cd0182dc80cd6783525260a43580965260608160075afa928c60808160065afa931016161616965195519615610539578a5197610300998a8a019b67ffffffffffffffff9c8d8c82109082111761050a578e528b368c378a5260208a01528b8901526060880152608087015260a086015260c085015260e08401527f10592dad298bacacd320a6d53b0aa6d19634663e74c9d0904d592f0f4c76c23f6101008401527f2e3568a669884dace0f7548ac91acb71f8e8ad34ad9fd989dc345cba33dc3c286101208401527f03bd0098c28f2dd6fcd74b88320067acb7fb385383edbd1585f436f84f348a216101408401527f2318952ecdb575b246b21b9dfe4d570c02afc29e92e969d1212e219f3b51998b6101608401527f1f905760a83cfc22560ae2e55747618102f280cdc3e0cd8a94d8dfbe209d65086101808401527f22aba4b0d677b5afb88a6a097d18a9ea0a583d8b0c0a3913090ab78cb3afde1a6101a08401527f2f41b4486c54a5721dd60b9097cf061e4465bf1b3c9c0a289301c0ace899bc556101c08401527f276d99ea794fdcb1475195104f95f535e6efb16aabdcc254754d5a930945a55a6101e08401527f17405def8568fbfa36af7ab9447a140709e550582489a4929276627df3dd97dd6102008401527f06d1101f5d40d84bab0cbb907522ec60be2d00b38333cd6652cb0070104e57276102208401526102408301526102608201527f234fc0b9aed3a32ea89c2325cf53add49651b855829cfcc9a1368bba013d9b786102808201527f1d1a6cf05f3129136408d8dc1ad2d8e6f60e087f8177f14b2deb7bbec4a37c146102a08201527f13775647cd78947d1d00fd5c997841ab39cb4f7723f375f48806fc067a5464d66102c08201527f12567ee8d519b6290e2c305a440418bc61b6ec775398a77bd9a5bd526904652c6102e08201528351926020840190848210908211176104de57845260209183918336843760085afa159081156104d1575b506104ab578280f35b517f7fcdd1f4000000000000000000000000000000000000000000000000000000008152fd5b60019150511415386104a2565b6024876041887f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b50505060248e60418f7f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b8b8b517fa54f8e27000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b503461056157610140807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126109715761010436811161096d57366101441161096d5782519060018483017f21479a04cf50085fc3d118540aa4416c0548c633cfb86f697bc20fba235f249b845260208401927f0e0de325f841aafcf07bf6c6e89c29396362176a1e8ff6c67a4eee690ead0ae684527f23e3faf88e315d58db00740d43a122ed4423a51351445140655e2763c6e5915f825260608501907f0cacfb35b768001e7e9de1cd30202a17d17328a4e1471e14f9ea95445908e486825235908760808701938385527f11d4df3cee8e649c7f4f47eab2f87cca5238ca416ccb5f5b00d125b5e9efd79d828260608160075afa947f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000019586858c60808160065afa92101616937f198ef1069f1cb584addcc1615c0154cce48a0f0ce946734e20cd0182dc80cd678352526101243580955260608160075afa91888760808160065afa931016161616915190519115610944577f03bd0098c28f2dd6fcd74b88320067acb7fb385383edbd1585f436f84f348a218451937f10592dad298bacacd320a6d53b0aa6d19634663e74c9d0904d592f0f4c76c23f610100808988378601527f2e3568a669884dace0f7548ac91acb71f8e8ad34ad9fd989dc345cba33dc3c286101208601528401527f2318952ecdb575b246b21b9dfe4d570c02afc29e92e969d1212e219f3b51998b6101608401527f1f905760a83cfc22560ae2e55747618102f280cdc3e0cd8a94d8dfbe209d65086101808401527f22aba4b0d677b5afb88a6a097d18a9ea0a583d8b0c0a3913090ab78cb3afde1a6101a08401527f2f41b4486c54a5721dd60b9097cf061e4465bf1b3c9c0a289301c0ace899bc556101c08401527f276d99ea794fdcb1475195104f95f535e6efb16aabdcc254754d5a930945a55a6101e08401527f17405def8568fbfa36af7ab9447a140709e550582489a4929276627df3dd97dd6102008401527f06d1101f5d40d84bab0cbb907522ec60be2d00b38333cd6652cb0070104e57276102208401526102408301526102608201527f234fc0b9aed3a32ea89c2325cf53add49651b855829cfcc9a1368bba013d9b786102808201527f1d1a6cf05f3129136408d8dc1ad2d8e6f60e087f8177f14b2deb7bbec4a37c146102a08201527f13775647cd78947d1d00fd5c997841ab39cb4f7723f375f48806fc067a5464d66102c08201527f12567ee8d519b6290e2c305a440418bc61b6ec775398a77bd9a5bd526904652c6102e08201526020816103008160085afa905116156104ab578280f35b505050517fa54f8e27000000000000000000000000000000000000000000000000000000008152fd5b8480fd5b8380fd5b93905034610a65576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610a65573661010411610a65576080840184811067ffffffffffffffff821117610a395782529260803682376109dd6024358435610a68565b81526109f360843560a435604435606435610bd1565b92906020938484015281830152610a0e60e43560c435610a68565b6060830152519390845b848310610a2457608086f35b83806001928451815201920192019190610a18565b6024826041867f4e487b7100000000000000000000000000000000000000000000000000000000835252fd5b80fd5b907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47808310801590610b16575b610adc57821580610b0e575b610b0657610ab88160038186818180090908610b20565b828103610ac85750505060011b90565b81900681030603610adc57600190811b1790565b60046040517f7fcdd1f4000000000000000000000000000000000000000000000000000000008152fd5b505050600090565b508115610aa1565b5080821015610a95565b90610b2a82610b56565b917f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4783800903610adc57565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760a083015260208260c08160055afa91519115610adc57565b909293917f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4794858310801590610db3575b8015610da9575b8015610d9f575b610adc578082868517171715610d935785610d0481808080808981808e9f8f9e9f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448280928709099e8f0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089b09818d8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775080681030695827f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481610ceb81808c80098187800908610b20565b84080983610cf882610b56565b80091415968791610dbd565b939080821480610d8a575b15610d3c57505050505090600090600014610d34575060ff60025b169060021b179190565b60ff90610d2a565b82809106810306149283610d79575b505050600014610adc5760019160009015610d71575060ff60025b169060021b17179190565b60ff90610d66565b819293500681030614388080610d4b565b50848414610d0f565b50600094508493505050565b5085811015610c10565b5085821015610c09565b5085851015610c02565b91610e2d7f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4919492947f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd479283809281610e1e81808c8009818c800908610b20565b91610ed1575b50870809610b20565b9381600286096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301528360a083015260208260c08160055afa91519115610adc5783826001920903610adc578290820993828080808880090681030681888009081491821592610ebe575b5050610adc57565b8091925084860960020914153880610eb6565b80910681030681610e24565b8015610f3b578060011c917f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780841015610adc57600180610f278360038189818180090908610b20565b941614610f315750565b8091920681030690565b50600090600090565b91909180158061106a575b61105d578060021c9280927f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4791828610801590611053575b610adc578584848080858180808080600280827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4461102a9f8f90839109099b16149a8a0981868181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508068103069409818b8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508610dbd565b909281936001808297161461103e57505050565b81900681038190069450908190068103069150565b5082811015610f87565b5060009150819081908190565b508215610f4f56fea26469706673582212206e6cdafb9e915626301e1f12da3807211d243781f6a63bd6f01341777a5bcf5564736f6c63430008140033",
}

// VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierMetaData.ABI instead.
var VerifierABI = VerifierMetaData.ABI

// VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifierMetaData.Bin instead.
var VerifierBin = VerifierMetaData.Bin

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_Verifier *VerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_Verifier *VerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _Verifier.Contract.CompressProof(&_Verifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_Verifier *VerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _Verifier.Contract.CompressProof(&_Verifier.CallOpts, proof)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_Verifier *VerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [2]*big.Int) error {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_Verifier *VerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [2]*big.Int) error {
	return _Verifier.Contract.VerifyCompressedProof(&_Verifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_Verifier *VerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [2]*big.Int) error {
	return _Verifier.Contract.VerifyCompressedProof(&_Verifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [2]*big.Int) error {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_Verifier *VerifierSession) VerifyProof(proof [8]*big.Int, input [2]*big.Int) error {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_Verifier *VerifierCallerSession) VerifyProof(proof [8]*big.Int, input [2]*big.Int) error {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, proof, input)
}
