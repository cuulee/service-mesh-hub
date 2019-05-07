// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/external/prometheus/v1/prometheus_config_client.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1 "github.com/solo-io/supergloo/pkg/api/external/prometheus/v1"
)

// MockPrometheusConfigClient is a mock of PrometheusConfigClient interface
type MockPrometheusConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockPrometheusConfigClientMockRecorder
}

// MockPrometheusConfigClientMockRecorder is the mock recorder for MockPrometheusConfigClient
type MockPrometheusConfigClientMockRecorder struct {
	mock *MockPrometheusConfigClient
}

// NewMockPrometheusConfigClient creates a new mock instance
func NewMockPrometheusConfigClient(ctrl *gomock.Controller) *MockPrometheusConfigClient {
	mock := &MockPrometheusConfigClient{ctrl: ctrl}
	mock.recorder = &MockPrometheusConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrometheusConfigClient) EXPECT() *MockPrometheusConfigClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockPrometheusConfigClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockPrometheusConfigClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockPrometheusConfigClient)(nil).BaseClient))
}

// Register mocks base method
func (m *MockPrometheusConfigClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockPrometheusConfigClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPrometheusConfigClient)(nil).Register))
}

// Read mocks base method
func (m *MockPrometheusConfigClient) Read(namespace, name string, opts clients.ReadOpts) (*v1.PrometheusConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", namespace, name, opts)
	ret0, _ := ret[0].(*v1.PrometheusConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockPrometheusConfigClientMockRecorder) Read(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockPrometheusConfigClient)(nil).Read), namespace, name, opts)
}

// Write mocks base method
func (m *MockPrometheusConfigClient) Write(resource *v1.PrometheusConfig, opts clients.WriteOpts) (*v1.PrometheusConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", resource, opts)
	ret0, _ := ret[0].(*v1.PrometheusConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockPrometheusConfigClientMockRecorder) Write(resource, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockPrometheusConfigClient)(nil).Write), resource, opts)
}

// Delete mocks base method
func (m *MockPrometheusConfigClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", namespace, name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPrometheusConfigClientMockRecorder) Delete(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPrometheusConfigClient)(nil).Delete), namespace, name, opts)
}

// List mocks base method
func (m *MockPrometheusConfigClient) List(namespace string, opts clients.ListOpts) (v1.PrometheusConfigList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", namespace, opts)
	ret0, _ := ret[0].(v1.PrometheusConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPrometheusConfigClientMockRecorder) List(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPrometheusConfigClient)(nil).List), namespace, opts)
}

// Watch mocks base method
func (m *MockPrometheusConfigClient) Watch(namespace string, opts clients.WatchOpts) (<-chan v1.PrometheusConfigList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1.PrometheusConfigList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockPrometheusConfigClientMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockPrometheusConfigClient)(nil).Watch), namespace, opts)
}