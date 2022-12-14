// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nats-io/stan.go (interfaces: Conn)

// Package mock_stan_go is a generated GoMock package.
package mock_stan_go

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

// MockConn is a mock of Conn interface.
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *MockConnMockRecorder
}

// MockConnMockRecorder is the mock recorder for MockConn.
type MockConnMockRecorder struct {
	mock *MockConn
}

// NewMockConn creates a new mock instance.
func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &MockConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConn) EXPECT() *MockConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConn)(nil).Close))
}

// NatsConn mocks base method.
func (m *MockConn) NatsConn() *nats.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NatsConn")
	ret0, _ := ret[0].(*nats.Conn)
	return ret0
}

// NatsConn indicates an expected call of NatsConn.
func (mr *MockConnMockRecorder) NatsConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NatsConn", reflect.TypeOf((*MockConn)(nil).NatsConn))
}

// Publish mocks base method.
func (m *MockConn) Publish(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockConnMockRecorder) Publish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockConn)(nil).Publish), arg0, arg1)
}

// PublishAsync mocks base method.
func (m *MockConn) PublishAsync(arg0 string, arg1 []byte, arg2 stan.AckHandler) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishAsync", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishAsync indicates an expected call of PublishAsync.
func (mr *MockConnMockRecorder) PublishAsync(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAsync", reflect.TypeOf((*MockConn)(nil).PublishAsync), arg0, arg1, arg2)
}

// QueueSubscribe mocks base method.
func (m *MockConn) QueueSubscribe(arg0, arg1 string, arg2 stan.MsgHandler, arg3 ...stan.SubscriptionOption) (stan.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueueSubscribe", varargs...)
	ret0, _ := ret[0].(stan.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueSubscribe indicates an expected call of QueueSubscribe.
func (mr *MockConnMockRecorder) QueueSubscribe(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueSubscribe", reflect.TypeOf((*MockConn)(nil).QueueSubscribe), varargs...)
}

// Subscribe mocks base method.
func (m *MockConn) Subscribe(arg0 string, arg1 stan.MsgHandler, arg2 ...stan.SubscriptionOption) (stan.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(stan.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockConnMockRecorder) Subscribe(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockConn)(nil).Subscribe), varargs...)
}
