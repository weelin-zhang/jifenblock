// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jifen

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

// JifenABI is the input ABI used to generate the binding from.
const JifenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"txid\",\"type\":\"bytes32\"},{\"name\":\"memberOID\",\"type\":\"bytes32\"},{\"name\":\"companyOID\",\"type\":\"string\"},{\"name\":\"siteOID\",\"type\":\"string\"},{\"name\":\"createTime\",\"type\":\"string\"},{\"name\":\"updateTime\",\"type\":\"string\"},{\"name\":\"integral\",\"type\":\"string\"},{\"name\":\"types_deductibleAmount_consumeAmount_goods_goodsQty\",\"type\":\"string\"}],\"name\":\"deductJiFen_flow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tx_details\",\"outputs\":[{\"name\":\"txid\",\"type\":\"bytes32\"},{\"name\":\"memberOID\",\"type\":\"bytes32\"},{\"name\":\"companyOID\",\"type\":\"string\"},{\"name\":\"siteOID\",\"type\":\"string\"},{\"name\":\"createTime\",\"type\":\"string\"},{\"name\":\"updateTime\",\"type\":\"string\"},{\"name\":\"integral\",\"type\":\"string\"},{\"name\":\"types_deductibleAmount_consumeAmount_goods_goodsQty\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_txs\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"memberOID\",\"type\":\"bytes32\"}],\"name\":\"getUser_txsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"txid\",\"type\":\"bytes32\"},{\"name\":\"memberOID\",\"type\":\"bytes32\"},{\"name\":\"companyOID\",\"type\":\"string\"},{\"name\":\"siteOID\",\"type\":\"string\"},{\"name\":\"createTime\",\"type\":\"string\"},{\"name\":\"updateTime\",\"type\":\"string\"},{\"name\":\"integral\",\"type\":\"string\"},{\"name\":\"types_deductibleAmount_consumeAmount_goods_goodsQty\",\"type\":\"string\"}],\"name\":\"distributeJiFen_flow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"memberOID\",\"type\":\"bytes32\"}],\"name\":\"Instructor\",\"type\":\"event\"}]"

// JifenBin is the compiled bytecode used for deploying new contracts.
const JifenBin = `608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610f29806100326000396000f30060806040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630124a14981146100925780633dc59c16146102ae5780636c7f1542146105545780638da5cb5b1461057e578063e3843db1146105af578063e82b0886146105ca578063f2fde38b146105e2578063f69ff64e14610605575b600080fd5b34801561009e57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261022b948035946024803595604435953695608494930191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061079e9650505050505050565b604051808315151515815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561027257818101518382015260200161025a565b50505050905090810190601f16801561029f5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156102ba57600080fd5b506102c66004356109de565b6040805189815260208082018a905261010092820183815289519383019390935288519192916060840191608085019160a086019160c087019160e0880191610120890191908f019080838360005b8381101561032d578181015183820152602001610315565b50505050905090810190601f16801561035a5780820380516001836020036101000a031916815260200191505b5087810386528c5181528c516020918201918e019080838360005b8381101561038d578181015183820152602001610375565b50505050905090810190601f1680156103ba5780820380516001836020036101000a031916815260200191505b5087810385528b5181528b516020918201918d019080838360005b838110156103ed5781810151838201526020016103d5565b50505050905090810190601f16801561041a5780820380516001836020036101000a031916815260200191505b5087810384528a5181528a516020918201918c019080838360005b8381101561044d578181015183820152602001610435565b50505050905090810190601f16801561047a5780820380516001836020036101000a031916815260200191505b5087810383528951815289516020918201918b019080838360005b838110156104ad578181015183820152602001610495565b50505050905090810190601f1680156104da5780820380516001836020036101000a031916815260200191505b5087810382528851815288516020918201918a019080838360005b8381101561050d5781810151838201526020016104f5565b50505050905090810190601f16801561053a5780820380516001836020036101000a031916815260200191505b509e50505050505050505050505050505060405180910390f35b34801561056057600080fd5b5061056c600435610d57565b60408051918252519081900360200190f35b34801561058a57600080fd5b50610593610d69565b60408051600160a060020a039092168252519081900360200190f35b3480156105bb57600080fd5b5061056c600435602435610d78565b3480156105d657600080fd5b5061056c600435610da8565b3480156105ee57600080fd5b50610603600160a060020a0360043516610dba565b005b34801561061157600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261022b948035946024803595604435953695608494930191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750610e009650505050505050565b60008054606090600160a060020a031633146107b957600080fd5b6000898152600160205260409020548b11156107d457600080fd5b6000898152600160205260409020546107f3908c63ffffffff610e3a16565b600160008b6000191660001916815260200190815260200160002081905550610100604051908101604052808b6000191681526020018a60001916815260200189815260200188815260200187815260200186815260200185815260200184815250600360008c60001916600019168152602001908152602001600020600082015181600001906000191690556020820151816001019060001916905560408201518160020190805190602001906108ac929190610e62565b50606082015180516108c8916003840191602090910190610e62565b50608082015180516108e4916004840191602090910190610e62565b5060a08201518051610900916005840191602090910190610e62565b5060c0820151805161091c916006840191602090910190610e62565b5060e08201518051610938916007840191602090910190610e62565b5050604080518c8152602081018c905281517fbec0f0a63d138d2b144ed4730f49d113e75a3a6a2df260d840bd9b0a0b0277b793509081900390910190a15050506000958652505060026020908152604080862080546001808201835591885296839020909601969096558551808701909652600786527f7375636365737300000000000000000000000000000000000000000000000000908601525091949293505050565b60036020908152600091825260409182902080546001808301546002808501805488519481161561010002600019011691909104601f810187900487028401870190975286835292959094919291830182828015610a7d5780601f10610a5257610100808354040283529160200191610a7d565b820191906000526020600020905b815481529060010190602001808311610a6057829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610b0d5780601f10610ae257610100808354040283529160200191610b0d565b820191906000526020600020905b815481529060010190602001808311610af057829003601f168201915b5050505060048301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610b9d5780601f10610b7257610100808354040283529160200191610b9d565b820191906000526020600020905b815481529060010190602001808311610b8057829003601f168201915b5050505060058301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610c2d5780601f10610c0257610100808354040283529160200191610c2d565b820191906000526020600020905b815481529060010190602001808311610c1057829003601f168201915b5050505060068301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610cbd5780601f10610c9257610100808354040283529160200191610cbd565b820191906000526020600020905b815481529060010190602001808311610ca057829003601f168201915b5050505060078301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610d4d5780601f10610d2257610100808354040283529160200191610d4d565b820191906000526020600020905b815481529060010190602001808311610d3057829003601f168201915b5050505050905088565b60016020526000908152604090205481565b600054600160a060020a031681565b600260205281600052604060002081815481101515610d9357fe5b90600052602060002001600091509150505481565b60009081526002602052604090205490565b600054600160a060020a03163314610dd157600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008054606090600160a060020a03163314610e1b57600080fd5b6000898152600160205260409020546107f3908c63ffffffff610e4c16565b600082821115610e4657fe5b50900390565b600082820183811015610e5b57fe5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ea357805160ff1916838001178555610ed0565b82800160010185558215610ed0579182015b82811115610ed0578251825591602001919060010190610eb5565b50610edc929150610ee0565b5090565b610efa91905b80821115610edc5760008155600101610ee6565b905600a165627a7a7230582090a1d248f232fe0cdfa79ebb48570f41a9e0d73d18d88744077f7bba86b6c4a40029`

// DeployJifen deploys a new Ethereum contract, binding an instance of Jifen to it.
func DeployJifen(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Jifen, error) {
	parsed, err := abi.JSON(strings.NewReader(JifenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JifenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jifen{JifenCaller: JifenCaller{contract: contract}, JifenTransactor: JifenTransactor{contract: contract}, JifenFilterer: JifenFilterer{contract: contract}}, nil
}

// Jifen is an auto generated Go binding around an Ethereum contract.
type Jifen struct {
	JifenCaller     // Read-only binding to the contract
	JifenTransactor // Write-only binding to the contract
	JifenFilterer   // Log filterer for contract events
}

// JifenCaller is an auto generated read-only Go binding around an Ethereum contract.
type JifenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JifenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JifenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JifenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JifenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JifenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JifenSession struct {
	Contract     *Jifen            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JifenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JifenCallerSession struct {
	Contract *JifenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JifenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JifenTransactorSession struct {
	Contract     *JifenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JifenRaw is an auto generated low-level Go binding around an Ethereum contract.
type JifenRaw struct {
	Contract *Jifen // Generic contract binding to access the raw methods on
}

// JifenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JifenCallerRaw struct {
	Contract *JifenCaller // Generic read-only contract binding to access the raw methods on
}

// JifenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JifenTransactorRaw struct {
	Contract *JifenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJifen creates a new instance of Jifen, bound to a specific deployed contract.
func NewJifen(address common.Address, backend bind.ContractBackend) (*Jifen, error) {
	contract, err := bindJifen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jifen{JifenCaller: JifenCaller{contract: contract}, JifenTransactor: JifenTransactor{contract: contract}, JifenFilterer: JifenFilterer{contract: contract}}, nil
}

// NewJifenCaller creates a new read-only instance of Jifen, bound to a specific deployed contract.
func NewJifenCaller(address common.Address, caller bind.ContractCaller) (*JifenCaller, error) {
	contract, err := bindJifen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JifenCaller{contract: contract}, nil
}

// NewJifenTransactor creates a new write-only instance of Jifen, bound to a specific deployed contract.
func NewJifenTransactor(address common.Address, transactor bind.ContractTransactor) (*JifenTransactor, error) {
	contract, err := bindJifen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JifenTransactor{contract: contract}, nil
}

// NewJifenFilterer creates a new log filterer instance of Jifen, bound to a specific deployed contract.
func NewJifenFilterer(address common.Address, filterer bind.ContractFilterer) (*JifenFilterer, error) {
	contract, err := bindJifen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JifenFilterer{contract: contract}, nil
}

// bindJifen binds a generic wrapper to an already deployed contract.
func bindJifen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JifenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jifen *JifenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Jifen.Contract.JifenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jifen *JifenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jifen.Contract.JifenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jifen *JifenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jifen.Contract.JifenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jifen *JifenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Jifen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jifen *JifenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jifen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jifen *JifenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jifen.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x6c7f1542.
//
// Solidity: function balanceOf( bytes32) constant returns(uint256)
func (_Jifen *JifenCaller) BalanceOf(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Jifen.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x6c7f1542.
//
// Solidity: function balanceOf( bytes32) constant returns(uint256)
func (_Jifen *JifenSession) BalanceOf(arg0 [32]byte) (*big.Int, error) {
	return _Jifen.Contract.BalanceOf(&_Jifen.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x6c7f1542.
//
// Solidity: function balanceOf( bytes32) constant returns(uint256)
func (_Jifen *JifenCallerSession) BalanceOf(arg0 [32]byte) (*big.Int, error) {
	return _Jifen.Contract.BalanceOf(&_Jifen.CallOpts, arg0)
}

// GetUserTxsLength is a free data retrieval call binding the contract method 0xe82b0886.
//
// Solidity: function getUser_txsLength(memberOID bytes32) constant returns(uint256)
func (_Jifen *JifenCaller) GetUserTxsLength(opts *bind.CallOpts, memberOID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Jifen.contract.Call(opts, out, "getUser_txsLength", memberOID)
	return *ret0, err
}

// GetUserTxsLength is a free data retrieval call binding the contract method 0xe82b0886.
//
// Solidity: function getUser_txsLength(memberOID bytes32) constant returns(uint256)
func (_Jifen *JifenSession) GetUserTxsLength(memberOID [32]byte) (*big.Int, error) {
	return _Jifen.Contract.GetUserTxsLength(&_Jifen.CallOpts, memberOID)
}

// GetUserTxsLength is a free data retrieval call binding the contract method 0xe82b0886.
//
// Solidity: function getUser_txsLength(memberOID bytes32) constant returns(uint256)
func (_Jifen *JifenCallerSession) GetUserTxsLength(memberOID [32]byte) (*big.Int, error) {
	return _Jifen.Contract.GetUserTxsLength(&_Jifen.CallOpts, memberOID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Jifen *JifenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Jifen.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Jifen *JifenSession) Owner() (common.Address, error) {
	return _Jifen.Contract.Owner(&_Jifen.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Jifen *JifenCallerSession) Owner() (common.Address, error) {
	return _Jifen.Contract.Owner(&_Jifen.CallOpts)
}

// TxDetails is a free data retrieval call binding the contract method 0x3dc59c16.
//
// Solidity: function tx_details( bytes32) constant returns(txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string)
func (_Jifen *JifenCaller) TxDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Txid                                            [32]byte
	MemberOID                                       [32]byte
	CompanyOID                                      string
	SiteOID                                         string
	CreateTime                                      string
	UpdateTime                                      string
	Integral                                        string
	TypesDeductibleAmountConsumeAmountGoodsGoodsQty string
}, error) {
	ret := new(struct {
		Txid                                            [32]byte
		MemberOID                                       [32]byte
		CompanyOID                                      string
		SiteOID                                         string
		CreateTime                                      string
		UpdateTime                                      string
		Integral                                        string
		TypesDeductibleAmountConsumeAmountGoodsGoodsQty string
	})
	out := ret
	err := _Jifen.contract.Call(opts, out, "tx_details", arg0)
	return *ret, err
}

// TxDetails is a free data retrieval call binding the contract method 0x3dc59c16.
//
// Solidity: function tx_details( bytes32) constant returns(txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string)
func (_Jifen *JifenSession) TxDetails(arg0 [32]byte) (struct {
	Txid                                            [32]byte
	MemberOID                                       [32]byte
	CompanyOID                                      string
	SiteOID                                         string
	CreateTime                                      string
	UpdateTime                                      string
	Integral                                        string
	TypesDeductibleAmountConsumeAmountGoodsGoodsQty string
}, error) {
	return _Jifen.Contract.TxDetails(&_Jifen.CallOpts, arg0)
}

// TxDetails is a free data retrieval call binding the contract method 0x3dc59c16.
//
// Solidity: function tx_details( bytes32) constant returns(txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string)
func (_Jifen *JifenCallerSession) TxDetails(arg0 [32]byte) (struct {
	Txid                                            [32]byte
	MemberOID                                       [32]byte
	CompanyOID                                      string
	SiteOID                                         string
	CreateTime                                      string
	UpdateTime                                      string
	Integral                                        string
	TypesDeductibleAmountConsumeAmountGoodsGoodsQty string
}, error) {
	return _Jifen.Contract.TxDetails(&_Jifen.CallOpts, arg0)
}

// UserTxs is a free data retrieval call binding the contract method 0xe3843db1.
//
// Solidity: function user_txs( bytes32,  uint256) constant returns(bytes32)
func (_Jifen *JifenCaller) UserTxs(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Jifen.contract.Call(opts, out, "user_txs", arg0, arg1)
	return *ret0, err
}

// UserTxs is a free data retrieval call binding the contract method 0xe3843db1.
//
// Solidity: function user_txs( bytes32,  uint256) constant returns(bytes32)
func (_Jifen *JifenSession) UserTxs(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _Jifen.Contract.UserTxs(&_Jifen.CallOpts, arg0, arg1)
}

// UserTxs is a free data retrieval call binding the contract method 0xe3843db1.
//
// Solidity: function user_txs( bytes32,  uint256) constant returns(bytes32)
func (_Jifen *JifenCallerSession) UserTxs(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _Jifen.Contract.UserTxs(&_Jifen.CallOpts, arg0, arg1)
}

// DeductJiFenFlow is a paid mutator transaction binding the contract method 0x0124a149.
//
// Solidity: function deductJiFen_flow(value uint256, txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) returns(bool, string)
func (_Jifen *JifenTransactor) DeductJiFenFlow(opts *bind.TransactOpts, value *big.Int, txid [32]byte, memberOID [32]byte, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) (*types.Transaction, error) {
	return _Jifen.contract.Transact(opts, "deductJiFen_flow", value, txid, memberOID, companyOID, siteOID, createTime, updateTime, integral, types_deductibleAmount_consumeAmount_goods_goodsQty)
}

// DeductJiFenFlow is a paid mutator transaction binding the contract method 0x0124a149.
//
// Solidity: function deductJiFen_flow(value uint256, txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) returns(bool, string)
func (_Jifen *JifenSession) DeductJiFenFlow(value *big.Int, txid [32]byte, memberOID [32]byte, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) (*types.Transaction, error) {
	return _Jifen.Contract.DeductJiFenFlow(&_Jifen.TransactOpts, value, txid, memberOID, companyOID, siteOID, createTime, updateTime, integral, types_deductibleAmount_consumeAmount_goods_goodsQty)
}

// DeductJiFenFlow is a paid mutator transaction binding the contract method 0x0124a149.
//
// Solidity: function deductJiFen_flow(value uint256, txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) returns(bool, string)
func (_Jifen *JifenTransactorSession) DeductJiFenFlow(value *big.Int, txid [32]byte, memberOID [32]byte, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) (*types.Transaction, error) {
	return _Jifen.Contract.DeductJiFenFlow(&_Jifen.TransactOpts, value, txid, memberOID, companyOID, siteOID, createTime, updateTime, integral, types_deductibleAmount_consumeAmount_goods_goodsQty)
}

// DistributeJiFenFlow is a paid mutator transaction binding the contract method 0xf69ff64e.
//
// Solidity: function distributeJiFen_flow(value uint256, txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) returns(bool, string)
func (_Jifen *JifenTransactor) DistributeJiFenFlow(opts *bind.TransactOpts, value *big.Int, txid [32]byte, memberOID [32]byte, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) (*types.Transaction, error) {
	return _Jifen.contract.Transact(opts, "distributeJiFen_flow", value, txid, memberOID, companyOID, siteOID, createTime, updateTime, integral, types_deductibleAmount_consumeAmount_goods_goodsQty)
}

// DistributeJiFenFlow is a paid mutator transaction binding the contract method 0xf69ff64e.
//
// Solidity: function distributeJiFen_flow(value uint256, txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) returns(bool, string)
func (_Jifen *JifenSession) DistributeJiFenFlow(value *big.Int, txid [32]byte, memberOID [32]byte, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) (*types.Transaction, error) {
	return _Jifen.Contract.DistributeJiFenFlow(&_Jifen.TransactOpts, value, txid, memberOID, companyOID, siteOID, createTime, updateTime, integral, types_deductibleAmount_consumeAmount_goods_goodsQty)
}

// DistributeJiFenFlow is a paid mutator transaction binding the contract method 0xf69ff64e.
//
// Solidity: function distributeJiFen_flow(value uint256, txid bytes32, memberOID bytes32, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) returns(bool, string)
func (_Jifen *JifenTransactorSession) DistributeJiFenFlow(value *big.Int, txid [32]byte, memberOID [32]byte, companyOID string, siteOID string, createTime string, updateTime string, integral string, types_deductibleAmount_consumeAmount_goods_goodsQty string) (*types.Transaction, error) {
	return _Jifen.Contract.DistributeJiFenFlow(&_Jifen.TransactOpts, value, txid, memberOID, companyOID, siteOID, createTime, updateTime, integral, types_deductibleAmount_consumeAmount_goods_goodsQty)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Jifen *JifenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Jifen.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Jifen *JifenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jifen.Contract.TransferOwnership(&_Jifen.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Jifen *JifenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jifen.Contract.TransferOwnership(&_Jifen.TransactOpts, newOwner)
}

// JifenInstructorIterator is returned from FilterInstructor and is used to iterate over the raw logs and unpacked data for Instructor events raised by the Jifen contract.
type JifenInstructorIterator struct {
	Event *JifenInstructor // Event containing the contract specifics and raw log

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
func (it *JifenInstructorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JifenInstructor)
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
		it.Event = new(JifenInstructor)
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
func (it *JifenInstructorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JifenInstructorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JifenInstructor represents a Instructor event raised by the Jifen contract.
type JifenInstructor struct {
	Txid      [32]byte
	MemberOID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInstructor is a free log retrieval operation binding the contract event 0xbec0f0a63d138d2b144ed4730f49d113e75a3a6a2df260d840bd9b0a0b0277b7.
//
// Solidity: e Instructor(txid bytes32, memberOID bytes32)
func (_Jifen *JifenFilterer) FilterInstructor(opts *bind.FilterOpts) (*JifenInstructorIterator, error) {

	logs, sub, err := _Jifen.contract.FilterLogs(opts, "Instructor")
	if err != nil {
		return nil, err
	}
	return &JifenInstructorIterator{contract: _Jifen.contract, event: "Instructor", logs: logs, sub: sub}, nil
}

// WatchInstructor is a free log subscription operation binding the contract event 0xbec0f0a63d138d2b144ed4730f49d113e75a3a6a2df260d840bd9b0a0b0277b7.
//
// Solidity: e Instructor(txid bytes32, memberOID bytes32)
func (_Jifen *JifenFilterer) WatchInstructor(opts *bind.WatchOpts, sink chan<- *JifenInstructor) (event.Subscription, error) {

	logs, sub, err := _Jifen.contract.WatchLogs(opts, "Instructor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JifenInstructor)
				if err := _Jifen.contract.UnpackLog(event, "Instructor", log); err != nil {
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
