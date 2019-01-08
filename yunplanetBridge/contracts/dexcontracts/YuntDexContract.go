// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dexcontracts

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

// DexcontractsABI is the input ABI used to generate the binding from.
const DexcontractsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_yuntToken\",\"type\":\"address\"}],\"name\":\"config\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"exChangeYUNT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"Receive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"mylog\",\"type\":\"event\"}]"

// Dexcontracts is an auto generated Go binding around an Ethereum contract.
type Dexcontracts struct {
	DexcontractsCaller     // Read-only binding to the contract
	DexcontractsTransactor // Write-only binding to the contract
	DexcontractsFilterer   // Log filterer for contract events
}

// DexcontractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexcontractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexcontractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexcontractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexcontractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexcontractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexcontractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexcontractsSession struct {
	Contract     *Dexcontracts     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexcontractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexcontractsCallerSession struct {
	Contract *DexcontractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DexcontractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexcontractsTransactorSession struct {
	Contract     *DexcontractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DexcontractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexcontractsRaw struct {
	Contract *Dexcontracts // Generic contract binding to access the raw methods on
}

// DexcontractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexcontractsCallerRaw struct {
	Contract *DexcontractsCaller // Generic read-only contract binding to access the raw methods on
}

// DexcontractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexcontractsTransactorRaw struct {
	Contract *DexcontractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDexcontracts creates a new instance of Dexcontracts, bound to a specific deployed contract.
func NewDexcontracts(address common.Address, backend bind.ContractBackend) (*Dexcontracts, error) {
	contract, err := bindDexcontracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dexcontracts{DexcontractsCaller: DexcontractsCaller{contract: contract}, DexcontractsTransactor: DexcontractsTransactor{contract: contract}, DexcontractsFilterer: DexcontractsFilterer{contract: contract}}, nil
}

// NewDexcontractsCaller creates a new read-only instance of Dexcontracts, bound to a specific deployed contract.
func NewDexcontractsCaller(address common.Address, caller bind.ContractCaller) (*DexcontractsCaller, error) {
	contract, err := bindDexcontracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexcontractsCaller{contract: contract}, nil
}

// NewDexcontractsTransactor creates a new write-only instance of Dexcontracts, bound to a specific deployed contract.
func NewDexcontractsTransactor(address common.Address, transactor bind.ContractTransactor) (*DexcontractsTransactor, error) {
	contract, err := bindDexcontracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexcontractsTransactor{contract: contract}, nil
}

// NewDexcontractsFilterer creates a new log filterer instance of Dexcontracts, bound to a specific deployed contract.
func NewDexcontractsFilterer(address common.Address, filterer bind.ContractFilterer) (*DexcontractsFilterer, error) {
	contract, err := bindDexcontracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexcontractsFilterer{contract: contract}, nil
}

// bindDexcontracts binds a generic wrapper to an already deployed contract.
func bindDexcontracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DexcontractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dexcontracts *DexcontractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dexcontracts.Contract.DexcontractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dexcontracts *DexcontractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dexcontracts.Contract.DexcontractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dexcontracts *DexcontractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dexcontracts.Contract.DexcontractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dexcontracts *DexcontractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dexcontracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dexcontracts *DexcontractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dexcontracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dexcontracts *DexcontractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dexcontracts.Contract.contract.Transact(opts, method, params...)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_Dexcontracts *DexcontractsTransactor) Callback(opts *bind.TransactOpts, myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _Dexcontracts.contract.Transact(opts, "__callback", myid, result, proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_Dexcontracts *DexcontractsSession) Callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _Dexcontracts.Contract.Callback(&_Dexcontracts.TransactOpts, myid, result, proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_Dexcontracts *DexcontractsTransactorSession) Callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _Dexcontracts.Contract.Callback(&_Dexcontracts.TransactOpts, myid, result, proof)
}

// Config is a paid mutator transaction binding the contract method 0x0e68ec95.
//
// Solidity: function config(_yuntToken address) returns()
func (_Dexcontracts *DexcontractsTransactor) Config(opts *bind.TransactOpts, _yuntToken common.Address) (*types.Transaction, error) {
	return _Dexcontracts.contract.Transact(opts, "config", _yuntToken)
}

// Config is a paid mutator transaction binding the contract method 0x0e68ec95.
//
// Solidity: function config(_yuntToken address) returns()
func (_Dexcontracts *DexcontractsSession) Config(_yuntToken common.Address) (*types.Transaction, error) {
	return _Dexcontracts.Contract.Config(&_Dexcontracts.TransactOpts, _yuntToken)
}

// Config is a paid mutator transaction binding the contract method 0x0e68ec95.
//
// Solidity: function config(_yuntToken address) returns()
func (_Dexcontracts *DexcontractsTransactorSession) Config(_yuntToken common.Address) (*types.Transaction, error) {
	return _Dexcontracts.Contract.Config(&_Dexcontracts.TransactOpts, _yuntToken)
}

// ExChangeYUNT is a paid mutator transaction binding the contract method 0x788be027.
//
// Solidity: function exChangeYUNT(_from address, _to address, _value uint256) returns()
func (_Dexcontracts *DexcontractsTransactor) ExChangeYUNT(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dexcontracts.contract.Transact(opts, "exChangeYUNT", _from, _to, _value)
}

// ExChangeYUNT is a paid mutator transaction binding the contract method 0x788be027.
//
// Solidity: function exChangeYUNT(_from address, _to address, _value uint256) returns()
func (_Dexcontracts *DexcontractsSession) ExChangeYUNT(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dexcontracts.Contract.ExChangeYUNT(&_Dexcontracts.TransactOpts, _from, _to, _value)
}

// ExChangeYUNT is a paid mutator transaction binding the contract method 0x788be027.
//
// Solidity: function exChangeYUNT(_from address, _to address, _value uint256) returns()
func (_Dexcontracts *DexcontractsTransactorSession) ExChangeYUNT(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dexcontracts.Contract.ExChangeYUNT(&_Dexcontracts.TransactOpts, _from, _to, _value)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_Dexcontracts *DexcontractsTransactor) ReceiveApproval(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Dexcontracts.contract.Transact(opts, "receiveApproval", _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_Dexcontracts *DexcontractsSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Dexcontracts.Contract.ReceiveApproval(&_Dexcontracts.TransactOpts, _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_Dexcontracts *DexcontractsTransactorSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Dexcontracts.Contract.ReceiveApproval(&_Dexcontracts.TransactOpts, _from, _value, _token, _extraData)
}

// DexcontractsReceiveIterator is returned from FilterReceive and is used to iterate over the raw logs and unpacked data for Receive events raised by the Dexcontracts contract.
type DexcontractsReceiveIterator struct {
	Event *DexcontractsReceive // Event containing the contract specifics and raw log

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
func (it *DexcontractsReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexcontractsReceive)
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
		it.Event = new(DexcontractsReceive)
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
func (it *DexcontractsReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexcontractsReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexcontractsReceive represents a Receive event raised by the Dexcontracts contract.
type DexcontractsReceive struct {
	From      common.Address
	Value     *big.Int
	Token     common.Address
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceive is a free log retrieval operation binding the contract event 0x676ee85d878585bb8ad2099ba98bda55b604b8fd3cfed77c72a035865d30a93e.
//
// Solidity: e Receive(from address, value uint256, token address, extraData bytes)
func (_Dexcontracts *DexcontractsFilterer) FilterReceive(opts *bind.FilterOpts) (*DexcontractsReceiveIterator, error) {

	logs, sub, err := _Dexcontracts.contract.FilterLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return &DexcontractsReceiveIterator{contract: _Dexcontracts.contract, event: "Receive", logs: logs, sub: sub}, nil
}

// WatchReceive is a free log subscription operation binding the contract event 0x676ee85d878585bb8ad2099ba98bda55b604b8fd3cfed77c72a035865d30a93e.
//
// Solidity: e Receive(from address, value uint256, token address, extraData bytes)
func (_Dexcontracts *DexcontractsFilterer) WatchReceive(opts *bind.WatchOpts, sink chan<- *DexcontractsReceive) (event.Subscription, error) {

	logs, sub, err := _Dexcontracts.contract.WatchLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexcontractsReceive)
				if err := _Dexcontracts.contract.UnpackLog(event, "Receive", log); err != nil {
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

// DexcontractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dexcontracts contract.
type DexcontractsTransferIterator struct {
	Event *DexcontractsTransfer // Event containing the contract specifics and raw log

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
func (it *DexcontractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexcontractsTransfer)
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
		it.Event = new(DexcontractsTransfer)
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
func (it *DexcontractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexcontractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexcontractsTransfer represents a Transfer event raised by the Dexcontracts contract.
type DexcontractsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Dexcontracts *DexcontractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DexcontractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dexcontracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DexcontractsTransferIterator{contract: _Dexcontracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Dexcontracts *DexcontractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DexcontractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dexcontracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexcontractsTransfer)
				if err := _Dexcontracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// DexcontractsMylogIterator is returned from FilterMylog and is used to iterate over the raw logs and unpacked data for Mylog events raised by the Dexcontracts contract.
type DexcontractsMylogIterator struct {
	Event *DexcontractsMylog // Event containing the contract specifics and raw log

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
func (it *DexcontractsMylogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexcontractsMylog)
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
		it.Event = new(DexcontractsMylog)
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
func (it *DexcontractsMylogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexcontractsMylogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexcontractsMylog represents a Mylog event raised by the Dexcontracts contract.
type DexcontractsMylog struct {
	Code *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMylog is a free log retrieval operation binding the contract event 0x67631d1bdda6fede513a18dab2d7282eb1233c8a01cd01f9580aa9543dd82b32.
//
// Solidity: e mylog(code uint256)
func (_Dexcontracts *DexcontractsFilterer) FilterMylog(opts *bind.FilterOpts) (*DexcontractsMylogIterator, error) {

	logs, sub, err := _Dexcontracts.contract.FilterLogs(opts, "mylog")
	if err != nil {
		return nil, err
	}
	return &DexcontractsMylogIterator{contract: _Dexcontracts.contract, event: "mylog", logs: logs, sub: sub}, nil
}

// WatchMylog is a free log subscription operation binding the contract event 0x67631d1bdda6fede513a18dab2d7282eb1233c8a01cd01f9580aa9543dd82b32.
//
// Solidity: e mylog(code uint256)
func (_Dexcontracts *DexcontractsFilterer) WatchMylog(opts *bind.WatchOpts, sink chan<- *DexcontractsMylog) (event.Subscription, error) {

	logs, sub, err := _Dexcontracts.contract.WatchLogs(opts, "mylog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexcontractsMylog)
				if err := _Dexcontracts.contract.UnpackLog(event, "mylog", log); err != nil {
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
