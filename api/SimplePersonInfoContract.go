// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"personIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAge\",\"type\":\"uint256\"}],\"name\":\"PersonInfoUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_personIndex\",\"type\":\"uint256\"}],\"name\":\"getPersonInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPersonsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"persons\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"setPersonInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506106bb8061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806333f3b2a41461004e5780638f97cff014610063578063a2f9eac614610078578063d336ac8014610099575b5f80fd5b61006161005c3660046103f0565b6100ac565b005b5f546040519081526020015b60405180910390f35b61008b61008636600461049f565b6101fe565b60405161006f9291906104b6565b61008b6100a736600461049f565b6102b2565b5f8251116101015760405162461bcd60e51b815260206004820152601860248201527f4e616d652073686f756c64206e6f7420626520656d707479000000000000000060448201526064015b60405180910390fd5b5f81116101505760405162461bcd60e51b815260206004820152601c60248201527f4167652073686f756c642062652067726561746572207468616e20300000000060448201526064016100f8565b60408051808201909152828152602081018290525f8054600181018255908052815182916002027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019081906101a6908261058c565b50602091909101516001918201555f546101c0919061064c565b7f96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec3284846040516101f19291906104b6565b60405180910390a2505050565b5f818154811061020c575f80fd5b905f5260205f2090600202015f91509050805f01805461022b90610508565b80601f016020809104026020016040519081016040528092919081815260200182805461025790610508565b80156102a25780601f10610279576101008083540402835291602001916102a2565b820191905f5260205f20905b81548152906001019060200180831161028557829003601f168201915b5050505050908060010154905082565b5f80546060919083106103075760405162461bcd60e51b815260206004820152601a60248201527f506572736f6e20696e646578206f7574206f6620626f756e647300000000000060448201526064016100f8565b5f80848154811061031a5761031a610671565b905f5260205f2090600202016040518060400160405290815f8201805461034090610508565b80601f016020809104026020016040519081016040528092919081815260200182805461036c90610508565b80156103b75780601f1061038e576101008083540402835291602001916103b7565b820191905f5260205f20905b81548152906001019060200180831161039a57829003601f168201915b5050509183525050600191909101546020918201528151910151909590945092505050565b634e487b7160e01b5f52604160045260245ffd5b5f8060408385031215610401575f80fd5b823567ffffffffffffffff80821115610418575f80fd5b818501915085601f83011261042b575f80fd5b81358181111561043d5761043d6103dc565b604051601f8201601f19908116603f01168101908382118183101715610465576104656103dc565b8160405282815288602084870101111561047d575f80fd5b826020860160208301375f602093820184015298969091013596505050505050565b5f602082840312156104af575f80fd5b5035919050565b604081525f83518060408401525f5b818110156104e257602081870181015160608684010152016104c5565b505f606082850101526060601f19601f8301168401019150508260208301529392505050565b600181811c9082168061051c57607f821691505b60208210810361053a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561058757805f5260205f20601f840160051c810160208510156105655750805b601f840160051c820191505b81811015610584575f8155600101610571565b50505b505050565b815167ffffffffffffffff8111156105a6576105a66103dc565b6105ba816105b48454610508565b84610540565b602080601f8311600181146105ed575f84156105d65750858301515b5f19600386901b1c1916600185901b178555610644565b5f85815260208120601f198616915b8281101561061b578886015182559484019460019091019084016105fc565b508582101561063857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b8181038181111561066b57634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52603260045260245ffdfea2646970667358221220819f9f2251bc4c7ccd7ed04c3102e7a387c09d9aa0a77cbc418602b8dcb98adc64736f6c63430008180033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_Api *ApiCaller) GetPersonInfo(opts *bind.CallOpts, _personIndex *big.Int) (string, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPersonInfo", _personIndex)

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_Api *ApiSession) GetPersonInfo(_personIndex *big.Int) (string, *big.Int, error) {
	return _Api.Contract.GetPersonInfo(&_Api.CallOpts, _personIndex)
}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_Api *ApiCallerSession) GetPersonInfo(_personIndex *big.Int) (string, *big.Int, error) {
	return _Api.Contract.GetPersonInfo(&_Api.CallOpts, _personIndex)
}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_Api *ApiCaller) GetPersonsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPersonsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_Api *ApiSession) GetPersonsCount() (*big.Int, error) {
	return _Api.Contract.GetPersonsCount(&_Api.CallOpts)
}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_Api *ApiCallerSession) GetPersonsCount() (*big.Int, error) {
	return _Api.Contract.GetPersonsCount(&_Api.CallOpts)
}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_Api *ApiCaller) Persons(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "persons", arg0)

	outstruct := new(struct {
		Name string
		Age  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Age = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_Api *ApiSession) Persons(arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	return _Api.Contract.Persons(&_Api.CallOpts, arg0)
}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_Api *ApiCallerSession) Persons(arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	return _Api.Contract.Persons(&_Api.CallOpts, arg0)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_Api *ApiTransactor) SetPersonInfo(opts *bind.TransactOpts, _name string, _age *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPersonInfo", _name, _age)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_Api *ApiSession) SetPersonInfo(_name string, _age *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetPersonInfo(&_Api.TransactOpts, _name, _age)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_Api *ApiTransactorSession) SetPersonInfo(_name string, _age *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetPersonInfo(&_Api.TransactOpts, _name, _age)
}

// ApiPersonInfoUpdatedIterator is returned from FilterPersonInfoUpdated and is used to iterate over the raw logs and unpacked data for PersonInfoUpdated events raised by the Api contract.
type ApiPersonInfoUpdatedIterator struct {
	Event *ApiPersonInfoUpdated // Event containing the contract specifics and raw log

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
func (it *ApiPersonInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiPersonInfoUpdated)
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
		it.Event = new(ApiPersonInfoUpdated)
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
func (it *ApiPersonInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiPersonInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiPersonInfoUpdated represents a PersonInfoUpdated event raised by the Api contract.
type ApiPersonInfoUpdated struct {
	PersonIndex *big.Int
	NewName     string
	NewAge      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPersonInfoUpdated is a free log retrieval operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_Api *ApiFilterer) FilterPersonInfoUpdated(opts *bind.FilterOpts, personIndex []*big.Int) (*ApiPersonInfoUpdatedIterator, error) {

	var personIndexRule []interface{}
	for _, personIndexItem := range personIndex {
		personIndexRule = append(personIndexRule, personIndexItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "PersonInfoUpdated", personIndexRule)
	if err != nil {
		return nil, err
	}
	return &ApiPersonInfoUpdatedIterator{contract: _Api.contract, event: "PersonInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPersonInfoUpdated is a free log subscription operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_Api *ApiFilterer) WatchPersonInfoUpdated(opts *bind.WatchOpts, sink chan<- *ApiPersonInfoUpdated, personIndex []*big.Int) (event.Subscription, error) {

	var personIndexRule []interface{}
	for _, personIndexItem := range personIndex {
		personIndexRule = append(personIndexRule, personIndexItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "PersonInfoUpdated", personIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiPersonInfoUpdated)
				if err := _Api.contract.UnpackLog(event, "PersonInfoUpdated", log); err != nil {
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

// ParsePersonInfoUpdated is a log parse operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_Api *ApiFilterer) ParsePersonInfoUpdated(log types.Log) (*ApiPersonInfoUpdated, error) {
	event := new(ApiPersonInfoUpdated)
	if err := _Api.contract.UnpackLog(event, "PersonInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
