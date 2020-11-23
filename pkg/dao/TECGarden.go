// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dao

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TECGardenABI is the input ABI used to generate the binding from.
const TECGardenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"MARKET_MAKER_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONVICTION_VOTING_APP_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PRESALE_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REDEMPTIONS_APP_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOLLGATE_APP_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MARKETPLACE_CONTROLLER_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DANDELION_VOTING_APP_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"HOOKED_TOKEN_MANAGER_APP_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_FORMULA_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_daoFactory\",\"type\":\"address\"},{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_miniMeFactory\",\"type\":\"address\"},{\"name\":\"_aragonID\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"dao\",\"type\":\"address\"}],\"name\":\"DeployDao\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"dao\",\"type\":\"address\"}],\"name\":\"SetupDao\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"DeployToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"appProxy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"appId\",\"type\":\"bytes32\"}],\"name\":\"InstalledApp\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteTokenName\",\"type\":\"string\"},{\"name\":\"_voteTokenSymbol\",\"type\":\"string\"},{\"name\":\"_votingSettings\",\"type\":\"uint64[5]\"},{\"name\":\"_useAgentAsVault\",\"type\":\"bool\"},{\"name\":\"_permissionManager\",\"type\":\"address\"}],\"name\":\"createDaoTxOne\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tollgateFeeToken\",\"type\":\"address\"},{\"name\":\"_tollgateFeeAmount\",\"type\":\"uint256\"},{\"name\":\"_redeemableTokens\",\"type\":\"address[]\"},{\"name\":\"_convictionSettings\",\"type\":\"uint256[4]\"},{\"name\":\"_collateralToken\",\"type\":\"address\"}],\"name\":\"createDaoTxTwo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_holders\",\"type\":\"address[]\"},{\"name\":\"_stakes\",\"type\":\"uint256[]\"},{\"name\":\"_openDate\",\"type\":\"uint64\"},{\"name\":\"_vestingCliffPeriod\",\"type\":\"uint64\"},{\"name\":\"_vestingCompletePeriod\",\"type\":\"uint64\"}],\"name\":\"createTxTokenHolders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_goal\",\"type\":\"uint256\"},{\"name\":\"_period\",\"type\":\"uint64\"},{\"name\":\"_exchangeRate\",\"type\":\"uint256\"},{\"name\":\"_vestingCliffPeriod\",\"type\":\"uint64\"},{\"name\":\"_vestingCompletePeriod\",\"type\":\"uint64\"},{\"name\":\"_supplyOfferedPct\",\"type\":\"uint256\"},{\"name\":\"_fundingForBeneficiaryPct\",\"type\":\"uint256\"},{\"name\":\"_openDate\",\"type\":\"uint64\"},{\"name\":\"_buyFeePct\",\"type\":\"uint256\"},{\"name\":\"_sellFeePct\",\"type\":\"uint256\"}],\"name\":\"createDaoTxThree\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_virtualSupply\",\"type\":\"uint256\"},{\"name\":\"_virtualBalance\",\"type\":\"uint256\"},{\"name\":\"_reserveRatio\",\"type\":\"uint32\"}],\"name\":\"createDaoTxFour\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TECGarden is an auto generated Go binding around an Ethereum contract.
type TECGarden struct {
	TECGardenCaller     // Read-only binding to the contract
	TECGardenTransactor // Write-only binding to the contract
	TECGardenFilterer   // Log filterer for contract events
}

// TECGardenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TECGardenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TECGardenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TECGardenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TECGardenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TECGardenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TECGardenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TECGardenSession struct {
	Contract     *TECGarden        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TECGardenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TECGardenCallerSession struct {
	Contract *TECGardenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TECGardenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TECGardenTransactorSession struct {
	Contract     *TECGardenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TECGardenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TECGardenRaw struct {
	Contract *TECGarden // Generic contract binding to access the raw methods on
}

// TECGardenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TECGardenCallerRaw struct {
	Contract *TECGardenCaller // Generic read-only contract binding to access the raw methods on
}

// TECGardenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TECGardenTransactorRaw struct {
	Contract *TECGardenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTECGarden creates a new instance of TECGarden, bound to a specific deployed contract.
func NewTECGarden(address common.Address, backend bind.ContractBackend) (*TECGarden, error) {
	contract, err := bindTECGarden(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TECGarden{TECGardenCaller: TECGardenCaller{contract: contract}, TECGardenTransactor: TECGardenTransactor{contract: contract}, TECGardenFilterer: TECGardenFilterer{contract: contract}}, nil
}

// NewTECGardenCaller creates a new read-only instance of TECGarden, bound to a specific deployed contract.
func NewTECGardenCaller(address common.Address, caller bind.ContractCaller) (*TECGardenCaller, error) {
	contract, err := bindTECGarden(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TECGardenCaller{contract: contract}, nil
}

// NewTECGardenTransactor creates a new write-only instance of TECGarden, bound to a specific deployed contract.
func NewTECGardenTransactor(address common.Address, transactor bind.ContractTransactor) (*TECGardenTransactor, error) {
	contract, err := bindTECGarden(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TECGardenTransactor{contract: contract}, nil
}

// NewTECGardenFilterer creates a new log filterer instance of TECGarden, bound to a specific deployed contract.
func NewTECGardenFilterer(address common.Address, filterer bind.ContractFilterer) (*TECGardenFilterer, error) {
	contract, err := bindTECGarden(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TECGardenFilterer{contract: contract}, nil
}

// bindTECGarden binds a generic wrapper to an already deployed contract.
func bindTECGarden(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TECGardenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TECGarden *TECGardenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TECGarden.Contract.TECGardenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TECGarden *TECGardenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TECGarden.Contract.TECGardenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TECGarden *TECGardenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TECGarden.Contract.TECGardenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TECGarden *TECGardenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TECGarden.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TECGarden *TECGardenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TECGarden.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TECGarden *TECGardenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TECGarden.Contract.contract.Transact(opts, method, params...)
}

// BANCORFORMULAID is a free data retrieval call binding the contract method 0xfba33ebb.
//
// Solidity: function BANCOR_FORMULA_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) BANCORFORMULAID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "BANCOR_FORMULA_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORFORMULAID is a free data retrieval call binding the contract method 0xfba33ebb.
//
// Solidity: function BANCOR_FORMULA_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) BANCORFORMULAID() ([32]byte, error) {
	return _TECGarden.Contract.BANCORFORMULAID(&_TECGarden.CallOpts)
}

// BANCORFORMULAID is a free data retrieval call binding the contract method 0xfba33ebb.
//
// Solidity: function BANCOR_FORMULA_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) BANCORFORMULAID() ([32]byte, error) {
	return _TECGarden.Contract.BANCORFORMULAID(&_TECGarden.CallOpts)
}

// CONVICTIONVOTINGAPPID is a free data retrieval call binding the contract method 0x262337e9.
//
// Solidity: function CONVICTION_VOTING_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) CONVICTIONVOTINGAPPID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "CONVICTION_VOTING_APP_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONVICTIONVOTINGAPPID is a free data retrieval call binding the contract method 0x262337e9.
//
// Solidity: function CONVICTION_VOTING_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) CONVICTIONVOTINGAPPID() ([32]byte, error) {
	return _TECGarden.Contract.CONVICTIONVOTINGAPPID(&_TECGarden.CallOpts)
}

// CONVICTIONVOTINGAPPID is a free data retrieval call binding the contract method 0x262337e9.
//
// Solidity: function CONVICTION_VOTING_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) CONVICTIONVOTINGAPPID() ([32]byte, error) {
	return _TECGarden.Contract.CONVICTIONVOTINGAPPID(&_TECGarden.CallOpts)
}

// DANDELIONVOTINGAPPID is a free data retrieval call binding the contract method 0xb610454f.
//
// Solidity: function DANDELION_VOTING_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) DANDELIONVOTINGAPPID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "DANDELION_VOTING_APP_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DANDELIONVOTINGAPPID is a free data retrieval call binding the contract method 0xb610454f.
//
// Solidity: function DANDELION_VOTING_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) DANDELIONVOTINGAPPID() ([32]byte, error) {
	return _TECGarden.Contract.DANDELIONVOTINGAPPID(&_TECGarden.CallOpts)
}

// DANDELIONVOTINGAPPID is a free data retrieval call binding the contract method 0xb610454f.
//
// Solidity: function DANDELION_VOTING_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) DANDELIONVOTINGAPPID() ([32]byte, error) {
	return _TECGarden.Contract.DANDELIONVOTINGAPPID(&_TECGarden.CallOpts)
}

// HOOKEDTOKENMANAGERAPPID is a free data retrieval call binding the contract method 0xf21a4cf1.
//
// Solidity: function HOOKED_TOKEN_MANAGER_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) HOOKEDTOKENMANAGERAPPID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "HOOKED_TOKEN_MANAGER_APP_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HOOKEDTOKENMANAGERAPPID is a free data retrieval call binding the contract method 0xf21a4cf1.
//
// Solidity: function HOOKED_TOKEN_MANAGER_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) HOOKEDTOKENMANAGERAPPID() ([32]byte, error) {
	return _TECGarden.Contract.HOOKEDTOKENMANAGERAPPID(&_TECGarden.CallOpts)
}

// HOOKEDTOKENMANAGERAPPID is a free data retrieval call binding the contract method 0xf21a4cf1.
//
// Solidity: function HOOKED_TOKEN_MANAGER_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) HOOKEDTOKENMANAGERAPPID() ([32]byte, error) {
	return _TECGarden.Contract.HOOKEDTOKENMANAGERAPPID(&_TECGarden.CallOpts)
}

// MARKETPLACECONTROLLERID is a free data retrieval call binding the contract method 0x9b407b3a.
//
// Solidity: function MARKETPLACE_CONTROLLER_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) MARKETPLACECONTROLLERID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "MARKETPLACE_CONTROLLER_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETPLACECONTROLLERID is a free data retrieval call binding the contract method 0x9b407b3a.
//
// Solidity: function MARKETPLACE_CONTROLLER_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) MARKETPLACECONTROLLERID() ([32]byte, error) {
	return _TECGarden.Contract.MARKETPLACECONTROLLERID(&_TECGarden.CallOpts)
}

// MARKETPLACECONTROLLERID is a free data retrieval call binding the contract method 0x9b407b3a.
//
// Solidity: function MARKETPLACE_CONTROLLER_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) MARKETPLACECONTROLLERID() ([32]byte, error) {
	return _TECGarden.Contract.MARKETPLACECONTROLLERID(&_TECGarden.CallOpts)
}

// MARKETMAKERID is a free data retrieval call binding the contract method 0x25b1aa2b.
//
// Solidity: function MARKET_MAKER_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) MARKETMAKERID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "MARKET_MAKER_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETMAKERID is a free data retrieval call binding the contract method 0x25b1aa2b.
//
// Solidity: function MARKET_MAKER_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) MARKETMAKERID() ([32]byte, error) {
	return _TECGarden.Contract.MARKETMAKERID(&_TECGarden.CallOpts)
}

// MARKETMAKERID is a free data retrieval call binding the contract method 0x25b1aa2b.
//
// Solidity: function MARKET_MAKER_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) MARKETMAKERID() ([32]byte, error) {
	return _TECGarden.Contract.MARKETMAKERID(&_TECGarden.CallOpts)
}

// PRESALEID is a free data retrieval call binding the contract method 0x4baaaac3.
//
// Solidity: function PRESALE_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) PRESALEID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "PRESALE_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRESALEID is a free data retrieval call binding the contract method 0x4baaaac3.
//
// Solidity: function PRESALE_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) PRESALEID() ([32]byte, error) {
	return _TECGarden.Contract.PRESALEID(&_TECGarden.CallOpts)
}

// PRESALEID is a free data retrieval call binding the contract method 0x4baaaac3.
//
// Solidity: function PRESALE_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) PRESALEID() ([32]byte, error) {
	return _TECGarden.Contract.PRESALEID(&_TECGarden.CallOpts)
}

// REDEMPTIONSAPPID is a free data retrieval call binding the contract method 0x65932650.
//
// Solidity: function REDEMPTIONS_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) REDEMPTIONSAPPID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "REDEMPTIONS_APP_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REDEMPTIONSAPPID is a free data retrieval call binding the contract method 0x65932650.
//
// Solidity: function REDEMPTIONS_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) REDEMPTIONSAPPID() ([32]byte, error) {
	return _TECGarden.Contract.REDEMPTIONSAPPID(&_TECGarden.CallOpts)
}

// REDEMPTIONSAPPID is a free data retrieval call binding the contract method 0x65932650.
//
// Solidity: function REDEMPTIONS_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) REDEMPTIONSAPPID() ([32]byte, error) {
	return _TECGarden.Contract.REDEMPTIONSAPPID(&_TECGarden.CallOpts)
}

// TOLLGATEAPPID is a free data retrieval call binding the contract method 0x7b25aeb9.
//
// Solidity: function TOLLGATE_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCaller) TOLLGATEAPPID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TECGarden.contract.Call(opts, &out, "TOLLGATE_APP_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOLLGATEAPPID is a free data retrieval call binding the contract method 0x7b25aeb9.
//
// Solidity: function TOLLGATE_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenSession) TOLLGATEAPPID() ([32]byte, error) {
	return _TECGarden.Contract.TOLLGATEAPPID(&_TECGarden.CallOpts)
}

// TOLLGATEAPPID is a free data retrieval call binding the contract method 0x7b25aeb9.
//
// Solidity: function TOLLGATE_APP_ID() view returns(bytes32)
func (_TECGarden *TECGardenCallerSession) TOLLGATEAPPID() ([32]byte, error) {
	return _TECGarden.Contract.TOLLGATEAPPID(&_TECGarden.CallOpts)
}

// CreateDaoTxFour is a paid mutator transaction binding the contract method 0xbce1cb5e.
//
// Solidity: function createDaoTxFour(string _id, uint256 _virtualSupply, uint256 _virtualBalance, uint32 _reserveRatio) returns()
func (_TECGarden *TECGardenTransactor) CreateDaoTxFour(opts *bind.TransactOpts, _id string, _virtualSupply *big.Int, _virtualBalance *big.Int, _reserveRatio uint32) (*types.Transaction, error) {
	return _TECGarden.contract.Transact(opts, "createDaoTxFour", _id, _virtualSupply, _virtualBalance, _reserveRatio)
}

// CreateDaoTxFour is a paid mutator transaction binding the contract method 0xbce1cb5e.
//
// Solidity: function createDaoTxFour(string _id, uint256 _virtualSupply, uint256 _virtualBalance, uint32 _reserveRatio) returns()
func (_TECGarden *TECGardenSession) CreateDaoTxFour(_id string, _virtualSupply *big.Int, _virtualBalance *big.Int, _reserveRatio uint32) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateDaoTxFour(&_TECGarden.TransactOpts, _id, _virtualSupply, _virtualBalance, _reserveRatio)
}

// CreateDaoTxFour is a paid mutator transaction binding the contract method 0xbce1cb5e.
//
// Solidity: function createDaoTxFour(string _id, uint256 _virtualSupply, uint256 _virtualBalance, uint32 _reserveRatio) returns()
func (_TECGarden *TECGardenTransactorSession) CreateDaoTxFour(_id string, _virtualSupply *big.Int, _virtualBalance *big.Int, _reserveRatio uint32) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateDaoTxFour(&_TECGarden.TransactOpts, _id, _virtualSupply, _virtualBalance, _reserveRatio)
}

// CreateDaoTxOne is a paid mutator transaction binding the contract method 0x5618021c.
//
// Solidity: function createDaoTxOne(string _voteTokenName, string _voteTokenSymbol, uint64[5] _votingSettings, bool _useAgentAsVault, address _permissionManager) returns()
func (_TECGarden *TECGardenTransactor) CreateDaoTxOne(opts *bind.TransactOpts, _voteTokenName string, _voteTokenSymbol string, _votingSettings [5]uint64, _useAgentAsVault bool, _permissionManager common.Address) (*types.Transaction, error) {
	return _TECGarden.contract.Transact(opts, "createDaoTxOne", _voteTokenName, _voteTokenSymbol, _votingSettings, _useAgentAsVault, _permissionManager)
}

// CreateDaoTxOne is a paid mutator transaction binding the contract method 0x5618021c.
//
// Solidity: function createDaoTxOne(string _voteTokenName, string _voteTokenSymbol, uint64[5] _votingSettings, bool _useAgentAsVault, address _permissionManager) returns()
func (_TECGarden *TECGardenSession) CreateDaoTxOne(_voteTokenName string, _voteTokenSymbol string, _votingSettings [5]uint64, _useAgentAsVault bool, _permissionManager common.Address) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateDaoTxOne(&_TECGarden.TransactOpts, _voteTokenName, _voteTokenSymbol, _votingSettings, _useAgentAsVault, _permissionManager)
}

// CreateDaoTxOne is a paid mutator transaction binding the contract method 0x5618021c.
//
// Solidity: function createDaoTxOne(string _voteTokenName, string _voteTokenSymbol, uint64[5] _votingSettings, bool _useAgentAsVault, address _permissionManager) returns()
func (_TECGarden *TECGardenTransactorSession) CreateDaoTxOne(_voteTokenName string, _voteTokenSymbol string, _votingSettings [5]uint64, _useAgentAsVault bool, _permissionManager common.Address) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateDaoTxOne(&_TECGarden.TransactOpts, _voteTokenName, _voteTokenSymbol, _votingSettings, _useAgentAsVault, _permissionManager)
}

// CreateDaoTxThree is a paid mutator transaction binding the contract method 0xa74dd284.
//
// Solidity: function createDaoTxThree(uint256 _goal, uint64 _period, uint256 _exchangeRate, uint64 _vestingCliffPeriod, uint64 _vestingCompletePeriod, uint256 _supplyOfferedPct, uint256 _fundingForBeneficiaryPct, uint64 _openDate, uint256 _buyFeePct, uint256 _sellFeePct) returns()
func (_TECGarden *TECGardenTransactor) CreateDaoTxThree(opts *bind.TransactOpts, _goal *big.Int, _period uint64, _exchangeRate *big.Int, _vestingCliffPeriod uint64, _vestingCompletePeriod uint64, _supplyOfferedPct *big.Int, _fundingForBeneficiaryPct *big.Int, _openDate uint64, _buyFeePct *big.Int, _sellFeePct *big.Int) (*types.Transaction, error) {
	return _TECGarden.contract.Transact(opts, "createDaoTxThree", _goal, _period, _exchangeRate, _vestingCliffPeriod, _vestingCompletePeriod, _supplyOfferedPct, _fundingForBeneficiaryPct, _openDate, _buyFeePct, _sellFeePct)
}

// CreateDaoTxThree is a paid mutator transaction binding the contract method 0xa74dd284.
//
// Solidity: function createDaoTxThree(uint256 _goal, uint64 _period, uint256 _exchangeRate, uint64 _vestingCliffPeriod, uint64 _vestingCompletePeriod, uint256 _supplyOfferedPct, uint256 _fundingForBeneficiaryPct, uint64 _openDate, uint256 _buyFeePct, uint256 _sellFeePct) returns()
func (_TECGarden *TECGardenSession) CreateDaoTxThree(_goal *big.Int, _period uint64, _exchangeRate *big.Int, _vestingCliffPeriod uint64, _vestingCompletePeriod uint64, _supplyOfferedPct *big.Int, _fundingForBeneficiaryPct *big.Int, _openDate uint64, _buyFeePct *big.Int, _sellFeePct *big.Int) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateDaoTxThree(&_TECGarden.TransactOpts, _goal, _period, _exchangeRate, _vestingCliffPeriod, _vestingCompletePeriod, _supplyOfferedPct, _fundingForBeneficiaryPct, _openDate, _buyFeePct, _sellFeePct)
}

// CreateDaoTxThree is a paid mutator transaction binding the contract method 0xa74dd284.
//
// Solidity: function createDaoTxThree(uint256 _goal, uint64 _period, uint256 _exchangeRate, uint64 _vestingCliffPeriod, uint64 _vestingCompletePeriod, uint256 _supplyOfferedPct, uint256 _fundingForBeneficiaryPct, uint64 _openDate, uint256 _buyFeePct, uint256 _sellFeePct) returns()
func (_TECGarden *TECGardenTransactorSession) CreateDaoTxThree(_goal *big.Int, _period uint64, _exchangeRate *big.Int, _vestingCliffPeriod uint64, _vestingCompletePeriod uint64, _supplyOfferedPct *big.Int, _fundingForBeneficiaryPct *big.Int, _openDate uint64, _buyFeePct *big.Int, _sellFeePct *big.Int) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateDaoTxThree(&_TECGarden.TransactOpts, _goal, _period, _exchangeRate, _vestingCliffPeriod, _vestingCompletePeriod, _supplyOfferedPct, _fundingForBeneficiaryPct, _openDate, _buyFeePct, _sellFeePct)
}

// CreateDaoTxTwo is a paid mutator transaction binding the contract method 0x813e6a99.
//
// Solidity: function createDaoTxTwo(address _tollgateFeeToken, uint256 _tollgateFeeAmount, address[] _redeemableTokens, uint256[4] _convictionSettings, address _collateralToken) returns()
func (_TECGarden *TECGardenTransactor) CreateDaoTxTwo(opts *bind.TransactOpts, _tollgateFeeToken common.Address, _tollgateFeeAmount *big.Int, _redeemableTokens []common.Address, _convictionSettings [4]*big.Int, _collateralToken common.Address) (*types.Transaction, error) {
	return _TECGarden.contract.Transact(opts, "createDaoTxTwo", _tollgateFeeToken, _tollgateFeeAmount, _redeemableTokens, _convictionSettings, _collateralToken)
}

// CreateDaoTxTwo is a paid mutator transaction binding the contract method 0x813e6a99.
//
// Solidity: function createDaoTxTwo(address _tollgateFeeToken, uint256 _tollgateFeeAmount, address[] _redeemableTokens, uint256[4] _convictionSettings, address _collateralToken) returns()
func (_TECGarden *TECGardenSession) CreateDaoTxTwo(_tollgateFeeToken common.Address, _tollgateFeeAmount *big.Int, _redeemableTokens []common.Address, _convictionSettings [4]*big.Int, _collateralToken common.Address) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateDaoTxTwo(&_TECGarden.TransactOpts, _tollgateFeeToken, _tollgateFeeAmount, _redeemableTokens, _convictionSettings, _collateralToken)
}

// CreateDaoTxTwo is a paid mutator transaction binding the contract method 0x813e6a99.
//
// Solidity: function createDaoTxTwo(address _tollgateFeeToken, uint256 _tollgateFeeAmount, address[] _redeemableTokens, uint256[4] _convictionSettings, address _collateralToken) returns()
func (_TECGarden *TECGardenTransactorSession) CreateDaoTxTwo(_tollgateFeeToken common.Address, _tollgateFeeAmount *big.Int, _redeemableTokens []common.Address, _convictionSettings [4]*big.Int, _collateralToken common.Address) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateDaoTxTwo(&_TECGarden.TransactOpts, _tollgateFeeToken, _tollgateFeeAmount, _redeemableTokens, _convictionSettings, _collateralToken)
}

// CreateTxTokenHolders is a paid mutator transaction binding the contract method 0xc9542822.
//
// Solidity: function createTxTokenHolders(address[] _holders, uint256[] _stakes, uint64 _openDate, uint64 _vestingCliffPeriod, uint64 _vestingCompletePeriod) returns()
func (_TECGarden *TECGardenTransactor) CreateTxTokenHolders(opts *bind.TransactOpts, _holders []common.Address, _stakes []*big.Int, _openDate uint64, _vestingCliffPeriod uint64, _vestingCompletePeriod uint64) (*types.Transaction, error) {
	return _TECGarden.contract.Transact(opts, "createTxTokenHolders", _holders, _stakes, _openDate, _vestingCliffPeriod, _vestingCompletePeriod)
}

// CreateTxTokenHolders is a paid mutator transaction binding the contract method 0xc9542822.
//
// Solidity: function createTxTokenHolders(address[] _holders, uint256[] _stakes, uint64 _openDate, uint64 _vestingCliffPeriod, uint64 _vestingCompletePeriod) returns()
func (_TECGarden *TECGardenSession) CreateTxTokenHolders(_holders []common.Address, _stakes []*big.Int, _openDate uint64, _vestingCliffPeriod uint64, _vestingCompletePeriod uint64) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateTxTokenHolders(&_TECGarden.TransactOpts, _holders, _stakes, _openDate, _vestingCliffPeriod, _vestingCompletePeriod)
}

// CreateTxTokenHolders is a paid mutator transaction binding the contract method 0xc9542822.
//
// Solidity: function createTxTokenHolders(address[] _holders, uint256[] _stakes, uint64 _openDate, uint64 _vestingCliffPeriod, uint64 _vestingCompletePeriod) returns()
func (_TECGarden *TECGardenTransactorSession) CreateTxTokenHolders(_holders []common.Address, _stakes []*big.Int, _openDate uint64, _vestingCliffPeriod uint64, _vestingCompletePeriod uint64) (*types.Transaction, error) {
	return _TECGarden.Contract.CreateTxTokenHolders(&_TECGarden.TransactOpts, _holders, _stakes, _openDate, _vestingCliffPeriod, _vestingCompletePeriod)
}

// TECGardenDeployDaoIterator is returned from FilterDeployDao and is used to iterate over the raw logs and unpacked data for DeployDao events raised by the TECGarden contract.
type TECGardenDeployDaoIterator struct {
	Event *TECGardenDeployDao // Event containing the contract specifics and raw log

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
func (it *TECGardenDeployDaoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TECGardenDeployDao)
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
		it.Event = new(TECGardenDeployDao)
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
func (it *TECGardenDeployDaoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TECGardenDeployDaoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TECGardenDeployDao represents a DeployDao event raised by the TECGarden contract.
type TECGardenDeployDao struct {
	Dao common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeployDao is a free log retrieval operation binding the contract event 0x0b13a9ab90735191cd544fd95ba68d1385144561cbdeb8acb8035de9a86432f5.
//
// Solidity: event DeployDao(address dao)
func (_TECGarden *TECGardenFilterer) FilterDeployDao(opts *bind.FilterOpts) (*TECGardenDeployDaoIterator, error) {

	logs, sub, err := _TECGarden.contract.FilterLogs(opts, "DeployDao")
	if err != nil {
		return nil, err
	}
	return &TECGardenDeployDaoIterator{contract: _TECGarden.contract, event: "DeployDao", logs: logs, sub: sub}, nil
}

// WatchDeployDao is a free log subscription operation binding the contract event 0x0b13a9ab90735191cd544fd95ba68d1385144561cbdeb8acb8035de9a86432f5.
//
// Solidity: event DeployDao(address dao)
func (_TECGarden *TECGardenFilterer) WatchDeployDao(opts *bind.WatchOpts, sink chan<- *TECGardenDeployDao) (event.Subscription, error) {

	logs, sub, err := _TECGarden.contract.WatchLogs(opts, "DeployDao")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TECGardenDeployDao)
				if err := _TECGarden.contract.UnpackLog(event, "DeployDao", log); err != nil {
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

// ParseDeployDao is a log parse operation binding the contract event 0x0b13a9ab90735191cd544fd95ba68d1385144561cbdeb8acb8035de9a86432f5.
//
// Solidity: event DeployDao(address dao)
func (_TECGarden *TECGardenFilterer) ParseDeployDao(log types.Log) (*TECGardenDeployDao, error) {
	event := new(TECGardenDeployDao)
	if err := _TECGarden.contract.UnpackLog(event, "DeployDao", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TECGardenDeployTokenIterator is returned from FilterDeployToken and is used to iterate over the raw logs and unpacked data for DeployToken events raised by the TECGarden contract.
type TECGardenDeployTokenIterator struct {
	Event *TECGardenDeployToken // Event containing the contract specifics and raw log

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
func (it *TECGardenDeployTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TECGardenDeployToken)
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
		it.Event = new(TECGardenDeployToken)
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
func (it *TECGardenDeployTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TECGardenDeployTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TECGardenDeployToken represents a DeployToken event raised by the TECGarden contract.
type TECGardenDeployToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeployToken is a free log retrieval operation binding the contract event 0xd18525bc6595a90cc21e9fbd517ada6fc07a7e87b7d2cdb6ee9284c450ebffa4.
//
// Solidity: event DeployToken(address token)
func (_TECGarden *TECGardenFilterer) FilterDeployToken(opts *bind.FilterOpts) (*TECGardenDeployTokenIterator, error) {

	logs, sub, err := _TECGarden.contract.FilterLogs(opts, "DeployToken")
	if err != nil {
		return nil, err
	}
	return &TECGardenDeployTokenIterator{contract: _TECGarden.contract, event: "DeployToken", logs: logs, sub: sub}, nil
}

// WatchDeployToken is a free log subscription operation binding the contract event 0xd18525bc6595a90cc21e9fbd517ada6fc07a7e87b7d2cdb6ee9284c450ebffa4.
//
// Solidity: event DeployToken(address token)
func (_TECGarden *TECGardenFilterer) WatchDeployToken(opts *bind.WatchOpts, sink chan<- *TECGardenDeployToken) (event.Subscription, error) {

	logs, sub, err := _TECGarden.contract.WatchLogs(opts, "DeployToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TECGardenDeployToken)
				if err := _TECGarden.contract.UnpackLog(event, "DeployToken", log); err != nil {
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

// ParseDeployToken is a log parse operation binding the contract event 0xd18525bc6595a90cc21e9fbd517ada6fc07a7e87b7d2cdb6ee9284c450ebffa4.
//
// Solidity: event DeployToken(address token)
func (_TECGarden *TECGardenFilterer) ParseDeployToken(log types.Log) (*TECGardenDeployToken, error) {
	event := new(TECGardenDeployToken)
	if err := _TECGarden.contract.UnpackLog(event, "DeployToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TECGardenInstalledAppIterator is returned from FilterInstalledApp and is used to iterate over the raw logs and unpacked data for InstalledApp events raised by the TECGarden contract.
type TECGardenInstalledAppIterator struct {
	Event *TECGardenInstalledApp // Event containing the contract specifics and raw log

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
func (it *TECGardenInstalledAppIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TECGardenInstalledApp)
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
		it.Event = new(TECGardenInstalledApp)
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
func (it *TECGardenInstalledAppIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TECGardenInstalledAppIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TECGardenInstalledApp represents a InstalledApp event raised by the TECGarden contract.
type TECGardenInstalledApp struct {
	AppProxy common.Address
	AppId    [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInstalledApp is a free log retrieval operation binding the contract event 0x2b183a501d4b1bbd30e6611ebac40ab18a00390e6c6bed324bf92a265c9ce6e3.
//
// Solidity: event InstalledApp(address appProxy, bytes32 appId)
func (_TECGarden *TECGardenFilterer) FilterInstalledApp(opts *bind.FilterOpts) (*TECGardenInstalledAppIterator, error) {

	logs, sub, err := _TECGarden.contract.FilterLogs(opts, "InstalledApp")
	if err != nil {
		return nil, err
	}
	return &TECGardenInstalledAppIterator{contract: _TECGarden.contract, event: "InstalledApp", logs: logs, sub: sub}, nil
}

// WatchInstalledApp is a free log subscription operation binding the contract event 0x2b183a501d4b1bbd30e6611ebac40ab18a00390e6c6bed324bf92a265c9ce6e3.
//
// Solidity: event InstalledApp(address appProxy, bytes32 appId)
func (_TECGarden *TECGardenFilterer) WatchInstalledApp(opts *bind.WatchOpts, sink chan<- *TECGardenInstalledApp) (event.Subscription, error) {

	logs, sub, err := _TECGarden.contract.WatchLogs(opts, "InstalledApp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TECGardenInstalledApp)
				if err := _TECGarden.contract.UnpackLog(event, "InstalledApp", log); err != nil {
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

// ParseInstalledApp is a log parse operation binding the contract event 0x2b183a501d4b1bbd30e6611ebac40ab18a00390e6c6bed324bf92a265c9ce6e3.
//
// Solidity: event InstalledApp(address appProxy, bytes32 appId)
func (_TECGarden *TECGardenFilterer) ParseInstalledApp(log types.Log) (*TECGardenInstalledApp, error) {
	event := new(TECGardenInstalledApp)
	if err := _TECGarden.contract.UnpackLog(event, "InstalledApp", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TECGardenSetupDaoIterator is returned from FilterSetupDao and is used to iterate over the raw logs and unpacked data for SetupDao events raised by the TECGarden contract.
type TECGardenSetupDaoIterator struct {
	Event *TECGardenSetupDao // Event containing the contract specifics and raw log

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
func (it *TECGardenSetupDaoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TECGardenSetupDao)
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
		it.Event = new(TECGardenSetupDao)
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
func (it *TECGardenSetupDaoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TECGardenSetupDaoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TECGardenSetupDao represents a SetupDao event raised by the TECGarden contract.
type TECGardenSetupDao struct {
	Dao common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetupDao is a free log retrieval operation binding the contract event 0x17592627a66846ce06d92a1708275bc653b2a3f34aec855584b819872a8ba413.
//
// Solidity: event SetupDao(address dao)
func (_TECGarden *TECGardenFilterer) FilterSetupDao(opts *bind.FilterOpts) (*TECGardenSetupDaoIterator, error) {

	logs, sub, err := _TECGarden.contract.FilterLogs(opts, "SetupDao")
	if err != nil {
		return nil, err
	}
	return &TECGardenSetupDaoIterator{contract: _TECGarden.contract, event: "SetupDao", logs: logs, sub: sub}, nil
}

// WatchSetupDao is a free log subscription operation binding the contract event 0x17592627a66846ce06d92a1708275bc653b2a3f34aec855584b819872a8ba413.
//
// Solidity: event SetupDao(address dao)
func (_TECGarden *TECGardenFilterer) WatchSetupDao(opts *bind.WatchOpts, sink chan<- *TECGardenSetupDao) (event.Subscription, error) {

	logs, sub, err := _TECGarden.contract.WatchLogs(opts, "SetupDao")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TECGardenSetupDao)
				if err := _TECGarden.contract.UnpackLog(event, "SetupDao", log); err != nil {
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

// ParseSetupDao is a log parse operation binding the contract event 0x17592627a66846ce06d92a1708275bc653b2a3f34aec855584b819872a8ba413.
//
// Solidity: event SetupDao(address dao)
func (_TECGarden *TECGardenFilterer) ParseSetupDao(log types.Log) (*TECGardenSetupDao, error) {
	event := new(TECGardenSetupDao)
	if err := _TECGarden.contract.UnpackLog(event, "SetupDao", log); err != nil {
		return nil, err
	}
	return event, nil
}
