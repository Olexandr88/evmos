// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	common "github.com/ethereum/go-ethereum/common"

	context "context"

	core "github.com/ethereum/go-ethereum/core"

	evmtypes "github.com/evmos/evmos/v19/x/evm/types"

	mock "github.com/stretchr/testify/mock"

	statedb "github.com/evmos/evmos/v19/x/evm/statedb"

	types "github.com/cosmos/cosmos-sdk/types"

<<<<<<< HEAD
	vm "github.com/evmos/evmos/v19/x/evm/core/vm"
=======
	vm "github.com/evmos/evmos/v19/x/evm/core/vm"
>>>>>>> main
)

// EVMKeeper is an autogenerated mock type for the EVMKeeper type
type EVMKeeper struct {
	mock.Mock
}

// ApplyMessage provides a mock function with given fields: ctx, msg, tracer, commit
func (_m *EVMKeeper) ApplyMessage(ctx types.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error) {
	ret := _m.Called(ctx, msg, tracer, commit)

	if len(ret) == 0 {
		panic("no return value specified for ApplyMessage")
	}

	var r0 *evmtypes.MsgEthereumTxResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, core.Message, vm.EVMLogger, bool) (*evmtypes.MsgEthereumTxResponse, error)); ok {
		return rf(ctx, msg, tracer, commit)
	}
	if rf, ok := ret.Get(0).(func(types.Context, core.Message, vm.EVMLogger, bool) *evmtypes.MsgEthereumTxResponse); ok {
		r0 = rf(ctx, msg, tracer, commit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*evmtypes.MsgEthereumTxResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, core.Message, vm.EVMLogger, bool) error); ok {
		r1 = rf(ctx, msg, tracer, commit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CallEVM provides a mock function with given fields: ctx, _a1, from, contract, commit, method, args
func (_m *EVMKeeper) CallEVM(ctx types.Context, _a1 abi.ABI, from common.Address, contract common.Address, commit bool, method string, args ...interface{}) (*evmtypes.MsgEthereumTxResponse, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, _a1, from, contract, commit, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CallEVM")
	}

	var r0 *evmtypes.MsgEthereumTxResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, abi.ABI, common.Address, common.Address, bool, string, ...interface{}) (*evmtypes.MsgEthereumTxResponse, error)); ok {
		return rf(ctx, _a1, from, contract, commit, method, args...)
	}
	if rf, ok := ret.Get(0).(func(types.Context, abi.ABI, common.Address, common.Address, bool, string, ...interface{}) *evmtypes.MsgEthereumTxResponse); ok {
		r0 = rf(ctx, _a1, from, contract, commit, method, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*evmtypes.MsgEthereumTxResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, abi.ABI, common.Address, common.Address, bool, string, ...interface{}) error); ok {
		r1 = rf(ctx, _a1, from, contract, commit, method, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CallEVMWithData provides a mock function with given fields: ctx, from, contract, data, commit
func (_m *EVMKeeper) CallEVMWithData(ctx types.Context, from common.Address, contract *common.Address, data []byte, commit bool) (*evmtypes.MsgEthereumTxResponse, error) {
	ret := _m.Called(ctx, from, contract, data, commit)

	if len(ret) == 0 {
		panic("no return value specified for CallEVMWithData")
	}

	var r0 *evmtypes.MsgEthereumTxResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, common.Address, *common.Address, []byte, bool) (*evmtypes.MsgEthereumTxResponse, error)); ok {
		return rf(ctx, from, contract, data, commit)
	}
	if rf, ok := ret.Get(0).(func(types.Context, common.Address, *common.Address, []byte, bool) *evmtypes.MsgEthereumTxResponse); ok {
		r0 = rf(ctx, from, contract, data, commit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*evmtypes.MsgEthereumTxResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, common.Address, *common.Address, []byte, bool) error); ok {
		r1 = rf(ctx, from, contract, data, commit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccount provides a mock function with given fields: ctx, addr
func (_m *EVMKeeper) DeleteAccount(ctx types.Context, addr common.Address) error {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, common.Address) error); ok {
		r0 = rf(ctx, addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EstimateGasInternal provides a mock function with given fields: c, req, fromType
func (_m *EVMKeeper) EstimateGasInternal(c context.Context, req *evmtypes.EthCallRequest, fromType evmtypes.CallType) (*evmtypes.EstimateGasResponse, error) {
	ret := _m.Called(c, req, fromType)

	if len(ret) == 0 {
		panic("no return value specified for EstimateGasInternal")
	}

	var r0 *evmtypes.EstimateGasResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *evmtypes.EthCallRequest, evmtypes.CallType) (*evmtypes.EstimateGasResponse, error)); ok {
		return rf(c, req, fromType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *evmtypes.EthCallRequest, evmtypes.CallType) *evmtypes.EstimateGasResponse); ok {
		r0 = rf(c, req, fromType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*evmtypes.EstimateGasResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *evmtypes.EthCallRequest, evmtypes.CallType) error); ok {
		r1 = rf(c, req, fromType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountWithoutBalance provides a mock function with given fields: ctx, addr
func (_m *EVMKeeper) GetAccountWithoutBalance(ctx types.Context, addr common.Address) *statedb.Account {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountWithoutBalance")
	}

	var r0 *statedb.Account
	if rf, ok := ret.Get(0).(func(types.Context, common.Address) *statedb.Account); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*statedb.Account)
		}
	}

	return r0
}

// GetParams provides a mock function with given fields: ctx
func (_m *EVMKeeper) GetParams(ctx types.Context) evmtypes.Params {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetParams")
	}

	var r0 evmtypes.Params
	if rf, ok := ret.Get(0).(func(types.Context) evmtypes.Params); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(evmtypes.Params)
	}

	return r0
}

// IsAvailableStaticPrecompile provides a mock function with given fields: params, address
func (_m *EVMKeeper) IsAvailableStaticPrecompile(params *evmtypes.Params, address common.Address) bool {
	ret := _m.Called(params, address)

	if len(ret) == 0 {
		panic("no return value specified for IsAvailableStaticPrecompile")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*evmtypes.Params, common.Address) bool); ok {
		r0 = rf(params, address)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewEVMKeeper creates a new instance of EVMKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEVMKeeper(t interface {
	mock.TestingT
	Cleanup(func())
},
) *EVMKeeper {
	mock := &EVMKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
