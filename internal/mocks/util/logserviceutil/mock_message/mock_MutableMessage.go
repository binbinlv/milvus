// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_message

import (
	message "github.com/milvus-io/milvus/internal/util/logserviceutil/message"
	mock "github.com/stretchr/testify/mock"
)

// MockMutableMessage is an autogenerated mock type for the MutableMessage type
type MockMutableMessage struct {
	mock.Mock
}

type MockMutableMessage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMutableMessage) EXPECT() *MockMutableMessage_Expecter {
	return &MockMutableMessage_Expecter{mock: &_m.Mock}
}

// EstimateSize provides a mock function with given fields:
func (_m *MockMutableMessage) EstimateSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockMutableMessage_EstimateSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateSize'
type MockMutableMessage_EstimateSize_Call struct {
	*mock.Call
}

// EstimateSize is a helper method to define mock.On call
func (_e *MockMutableMessage_Expecter) EstimateSize() *MockMutableMessage_EstimateSize_Call {
	return &MockMutableMessage_EstimateSize_Call{Call: _e.mock.On("EstimateSize")}
}

func (_c *MockMutableMessage_EstimateSize_Call) Run(run func()) *MockMutableMessage_EstimateSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMutableMessage_EstimateSize_Call) Return(_a0 int) *MockMutableMessage_EstimateSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMutableMessage_EstimateSize_Call) RunAndReturn(run func() int) *MockMutableMessage_EstimateSize_Call {
	_c.Call.Return(run)
	return _c
}

// MessageType provides a mock function with given fields:
func (_m *MockMutableMessage) MessageType() message.MessageType {
	ret := _m.Called()

	var r0 message.MessageType
	if rf, ok := ret.Get(0).(func() message.MessageType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(message.MessageType)
	}

	return r0
}

// MockMutableMessage_MessageType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MessageType'
type MockMutableMessage_MessageType_Call struct {
	*mock.Call
}

// MessageType is a helper method to define mock.On call
func (_e *MockMutableMessage_Expecter) MessageType() *MockMutableMessage_MessageType_Call {
	return &MockMutableMessage_MessageType_Call{Call: _e.mock.On("MessageType")}
}

func (_c *MockMutableMessage_MessageType_Call) Run(run func()) *MockMutableMessage_MessageType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMutableMessage_MessageType_Call) Return(_a0 message.MessageType) *MockMutableMessage_MessageType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMutableMessage_MessageType_Call) RunAndReturn(run func() message.MessageType) *MockMutableMessage_MessageType_Call {
	_c.Call.Return(run)
	return _c
}

// Payload provides a mock function with given fields:
func (_m *MockMutableMessage) Payload() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockMutableMessage_Payload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Payload'
type MockMutableMessage_Payload_Call struct {
	*mock.Call
}

// Payload is a helper method to define mock.On call
func (_e *MockMutableMessage_Expecter) Payload() *MockMutableMessage_Payload_Call {
	return &MockMutableMessage_Payload_Call{Call: _e.mock.On("Payload")}
}

func (_c *MockMutableMessage_Payload_Call) Run(run func()) *MockMutableMessage_Payload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMutableMessage_Payload_Call) Return(_a0 []byte) *MockMutableMessage_Payload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMutableMessage_Payload_Call) RunAndReturn(run func() []byte) *MockMutableMessage_Payload_Call {
	_c.Call.Return(run)
	return _c
}

// Properties provides a mock function with given fields:
func (_m *MockMutableMessage) Properties() message.Properties {
	ret := _m.Called()

	var r0 message.Properties
	if rf, ok := ret.Get(0).(func() message.Properties); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(message.Properties)
		}
	}

	return r0
}

// MockMutableMessage_Properties_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Properties'
type MockMutableMessage_Properties_Call struct {
	*mock.Call
}

// Properties is a helper method to define mock.On call
func (_e *MockMutableMessage_Expecter) Properties() *MockMutableMessage_Properties_Call {
	return &MockMutableMessage_Properties_Call{Call: _e.mock.On("Properties")}
}

func (_c *MockMutableMessage_Properties_Call) Run(run func()) *MockMutableMessage_Properties_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMutableMessage_Properties_Call) Return(_a0 message.Properties) *MockMutableMessage_Properties_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMutableMessage_Properties_Call) RunAndReturn(run func() message.Properties) *MockMutableMessage_Properties_Call {
	_c.Call.Return(run)
	return _c
}

// WithTimeTick provides a mock function with given fields: tt
func (_m *MockMutableMessage) WithTimeTick(tt uint64) message.MutableMessage {
	ret := _m.Called(tt)

	var r0 message.MutableMessage
	if rf, ok := ret.Get(0).(func(uint64) message.MutableMessage); ok {
		r0 = rf(tt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(message.MutableMessage)
		}
	}

	return r0
}

// MockMutableMessage_WithTimeTick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTimeTick'
type MockMutableMessage_WithTimeTick_Call struct {
	*mock.Call
}

// WithTimeTick is a helper method to define mock.On call
//   - tt uint64
func (_e *MockMutableMessage_Expecter) WithTimeTick(tt interface{}) *MockMutableMessage_WithTimeTick_Call {
	return &MockMutableMessage_WithTimeTick_Call{Call: _e.mock.On("WithTimeTick", tt)}
}

func (_c *MockMutableMessage_WithTimeTick_Call) Run(run func(tt uint64)) *MockMutableMessage_WithTimeTick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockMutableMessage_WithTimeTick_Call) Return(_a0 message.MutableMessage) *MockMutableMessage_WithTimeTick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMutableMessage_WithTimeTick_Call) RunAndReturn(run func(uint64) message.MutableMessage) *MockMutableMessage_WithTimeTick_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMutableMessage creates a new instance of MockMutableMessage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMutableMessage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMutableMessage {
	mock := &MockMutableMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
