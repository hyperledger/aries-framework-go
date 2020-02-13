// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service (interfaces: DIDComm,Event,Messenger,MessengerHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	reflect "reflect"
)

// MockDIDComm is a mock of DIDComm interface
type MockDIDComm struct {
	ctrl     *gomock.Controller
	recorder *MockDIDCommMockRecorder
}

// MockDIDCommMockRecorder is the mock recorder for MockDIDComm
type MockDIDCommMockRecorder struct {
	mock *MockDIDComm
}

// NewMockDIDComm creates a new mock instance
func NewMockDIDComm(ctrl *gomock.Controller) *MockDIDComm {
	mock := &MockDIDComm{ctrl: ctrl}
	mock.recorder = &MockDIDCommMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDIDComm) EXPECT() *MockDIDCommMockRecorder {
	return m.recorder
}

// HandleInbound mocks base method
func (m *MockDIDComm) HandleInbound(arg0 service.DIDCommMsg, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInbound", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleInbound indicates an expected call of HandleInbound
func (mr *MockDIDCommMockRecorder) HandleInbound(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInbound", reflect.TypeOf((*MockDIDComm)(nil).HandleInbound), arg0, arg1, arg2)
}

// HandleOutbound mocks base method
func (m *MockDIDComm) HandleOutbound(arg0 service.DIDCommMsg, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleOutbound", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleOutbound indicates an expected call of HandleOutbound
func (mr *MockDIDCommMockRecorder) HandleOutbound(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleOutbound", reflect.TypeOf((*MockDIDComm)(nil).HandleOutbound), arg0, arg1, arg2)
}

// RegisterActionEvent mocks base method
func (m *MockDIDComm) RegisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterActionEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterActionEvent indicates an expected call of RegisterActionEvent
func (mr *MockDIDCommMockRecorder) RegisterActionEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterActionEvent", reflect.TypeOf((*MockDIDComm)(nil).RegisterActionEvent), arg0)
}

// RegisterMsgEvent mocks base method
func (m *MockDIDComm) RegisterMsgEvent(arg0 chan<- service.StateMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMsgEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterMsgEvent indicates an expected call of RegisterMsgEvent
func (mr *MockDIDCommMockRecorder) RegisterMsgEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMsgEvent", reflect.TypeOf((*MockDIDComm)(nil).RegisterMsgEvent), arg0)
}

// UnregisterActionEvent mocks base method
func (m *MockDIDComm) UnregisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterActionEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterActionEvent indicates an expected call of UnregisterActionEvent
func (mr *MockDIDCommMockRecorder) UnregisterActionEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterActionEvent", reflect.TypeOf((*MockDIDComm)(nil).UnregisterActionEvent), arg0)
}

// UnregisterMsgEvent mocks base method
func (m *MockDIDComm) UnregisterMsgEvent(arg0 chan<- service.StateMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterMsgEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterMsgEvent indicates an expected call of UnregisterMsgEvent
func (mr *MockDIDCommMockRecorder) UnregisterMsgEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterMsgEvent", reflect.TypeOf((*MockDIDComm)(nil).UnregisterMsgEvent), arg0)
}

// MockEvent is a mock of Event interface
type MockEvent struct {
	ctrl     *gomock.Controller
	recorder *MockEventMockRecorder
}

// MockEventMockRecorder is the mock recorder for MockEvent
type MockEventMockRecorder struct {
	mock *MockEvent
}

// NewMockEvent creates a new mock instance
func NewMockEvent(ctrl *gomock.Controller) *MockEvent {
	mock := &MockEvent{ctrl: ctrl}
	mock.recorder = &MockEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEvent) EXPECT() *MockEventMockRecorder {
	return m.recorder
}

// RegisterActionEvent mocks base method
func (m *MockEvent) RegisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterActionEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterActionEvent indicates an expected call of RegisterActionEvent
func (mr *MockEventMockRecorder) RegisterActionEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterActionEvent", reflect.TypeOf((*MockEvent)(nil).RegisterActionEvent), arg0)
}

// RegisterMsgEvent mocks base method
func (m *MockEvent) RegisterMsgEvent(arg0 chan<- service.StateMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMsgEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterMsgEvent indicates an expected call of RegisterMsgEvent
func (mr *MockEventMockRecorder) RegisterMsgEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMsgEvent", reflect.TypeOf((*MockEvent)(nil).RegisterMsgEvent), arg0)
}

// UnregisterActionEvent mocks base method
func (m *MockEvent) UnregisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterActionEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterActionEvent indicates an expected call of UnregisterActionEvent
func (mr *MockEventMockRecorder) UnregisterActionEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterActionEvent", reflect.TypeOf((*MockEvent)(nil).UnregisterActionEvent), arg0)
}

// UnregisterMsgEvent mocks base method
func (m *MockEvent) UnregisterMsgEvent(arg0 chan<- service.StateMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterMsgEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterMsgEvent indicates an expected call of UnregisterMsgEvent
func (mr *MockEventMockRecorder) UnregisterMsgEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterMsgEvent", reflect.TypeOf((*MockEvent)(nil).UnregisterMsgEvent), arg0)
}

// MockMessenger is a mock of Messenger interface
type MockMessenger struct {
	ctrl     *gomock.Controller
	recorder *MockMessengerMockRecorder
}

// MockMessengerMockRecorder is the mock recorder for MockMessenger
type MockMessengerMockRecorder struct {
	mock *MockMessenger
}

// NewMockMessenger creates a new mock instance
func NewMockMessenger(ctrl *gomock.Controller) *MockMessenger {
	mock := &MockMessenger{ctrl: ctrl}
	mock.recorder = &MockMessengerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessenger) EXPECT() *MockMessengerMockRecorder {
	return m.recorder
}

// ReplyTo mocks base method
func (m *MockMessenger) ReplyTo(arg0 string, arg1 service.DIDCommMsgMap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyTo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplyTo indicates an expected call of ReplyTo
func (mr *MockMessengerMockRecorder) ReplyTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyTo", reflect.TypeOf((*MockMessenger)(nil).ReplyTo), arg0, arg1)
}

// ReplyToNested mocks base method
func (m *MockMessenger) ReplyToNested(arg0 string, arg1 service.DIDCommMsgMap, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyToNested", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplyToNested indicates an expected call of ReplyToNested
func (mr *MockMessengerMockRecorder) ReplyToNested(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyToNested", reflect.TypeOf((*MockMessenger)(nil).ReplyToNested), arg0, arg1, arg2, arg3)
}

// Send mocks base method
func (m *MockMessenger) Send(arg0 service.DIDCommMsgMap, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockMessengerMockRecorder) Send(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMessenger)(nil).Send), arg0, arg1, arg2)
}

// MockMessengerHandler is a mock of MessengerHandler interface
type MockMessengerHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMessengerHandlerMockRecorder
}

// MockMessengerHandlerMockRecorder is the mock recorder for MockMessengerHandler
type MockMessengerHandlerMockRecorder struct {
	mock *MockMessengerHandler
}

// NewMockMessengerHandler creates a new mock instance
func NewMockMessengerHandler(ctrl *gomock.Controller) *MockMessengerHandler {
	mock := &MockMessengerHandler{ctrl: ctrl}
	mock.recorder = &MockMessengerHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessengerHandler) EXPECT() *MockMessengerHandlerMockRecorder {
	return m.recorder
}

// HandleInbound mocks base method
func (m *MockMessengerHandler) HandleInbound(arg0 service.DIDCommMsgMap, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInbound", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleInbound indicates an expected call of HandleInbound
func (mr *MockMessengerHandlerMockRecorder) HandleInbound(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInbound", reflect.TypeOf((*MockMessengerHandler)(nil).HandleInbound), arg0, arg1, arg2)
}

// ReplyTo mocks base method
func (m *MockMessengerHandler) ReplyTo(arg0 string, arg1 service.DIDCommMsgMap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyTo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplyTo indicates an expected call of ReplyTo
func (mr *MockMessengerHandlerMockRecorder) ReplyTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyTo", reflect.TypeOf((*MockMessengerHandler)(nil).ReplyTo), arg0, arg1)
}

// ReplyToNested mocks base method
func (m *MockMessengerHandler) ReplyToNested(arg0 string, arg1 service.DIDCommMsgMap, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplyToNested", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplyToNested indicates an expected call of ReplyToNested
func (mr *MockMessengerHandlerMockRecorder) ReplyToNested(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyToNested", reflect.TypeOf((*MockMessengerHandler)(nil).ReplyToNested), arg0, arg1, arg2, arg3)
}

// Send mocks base method
func (m *MockMessengerHandler) Send(arg0 service.DIDCommMsgMap, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockMessengerHandlerMockRecorder) Send(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMessengerHandler)(nil).Send), arg0, arg1, arg2)
}
