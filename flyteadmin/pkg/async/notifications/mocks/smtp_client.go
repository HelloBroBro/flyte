// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	io "io"
	smtp "net/smtp"

	mock "github.com/stretchr/testify/mock"

	tls "crypto/tls"
)

// SMTPClient is an autogenerated mock type for the SMTPClient type
type SMTPClient struct {
	mock.Mock
}

type SMTPClient_Auth struct {
	*mock.Call
}

func (_m SMTPClient_Auth) Return(_a0 error) *SMTPClient_Auth {
	return &SMTPClient_Auth{Call: _m.Call.Return(_a0)}
}

func (_m *SMTPClient) OnAuth(a smtp.Auth) *SMTPClient_Auth {
	c_call := _m.On("Auth", a)
	return &SMTPClient_Auth{Call: c_call}
}

func (_m *SMTPClient) OnAuthMatch(matchers ...interface{}) *SMTPClient_Auth {
	c_call := _m.On("Auth", matchers...)
	return &SMTPClient_Auth{Call: c_call}
}

// Auth provides a mock function with given fields: a
func (_m *SMTPClient) Auth(a smtp.Auth) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(smtp.Auth) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SMTPClient_Close struct {
	*mock.Call
}

func (_m SMTPClient_Close) Return(_a0 error) *SMTPClient_Close {
	return &SMTPClient_Close{Call: _m.Call.Return(_a0)}
}

func (_m *SMTPClient) OnClose() *SMTPClient_Close {
	c_call := _m.On("Close")
	return &SMTPClient_Close{Call: c_call}
}

func (_m *SMTPClient) OnCloseMatch(matchers ...interface{}) *SMTPClient_Close {
	c_call := _m.On("Close", matchers...)
	return &SMTPClient_Close{Call: c_call}
}

// Close provides a mock function with given fields:
func (_m *SMTPClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SMTPClient_Data struct {
	*mock.Call
}

func (_m SMTPClient_Data) Return(_a0 io.WriteCloser, _a1 error) *SMTPClient_Data {
	return &SMTPClient_Data{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SMTPClient) OnData() *SMTPClient_Data {
	c_call := _m.On("Data")
	return &SMTPClient_Data{Call: c_call}
}

func (_m *SMTPClient) OnDataMatch(matchers ...interface{}) *SMTPClient_Data {
	c_call := _m.On("Data", matchers...)
	return &SMTPClient_Data{Call: c_call}
}

// Data provides a mock function with given fields:
func (_m *SMTPClient) Data() (io.WriteCloser, error) {
	ret := _m.Called()

	var r0 io.WriteCloser
	if rf, ok := ret.Get(0).(func() io.WriteCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.WriteCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type SMTPClient_Extension struct {
	*mock.Call
}

func (_m SMTPClient_Extension) Return(_a0 bool, _a1 string) *SMTPClient_Extension {
	return &SMTPClient_Extension{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SMTPClient) OnExtension(ext string) *SMTPClient_Extension {
	c_call := _m.On("Extension", ext)
	return &SMTPClient_Extension{Call: c_call}
}

func (_m *SMTPClient) OnExtensionMatch(matchers ...interface{}) *SMTPClient_Extension {
	c_call := _m.On("Extension", matchers...)
	return &SMTPClient_Extension{Call: c_call}
}

// Extension provides a mock function with given fields: ext
func (_m *SMTPClient) Extension(ext string) (bool, string) {
	ret := _m.Called(ext)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(ext)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(ext)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

type SMTPClient_Hello struct {
	*mock.Call
}

func (_m SMTPClient_Hello) Return(_a0 error) *SMTPClient_Hello {
	return &SMTPClient_Hello{Call: _m.Call.Return(_a0)}
}

func (_m *SMTPClient) OnHello(localName string) *SMTPClient_Hello {
	c_call := _m.On("Hello", localName)
	return &SMTPClient_Hello{Call: c_call}
}

func (_m *SMTPClient) OnHelloMatch(matchers ...interface{}) *SMTPClient_Hello {
	c_call := _m.On("Hello", matchers...)
	return &SMTPClient_Hello{Call: c_call}
}

// Hello provides a mock function with given fields: localName
func (_m *SMTPClient) Hello(localName string) error {
	ret := _m.Called(localName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(localName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SMTPClient_Mail struct {
	*mock.Call
}

func (_m SMTPClient_Mail) Return(_a0 error) *SMTPClient_Mail {
	return &SMTPClient_Mail{Call: _m.Call.Return(_a0)}
}

func (_m *SMTPClient) OnMail(from string) *SMTPClient_Mail {
	c_call := _m.On("Mail", from)
	return &SMTPClient_Mail{Call: c_call}
}

func (_m *SMTPClient) OnMailMatch(matchers ...interface{}) *SMTPClient_Mail {
	c_call := _m.On("Mail", matchers...)
	return &SMTPClient_Mail{Call: c_call}
}

// Mail provides a mock function with given fields: from
func (_m *SMTPClient) Mail(from string) error {
	ret := _m.Called(from)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(from)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SMTPClient_Noop struct {
	*mock.Call
}

func (_m SMTPClient_Noop) Return(_a0 error) *SMTPClient_Noop {
	return &SMTPClient_Noop{Call: _m.Call.Return(_a0)}
}

func (_m *SMTPClient) OnNoop() *SMTPClient_Noop {
	c_call := _m.On("Noop")
	return &SMTPClient_Noop{Call: c_call}
}

func (_m *SMTPClient) OnNoopMatch(matchers ...interface{}) *SMTPClient_Noop {
	c_call := _m.On("Noop", matchers...)
	return &SMTPClient_Noop{Call: c_call}
}

// Noop provides a mock function with given fields:
func (_m *SMTPClient) Noop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SMTPClient_Rcpt struct {
	*mock.Call
}

func (_m SMTPClient_Rcpt) Return(_a0 error) *SMTPClient_Rcpt {
	return &SMTPClient_Rcpt{Call: _m.Call.Return(_a0)}
}

func (_m *SMTPClient) OnRcpt(to string) *SMTPClient_Rcpt {
	c_call := _m.On("Rcpt", to)
	return &SMTPClient_Rcpt{Call: c_call}
}

func (_m *SMTPClient) OnRcptMatch(matchers ...interface{}) *SMTPClient_Rcpt {
	c_call := _m.On("Rcpt", matchers...)
	return &SMTPClient_Rcpt{Call: c_call}
}

// Rcpt provides a mock function with given fields: to
func (_m *SMTPClient) Rcpt(to string) error {
	ret := _m.Called(to)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SMTPClient_StartTLS struct {
	*mock.Call
}

func (_m SMTPClient_StartTLS) Return(_a0 error) *SMTPClient_StartTLS {
	return &SMTPClient_StartTLS{Call: _m.Call.Return(_a0)}
}

func (_m *SMTPClient) OnStartTLS(config *tls.Config) *SMTPClient_StartTLS {
	c_call := _m.On("StartTLS", config)
	return &SMTPClient_StartTLS{Call: c_call}
}

func (_m *SMTPClient) OnStartTLSMatch(matchers ...interface{}) *SMTPClient_StartTLS {
	c_call := _m.On("StartTLS", matchers...)
	return &SMTPClient_StartTLS{Call: c_call}
}

// StartTLS provides a mock function with given fields: config
func (_m *SMTPClient) StartTLS(config *tls.Config) error {
	ret := _m.Called(config)

	var r0 error
	if rf, ok := ret.Get(0).(func(*tls.Config) error); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}