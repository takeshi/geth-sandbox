// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ApproveAndCallFallBackABI is the input ABI used to generate the binding from.
const ApproveAndCallFallBackABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ApproveAndCallFallBackBin is the compiled bytecode used for deploying new contracts.
const ApproveAndCallFallBackBin = `0x`

// DeployApproveAndCallFallBack deploys a new Ethereum contract, binding an instance of ApproveAndCallFallBack to it.
func DeployApproveAndCallFallBack(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ApproveAndCallFallBack, error) {
	parsed, err := abi.JSON(strings.NewReader(ApproveAndCallFallBackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApproveAndCallFallBackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ApproveAndCallFallBack{ApproveAndCallFallBackCaller: ApproveAndCallFallBackCaller{contract: contract}, ApproveAndCallFallBackTransactor: ApproveAndCallFallBackTransactor{contract: contract}, ApproveAndCallFallBackFilterer: ApproveAndCallFallBackFilterer{contract: contract}}, nil
}

// ApproveAndCallFallBack is an auto generated Go binding around an Ethereum contract.
type ApproveAndCallFallBack struct {
	ApproveAndCallFallBackCaller     // Read-only binding to the contract
	ApproveAndCallFallBackTransactor // Write-only binding to the contract
	ApproveAndCallFallBackFilterer   // Log filterer for contract events
}

// ApproveAndCallFallBackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApproveAndCallFallBackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveAndCallFallBackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApproveAndCallFallBackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveAndCallFallBackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApproveAndCallFallBackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproveAndCallFallBackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApproveAndCallFallBackSession struct {
	Contract     *ApproveAndCallFallBack // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ApproveAndCallFallBackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApproveAndCallFallBackCallerSession struct {
	Contract *ApproveAndCallFallBackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ApproveAndCallFallBackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApproveAndCallFallBackTransactorSession struct {
	Contract     *ApproveAndCallFallBackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ApproveAndCallFallBackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApproveAndCallFallBackRaw struct {
	Contract *ApproveAndCallFallBack // Generic contract binding to access the raw methods on
}

// ApproveAndCallFallBackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApproveAndCallFallBackCallerRaw struct {
	Contract *ApproveAndCallFallBackCaller // Generic read-only contract binding to access the raw methods on
}

// ApproveAndCallFallBackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApproveAndCallFallBackTransactorRaw struct {
	Contract *ApproveAndCallFallBackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApproveAndCallFallBack creates a new instance of ApproveAndCallFallBack, bound to a specific deployed contract.
func NewApproveAndCallFallBack(address common.Address, backend bind.ContractBackend) (*ApproveAndCallFallBack, error) {
	contract, err := bindApproveAndCallFallBack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApproveAndCallFallBack{ApproveAndCallFallBackCaller: ApproveAndCallFallBackCaller{contract: contract}, ApproveAndCallFallBackTransactor: ApproveAndCallFallBackTransactor{contract: contract}, ApproveAndCallFallBackFilterer: ApproveAndCallFallBackFilterer{contract: contract}}, nil
}

// NewApproveAndCallFallBackCaller creates a new read-only instance of ApproveAndCallFallBack, bound to a specific deployed contract.
func NewApproveAndCallFallBackCaller(address common.Address, caller bind.ContractCaller) (*ApproveAndCallFallBackCaller, error) {
	contract, err := bindApproveAndCallFallBack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApproveAndCallFallBackCaller{contract: contract}, nil
}

// NewApproveAndCallFallBackTransactor creates a new write-only instance of ApproveAndCallFallBack, bound to a specific deployed contract.
func NewApproveAndCallFallBackTransactor(address common.Address, transactor bind.ContractTransactor) (*ApproveAndCallFallBackTransactor, error) {
	contract, err := bindApproveAndCallFallBack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApproveAndCallFallBackTransactor{contract: contract}, nil
}

// NewApproveAndCallFallBackFilterer creates a new log filterer instance of ApproveAndCallFallBack, bound to a specific deployed contract.
func NewApproveAndCallFallBackFilterer(address common.Address, filterer bind.ContractFilterer) (*ApproveAndCallFallBackFilterer, error) {
	contract, err := bindApproveAndCallFallBack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApproveAndCallFallBackFilterer{contract: contract}, nil
}

// bindApproveAndCallFallBack binds a generic wrapper to an already deployed contract.
func bindApproveAndCallFallBack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApproveAndCallFallBackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ApproveAndCallFallBack.Contract.ApproveAndCallFallBackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.ApproveAndCallFallBackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.ApproveAndCallFallBackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ApproveAndCallFallBack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApproveAndCallFallBack *ApproveAndCallFallBackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.contract.Transact(opts, method, params...)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, tokens uint256, token address, data bytes) returns()
func (_ApproveAndCallFallBack *ApproveAndCallFallBackTransactor) ReceiveApproval(opts *bind.TransactOpts, from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.contract.Transact(opts, "receiveApproval", from, tokens, token, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, tokens uint256, token address, data bytes) returns()
func (_ApproveAndCallFallBack *ApproveAndCallFallBackSession) ReceiveApproval(from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.ReceiveApproval(&_ApproveAndCallFallBack.TransactOpts, from, tokens, token, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(from address, tokens uint256, token address, data bytes) returns()
func (_ApproveAndCallFallBack *ApproveAndCallFallBackTransactorSession) ReceiveApproval(from common.Address, tokens *big.Int, token common.Address, data []byte) (*types.Transaction, error) {
	return _ApproveAndCallFallBack.Contract.ReceiveApproval(&_ApproveAndCallFallBack.TransactOpts, from, tokens, token, data)
}

// ERC20InterfaceABI is the input ABI used to generate the binding from.
const ERC20InterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20InterfaceBin is the compiled bytecode used for deploying new contracts.
const ERC20InterfaceBin = `0x`

// DeployERC20Interface deploys a new Ethereum contract, binding an instance of ERC20Interface to it.
func DeployERC20Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Interface{ERC20InterfaceCaller: ERC20InterfaceCaller{contract: contract}, ERC20InterfaceTransactor: ERC20InterfaceTransactor{contract: contract}, ERC20InterfaceFilterer: ERC20InterfaceFilterer{contract: contract}}, nil
}

// ERC20Interface is an auto generated Go binding around an Ethereum contract.
type ERC20Interface struct {
	ERC20InterfaceCaller     // Read-only binding to the contract
	ERC20InterfaceTransactor // Write-only binding to the contract
	ERC20InterfaceFilterer   // Log filterer for contract events
}

// ERC20InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20InterfaceSession struct {
	Contract     *ERC20Interface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20InterfaceCallerSession struct {
	Contract *ERC20InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20InterfaceTransactorSession struct {
	Contract     *ERC20InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20InterfaceRaw struct {
	Contract *ERC20Interface // Generic contract binding to access the raw methods on
}

// ERC20InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20InterfaceCallerRaw struct {
	Contract *ERC20InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactorRaw struct {
	Contract *ERC20InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Interface creates a new instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20Interface(address common.Address, backend bind.ContractBackend) (*ERC20Interface, error) {
	contract, err := bindERC20Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Interface{ERC20InterfaceCaller: ERC20InterfaceCaller{contract: contract}, ERC20InterfaceTransactor: ERC20InterfaceTransactor{contract: contract}, ERC20InterfaceFilterer: ERC20InterfaceFilterer{contract: contract}}, nil
}

// NewERC20InterfaceCaller creates a new read-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC20InterfaceCaller, error) {
	contract, err := bindERC20Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceCaller{contract: contract}, nil
}

// NewERC20InterfaceTransactor creates a new write-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20InterfaceTransactor, error) {
	contract, err := bindERC20Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceTransactor{contract: contract}, nil
}

// NewERC20InterfaceFilterer creates a new log filterer instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20InterfaceFilterer, error) {
	contract, err := bindERC20Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceFilterer{contract: contract}, nil
}

// bindERC20Interface binds a generic wrapper to an already deployed contract.
func bindERC20Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.ERC20InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_ERC20Interface *ERC20InterfaceCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_ERC20Interface *ERC20InterfaceSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.Allowance(&_ERC20Interface.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.Allowance(&_ERC20Interface.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_ERC20Interface *ERC20InterfaceCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_ERC20Interface *ERC20InterfaceSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.BalanceOf(&_ERC20Interface.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.BalanceOf(&_ERC20Interface.CallOpts, tokenOwner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceSession) TotalSupply() (*big.Int, error) {
	return _ERC20Interface.Contract.TotalSupply(&_ERC20Interface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Interface.Contract.TotalSupply(&_ERC20Interface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Approve(&_ERC20Interface.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Approve(&_ERC20Interface.TransactOpts, spender, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, from, to, tokens)
}

// ERC20InterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Interface contract.
type ERC20InterfaceApprovalIterator struct {
	Event *ERC20InterfaceApproval // Event containing the contract specifics and raw log

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
func (it *ERC20InterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterfaceApproval)
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
		it.Event = new(ERC20InterfaceApproval)
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
func (it *ERC20InterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterfaceApproval represents a Approval event raised by the ERC20Interface contract.
type ERC20InterfaceApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*ERC20InterfaceApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Interface.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceApprovalIterator{contract: _ERC20Interface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20InterfaceApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Interface.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterfaceApproval)
				if err := _ERC20Interface.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20InterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Interface contract.
type ERC20InterfaceTransferIterator struct {
	Event *ERC20InterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20InterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterfaceTransfer)
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
		it.Event = new(ERC20InterfaceTransfer)
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
func (it *ERC20InterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterfaceTransfer represents a Transfer event raised by the ERC20Interface contract.
type ERC20InterfaceTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, tokens uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20InterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Interface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceTransferIterator{contract: _ERC20Interface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, tokens uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20InterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Interface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterfaceTransfer)
				if err := _ERC20Interface.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// FixedSupplyTokenABI is the input ABI used to generate the binding from.
const FixedSupplyTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// FixedSupplyTokenBin is the compiled bytecode used for deploying new contracts.
const FixedSupplyTokenBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a03191633600160a060020a031617905560408051908101604052600581527f464958454400000000000000000000000000000000000000000000000000000060208201526002908051610071929160200190610135565b5060408051908101604052601a81527f4578616d706c6520466978656420537570706c7920546f6b656e000000000000602082015260039080516100b9929160200190610135565b5060048054601260ff19909116179081905560ff16600a0a620f424002600581905560008054600160a060020a039081168252600660205260408083208490558254909116927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef915190815260200160405180910390a36101d0565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061017657805160ff19168380011785556101a3565b828001600101855582156101a3579182015b828111156101a3578251825591602001919060010190610188565b506101af9291506101b3565b5090565b6101cd91905b808211156101af57600081556001016101b9565b90565b610ab0806101df6000396000f3006060604052600436106100cc5763ffffffff60e060020a60003504166306fdde0381146100d1578063095ea7b31461015b57806318160ddd1461019157806323b872dd146101b6578063313ce567146101de5780633eaaf86b1461020757806370a082311461021a57806379ba5097146102395780638da5cb5b1461024e57806395d89b411461027d578063a9059cbb14610290578063cae9ca51146102b2578063d4ee1d9014610317578063dc39d06d1461032a578063dd62ed3e1461034c578063f2fde38b14610371575b600080fd5b34156100dc57600080fd5b6100e4610390565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610120578082015183820152602001610108565b50505050905090810190601f16801561014d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016657600080fd5b61017d600160a060020a036004351660243561042e565b604051901515815260200160405180910390f35b341561019c57600080fd5b6101a461049b565b60405190815260200160405180910390f35b34156101c157600080fd5b61017d600160a060020a03600435811690602435166044356104cd565b34156101e957600080fd5b6101f16105e0565b60405160ff909116815260200160405180910390f35b341561021257600080fd5b6101a46105e9565b341561022557600080fd5b6101a4600160a060020a03600435166105ef565b341561024457600080fd5b61024c61060a565b005b341561025957600080fd5b610261610698565b604051600160a060020a03909116815260200160405180910390f35b341561028857600080fd5b6100e46106a7565b341561029b57600080fd5b61017d600160a060020a0360043516602435610712565b34156102bd57600080fd5b61017d60048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506107d195505050505050565b341561032257600080fd5b610261610938565b341561033557600080fd5b61017d600160a060020a0360043516602435610947565b341561035757600080fd5b6101a4600160a060020a03600435811690602435166109ea565b341561037c57600080fd5b61024c600160a060020a0360043516610a15565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104265780601f106103fb57610100808354040283529160200191610426565b820191906000526020600020905b81548152906001019060200180831161040957829003601f168201915b505050505081565b600160a060020a03338116600081815260076020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b6000805260066020527f54cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f8546005540390565b600160a060020a0383166000908152600660205260408120546104f6908363ffffffff610a5f16565b600160a060020a0380861660009081526006602090815260408083209490945560078152838220339093168252919091522054610539908363ffffffff610a5f16565b600160a060020a038086166000908152600760209081526040808320338516845282528083209490945591861681526006909152205461057f908363ffffffff610a7416565b600160a060020a03808516600081815260066020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60045460ff1681565b60055481565b600160a060020a031660009081526006602052604090205490565b60015433600160a060020a0390811691161461062557600080fd5b600154600054600160a060020a0391821691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104265780601f106103fb57610100808354040283529160200191610426565b600160a060020a03331660009081526006602052604081205461073b908363ffffffff610a5f16565b600160a060020a033381166000908152600660205260408082209390935590851681522054610770908363ffffffff610a7416565b600160a060020a0380851660008181526006602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a03338116600081815260076020908152604080832094881680845294909152808220869055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259086905190815260200160405180910390a383600160a060020a0316638f4ffcb1338530866040518563ffffffff1660e060020a0281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156108cc5780820151838201526020016108b4565b50505050905090810190601f1680156108f95780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b151561091a57600080fd5b6102c65a03f1151561092b57600080fd5b5060019695505050505050565b600154600160a060020a031681565b6000805433600160a060020a0390811691161461096357600080fd5b60008054600160a060020a038086169263a9059cbb929091169085906040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156109c957600080fd5b6102c65a03f115156109da57600080fd5b5050506040518051949350505050565b600160a060020a03918216600090815260076020908152604080832093909416825291909152205490565b60005433600160a060020a03908116911614610a3057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610a6e57600080fd5b50900390565b8181018281101561049557600080fd00a165627a7a723058201f4733f27f7e8c4eec5781823f75ee9437e6d25c5f041f8e5db13bf4d5acfa9c0029`

// DeployFixedSupplyToken deploys a new Ethereum contract, binding an instance of FixedSupplyToken to it.
func DeployFixedSupplyToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FixedSupplyToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FixedSupplyTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FixedSupplyTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FixedSupplyToken{FixedSupplyTokenCaller: FixedSupplyTokenCaller{contract: contract}, FixedSupplyTokenTransactor: FixedSupplyTokenTransactor{contract: contract}, FixedSupplyTokenFilterer: FixedSupplyTokenFilterer{contract: contract}}, nil
}

// FixedSupplyToken is an auto generated Go binding around an Ethereum contract.
type FixedSupplyToken struct {
	FixedSupplyTokenCaller     // Read-only binding to the contract
	FixedSupplyTokenTransactor // Write-only binding to the contract
	FixedSupplyTokenFilterer   // Log filterer for contract events
}

// FixedSupplyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FixedSupplyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedSupplyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FixedSupplyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedSupplyTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FixedSupplyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedSupplyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FixedSupplyTokenSession struct {
	Contract     *FixedSupplyToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FixedSupplyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FixedSupplyTokenCallerSession struct {
	Contract *FixedSupplyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FixedSupplyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FixedSupplyTokenTransactorSession struct {
	Contract     *FixedSupplyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FixedSupplyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FixedSupplyTokenRaw struct {
	Contract *FixedSupplyToken // Generic contract binding to access the raw methods on
}

// FixedSupplyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FixedSupplyTokenCallerRaw struct {
	Contract *FixedSupplyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FixedSupplyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FixedSupplyTokenTransactorRaw struct {
	Contract *FixedSupplyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFixedSupplyToken creates a new instance of FixedSupplyToken, bound to a specific deployed contract.
func NewFixedSupplyToken(address common.Address, backend bind.ContractBackend) (*FixedSupplyToken, error) {
	contract, err := bindFixedSupplyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyToken{FixedSupplyTokenCaller: FixedSupplyTokenCaller{contract: contract}, FixedSupplyTokenTransactor: FixedSupplyTokenTransactor{contract: contract}, FixedSupplyTokenFilterer: FixedSupplyTokenFilterer{contract: contract}}, nil
}

// NewFixedSupplyTokenCaller creates a new read-only instance of FixedSupplyToken, bound to a specific deployed contract.
func NewFixedSupplyTokenCaller(address common.Address, caller bind.ContractCaller) (*FixedSupplyTokenCaller, error) {
	contract, err := bindFixedSupplyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyTokenCaller{contract: contract}, nil
}

// NewFixedSupplyTokenTransactor creates a new write-only instance of FixedSupplyToken, bound to a specific deployed contract.
func NewFixedSupplyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FixedSupplyTokenTransactor, error) {
	contract, err := bindFixedSupplyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyTokenTransactor{contract: contract}, nil
}

// NewFixedSupplyTokenFilterer creates a new log filterer instance of FixedSupplyToken, bound to a specific deployed contract.
func NewFixedSupplyTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FixedSupplyTokenFilterer, error) {
	contract, err := bindFixedSupplyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyTokenFilterer{contract: contract}, nil
}

// bindFixedSupplyToken binds a generic wrapper to an already deployed contract.
func bindFixedSupplyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FixedSupplyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedSupplyToken *FixedSupplyTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FixedSupplyToken.Contract.FixedSupplyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedSupplyToken *FixedSupplyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.FixedSupplyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedSupplyToken *FixedSupplyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.FixedSupplyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedSupplyToken *FixedSupplyTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FixedSupplyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedSupplyToken *FixedSupplyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedSupplyToken *FixedSupplyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.contract.Transact(opts, method, params...)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_FixedSupplyToken *FixedSupplyTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FixedSupplyToken.contract.Call(opts, out, "_totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_FixedSupplyToken *FixedSupplyTokenSession) TotalSupply() (*big.Int, error) {
	return _FixedSupplyToken.Contract.TotalSupply(&_FixedSupplyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_FixedSupplyToken *FixedSupplyTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FixedSupplyToken.Contract.TotalSupply(&_FixedSupplyToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_FixedSupplyToken *FixedSupplyTokenCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FixedSupplyToken.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_FixedSupplyToken *FixedSupplyTokenSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _FixedSupplyToken.Contract.Allowance(&_FixedSupplyToken.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_FixedSupplyToken *FixedSupplyTokenCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _FixedSupplyToken.Contract.Allowance(&_FixedSupplyToken.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_FixedSupplyToken *FixedSupplyTokenCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FixedSupplyToken.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_FixedSupplyToken *FixedSupplyTokenSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _FixedSupplyToken.Contract.BalanceOf(&_FixedSupplyToken.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_FixedSupplyToken *FixedSupplyTokenCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _FixedSupplyToken.Contract.BalanceOf(&_FixedSupplyToken.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FixedSupplyToken *FixedSupplyTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _FixedSupplyToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FixedSupplyToken *FixedSupplyTokenSession) Decimals() (uint8, error) {
	return _FixedSupplyToken.Contract.Decimals(&_FixedSupplyToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FixedSupplyToken *FixedSupplyTokenCallerSession) Decimals() (uint8, error) {
	return _FixedSupplyToken.Contract.Decimals(&_FixedSupplyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FixedSupplyToken *FixedSupplyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FixedSupplyToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FixedSupplyToken *FixedSupplyTokenSession) Name() (string, error) {
	return _FixedSupplyToken.Contract.Name(&_FixedSupplyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FixedSupplyToken *FixedSupplyTokenCallerSession) Name() (string, error) {
	return _FixedSupplyToken.Contract.Name(&_FixedSupplyToken.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_FixedSupplyToken *FixedSupplyTokenCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FixedSupplyToken.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_FixedSupplyToken *FixedSupplyTokenSession) NewOwner() (common.Address, error) {
	return _FixedSupplyToken.Contract.NewOwner(&_FixedSupplyToken.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_FixedSupplyToken *FixedSupplyTokenCallerSession) NewOwner() (common.Address, error) {
	return _FixedSupplyToken.Contract.NewOwner(&_FixedSupplyToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FixedSupplyToken *FixedSupplyTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FixedSupplyToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FixedSupplyToken *FixedSupplyTokenSession) Owner() (common.Address, error) {
	return _FixedSupplyToken.Contract.Owner(&_FixedSupplyToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FixedSupplyToken *FixedSupplyTokenCallerSession) Owner() (common.Address, error) {
	return _FixedSupplyToken.Contract.Owner(&_FixedSupplyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FixedSupplyToken *FixedSupplyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FixedSupplyToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FixedSupplyToken *FixedSupplyTokenSession) Symbol() (string, error) {
	return _FixedSupplyToken.Contract.Symbol(&_FixedSupplyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FixedSupplyToken *FixedSupplyTokenCallerSession) Symbol() (string, error) {
	return _FixedSupplyToken.Contract.Symbol(&_FixedSupplyToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
// func (_FixedSupplyToken *FixedSupplyTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
// 	var (
// 		ret0 = new(*big.Int)
// 	)
// 	out := ret0
// 	err := _FixedSupplyToken.contract.Call(opts, out, "totalSupply")
// 	return *ret0, err
// }

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
// func (_FixedSupplyToken *FixedSupplyTokenSession) TotalSupply() (*big.Int, error) {
// 	return _FixedSupplyToken.Contract.TotalSupply(&_FixedSupplyToken.CallOpts)
// }

// // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
// //
// // Solidity: function totalSupply() constant returns(uint256)
// func (_FixedSupplyToken *FixedSupplyTokenCallerSession) TotalSupply() (*big.Int, error) {
// 	return _FixedSupplyToken.Contract.TotalSupply(&_FixedSupplyToken.CallOpts)
// }

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FixedSupplyToken *FixedSupplyTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedSupplyToken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FixedSupplyToken *FixedSupplyTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.AcceptOwnership(&_FixedSupplyToken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FixedSupplyToken *FixedSupplyTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.AcceptOwnership(&_FixedSupplyToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.Approve(&_FixedSupplyToken.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.Approve(&_FixedSupplyToken.TransactOpts, spender, tokens)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _FixedSupplyToken.contract.Transact(opts, "approveAndCall", spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.ApproveAndCall(&_FixedSupplyToken.TransactOpts, spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(spender address, tokens uint256, data bytes) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactorSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.ApproveAndCall(&_FixedSupplyToken.TransactOpts, spender, tokens, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.Transfer(&_FixedSupplyToken.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.Transfer(&_FixedSupplyToken.TransactOpts, to, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactor) TransferAnyERC20Token(opts *bind.TransactOpts, tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.contract.Transact(opts, "transferAnyERC20Token", tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.TransferAnyERC20Token(&_FixedSupplyToken.TransactOpts, tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(tokenAddress address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactorSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.TransferAnyERC20Token(&_FixedSupplyToken.TransactOpts, tokenAddress, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.TransferFrom(&_FixedSupplyToken.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_FixedSupplyToken *FixedSupplyTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.TransferFrom(&_FixedSupplyToken.TransactOpts, from, to, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_FixedSupplyToken *FixedSupplyTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _FixedSupplyToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_FixedSupplyToken *FixedSupplyTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.TransferOwnership(&_FixedSupplyToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_FixedSupplyToken *FixedSupplyTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _FixedSupplyToken.Contract.TransferOwnership(&_FixedSupplyToken.TransactOpts, _newOwner)
}

// FixedSupplyTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FixedSupplyToken contract.
type FixedSupplyTokenApprovalIterator struct {
	Event *FixedSupplyTokenApproval // Event containing the contract specifics and raw log

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
func (it *FixedSupplyTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedSupplyTokenApproval)
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
		it.Event = new(FixedSupplyTokenApproval)
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
func (it *FixedSupplyTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedSupplyTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedSupplyTokenApproval represents a Approval event raised by the FixedSupplyToken contract.
type FixedSupplyTokenApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_FixedSupplyToken *FixedSupplyTokenFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*FixedSupplyTokenApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FixedSupplyToken.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyTokenApprovalIterator{contract: _FixedSupplyToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(tokenOwner indexed address, spender indexed address, tokens uint256)
func (_FixedSupplyToken *FixedSupplyTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FixedSupplyTokenApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FixedSupplyToken.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedSupplyTokenApproval)
				if err := _FixedSupplyToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// FixedSupplyTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FixedSupplyToken contract.
type FixedSupplyTokenOwnershipTransferredIterator struct {
	Event *FixedSupplyTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FixedSupplyTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedSupplyTokenOwnershipTransferred)
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
		it.Event = new(FixedSupplyTokenOwnershipTransferred)
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
func (it *FixedSupplyTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedSupplyTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedSupplyTokenOwnershipTransferred represents a OwnershipTransferred event raised by the FixedSupplyToken contract.
type FixedSupplyTokenOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(_from indexed address, _to indexed address)
func (_FixedSupplyToken *FixedSupplyTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*FixedSupplyTokenOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _FixedSupplyToken.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyTokenOwnershipTransferredIterator{contract: _FixedSupplyToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(_from indexed address, _to indexed address)
func (_FixedSupplyToken *FixedSupplyTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FixedSupplyTokenOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _FixedSupplyToken.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedSupplyTokenOwnershipTransferred)
				if err := _FixedSupplyToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// FixedSupplyTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FixedSupplyToken contract.
type FixedSupplyTokenTransferIterator struct {
	Event *FixedSupplyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *FixedSupplyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedSupplyTokenTransfer)
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
		it.Event = new(FixedSupplyTokenTransfer)
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
func (it *FixedSupplyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedSupplyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedSupplyTokenTransfer represents a Transfer event raised by the FixedSupplyToken contract.
type FixedSupplyTokenTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, tokens uint256)
func (_FixedSupplyToken *FixedSupplyTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FixedSupplyTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FixedSupplyToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FixedSupplyTokenTransferIterator{contract: _FixedSupplyToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, tokens uint256)
func (_FixedSupplyToken *FixedSupplyTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FixedSupplyTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FixedSupplyToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedSupplyTokenTransfer)
				if err := _FixedSupplyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101fe8061003b6000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166379ba509781146100665780638da5cb5b1461007b578063d4ee1d90146100aa578063f2fde38b146100bd575b600080fd5b341561007157600080fd5b6100796100dc565b005b341561008657600080fd5b61008e61016a565b604051600160a060020a03909116815260200160405180910390f35b34156100b557600080fd5b61008e610179565b34156100c857600080fd5b610079600160a060020a0360043516610188565b60015433600160a060020a039081169116146100f757600080fd5b600154600054600160a060020a0391821691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600154600160a060020a031681565b60005433600160a060020a039081169116146101a357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820572477ccdb3d3010e3ebde3ec482c53c9bf4c34938afda56daf4bce02931a26f0029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Owned *OwnedCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Owned *OwnedSession) NewOwner() (common.Address, error) {
	return _Owned.Contract.NewOwner(&_Owned.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() constant returns(address)
func (_Owned *OwnedCallerSession) NewOwner() (common.Address, error) {
	return _Owned.Contract.NewOwner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Owned *OwnedTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Owned *OwnedSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Owned *OwnedTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, _newOwner)
}

// OwnedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Owned contract.
type OwnedOwnershipTransferredIterator struct {
	Event *OwnedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnedOwnershipTransferred)
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
		it.Event = new(OwnedOwnershipTransferred)
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
func (it *OwnedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnedOwnershipTransferred represents a OwnershipTransferred event raised by the Owned contract.
type OwnedOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(_from indexed address, _to indexed address)
func (_Owned *OwnedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*OwnedOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Owned.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &OwnedOwnershipTransferredIterator{contract: _Owned.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(_from indexed address, _to indexed address)
func (_Owned *OwnedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnedOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Owned.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnedOwnershipTransferred)
				if err := _Owned.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a72305820c7bed5292b3bb93dc8e349ba905add2ab7c8a0b540a83ba974168f76c28f41920029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
