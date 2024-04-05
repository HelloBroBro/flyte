// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// ExecutableArrayNode is an autogenerated mock type for the ExecutableArrayNode type
type ExecutableArrayNode struct {
	mock.Mock
}

type ExecutableArrayNode_GetMinSuccessRatio struct {
	*mock.Call
}

func (_m ExecutableArrayNode_GetMinSuccessRatio) Return(_a0 *float32) *ExecutableArrayNode_GetMinSuccessRatio {
	return &ExecutableArrayNode_GetMinSuccessRatio{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableArrayNode) OnGetMinSuccessRatio() *ExecutableArrayNode_GetMinSuccessRatio {
	c_call := _m.On("GetMinSuccessRatio")
	return &ExecutableArrayNode_GetMinSuccessRatio{Call: c_call}
}

func (_m *ExecutableArrayNode) OnGetMinSuccessRatioMatch(matchers ...interface{}) *ExecutableArrayNode_GetMinSuccessRatio {
	c_call := _m.On("GetMinSuccessRatio", matchers...)
	return &ExecutableArrayNode_GetMinSuccessRatio{Call: c_call}
}

// GetMinSuccessRatio provides a mock function with given fields:
func (_m *ExecutableArrayNode) GetMinSuccessRatio() *float32 {
	ret := _m.Called()

	var r0 *float32
	if rf, ok := ret.Get(0).(func() *float32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*float32)
		}
	}

	return r0
}

type ExecutableArrayNode_GetMinSuccesses struct {
	*mock.Call
}

func (_m ExecutableArrayNode_GetMinSuccesses) Return(_a0 *uint32) *ExecutableArrayNode_GetMinSuccesses {
	return &ExecutableArrayNode_GetMinSuccesses{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableArrayNode) OnGetMinSuccesses() *ExecutableArrayNode_GetMinSuccesses {
	c_call := _m.On("GetMinSuccesses")
	return &ExecutableArrayNode_GetMinSuccesses{Call: c_call}
}

func (_m *ExecutableArrayNode) OnGetMinSuccessesMatch(matchers ...interface{}) *ExecutableArrayNode_GetMinSuccesses {
	c_call := _m.On("GetMinSuccesses", matchers...)
	return &ExecutableArrayNode_GetMinSuccesses{Call: c_call}
}

// GetMinSuccesses provides a mock function with given fields:
func (_m *ExecutableArrayNode) GetMinSuccesses() *uint32 {
	ret := _m.Called()

	var r0 *uint32
	if rf, ok := ret.Get(0).(func() *uint32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*uint32)
		}
	}

	return r0
}

type ExecutableArrayNode_GetParallelism struct {
	*mock.Call
}

func (_m ExecutableArrayNode_GetParallelism) Return(_a0 int64) *ExecutableArrayNode_GetParallelism {
	return &ExecutableArrayNode_GetParallelism{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableArrayNode) OnGetParallelism() *ExecutableArrayNode_GetParallelism {
	c_call := _m.On("GetParallelism")
	return &ExecutableArrayNode_GetParallelism{Call: c_call}
}

func (_m *ExecutableArrayNode) OnGetParallelismMatch(matchers ...interface{}) *ExecutableArrayNode_GetParallelism {
	c_call := _m.On("GetParallelism", matchers...)
	return &ExecutableArrayNode_GetParallelism{Call: c_call}
}

// GetParallelism provides a mock function with given fields:
func (_m *ExecutableArrayNode) GetParallelism() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

type ExecutableArrayNode_GetSubNodeSpec struct {
	*mock.Call
}

func (_m ExecutableArrayNode_GetSubNodeSpec) Return(_a0 *v1alpha1.NodeSpec) *ExecutableArrayNode_GetSubNodeSpec {
	return &ExecutableArrayNode_GetSubNodeSpec{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableArrayNode) OnGetSubNodeSpec() *ExecutableArrayNode_GetSubNodeSpec {
	c_call := _m.On("GetSubNodeSpec")
	return &ExecutableArrayNode_GetSubNodeSpec{Call: c_call}
}

func (_m *ExecutableArrayNode) OnGetSubNodeSpecMatch(matchers ...interface{}) *ExecutableArrayNode_GetSubNodeSpec {
	c_call := _m.On("GetSubNodeSpec", matchers...)
	return &ExecutableArrayNode_GetSubNodeSpec{Call: c_call}
}

// GetSubNodeSpec provides a mock function with given fields:
func (_m *ExecutableArrayNode) GetSubNodeSpec() *v1alpha1.NodeSpec {
	ret := _m.Called()

	var r0 *v1alpha1.NodeSpec
	if rf, ok := ret.Get(0).(func() *v1alpha1.NodeSpec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.NodeSpec)
		}
	}

	return r0
}
