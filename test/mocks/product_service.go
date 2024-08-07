// Code generated by MockGen. DO NOT EDIT.
// Source: domain/enterprise/services/product_service.go
//
// Generated by this command:
//
//	mockgen -destination=domain/application/mocks/product_service.go -source=domain/enterprise/services/product_service.go -package=application_mocks
//

// Package application_mocks is a generated GoMock package.
package application_mocks

import (
	reflect "reflect"

	entities "github.com/luizeduu/hexagonal-arch-go/domain/enterprise/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockProductServiceInterface is a mock of ProductServiceInterface interface.
type MockProductServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceInterfaceMockRecorder
}

// MockProductServiceInterfaceMockRecorder is the mock recorder for MockProductServiceInterface.
type MockProductServiceInterfaceMockRecorder struct {
	mock *MockProductServiceInterface
}

// NewMockProductServiceInterface creates a new mock instance.
func NewMockProductServiceInterface(ctrl *gomock.Controller) *MockProductServiceInterface {
	mock := &MockProductServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProductServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceInterface) EXPECT() *MockProductServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductServiceInterface) Create(product entities.ProductInterface) (entities.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", product)
	ret0, _ := ret[0].(entities.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceInterfaceMockRecorder) Create(product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductServiceInterface)(nil).Create), product)
}

// Disable mocks base method.
func (m *MockProductServiceInterface) Disable(product entities.ProductInterface) (entities.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", product)
	ret0, _ := ret[0].(entities.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disable indicates an expected call of Disable.
func (mr *MockProductServiceInterfaceMockRecorder) Disable(product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockProductServiceInterface)(nil).Disable), product)
}

// Enable mocks base method.
func (m *MockProductServiceInterface) Enable(product entities.ProductInterface) (entities.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", product)
	ret0, _ := ret[0].(entities.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enable indicates an expected call of Enable.
func (mr *MockProductServiceInterfaceMockRecorder) Enable(product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockProductServiceInterface)(nil).Enable), product)
}

// Get mocks base method.
func (m *MockProductServiceInterface) Get(id string) (entities.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(entities.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductServiceInterfaceMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductServiceInterface)(nil).Get), id)
}

// MockProductReader is a mock of ProductReader interface.
type MockProductReader struct {
	ctrl     *gomock.Controller
	recorder *MockProductReaderMockRecorder
}

// MockProductReaderMockRecorder is the mock recorder for MockProductReader.
type MockProductReaderMockRecorder struct {
	mock *MockProductReader
}

// NewMockProductReader creates a new mock instance.
func NewMockProductReader(ctrl *gomock.Controller) *MockProductReader {
	mock := &MockProductReader{ctrl: ctrl}
	mock.recorder = &MockProductReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductReader) EXPECT() *MockProductReaderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProductReader) Get(id string) (entities.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(entities.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductReaderMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductReader)(nil).Get), id)
}

// MockProductWriter is a mock of ProductWriter interface.
type MockProductWriter struct {
	ctrl     *gomock.Controller
	recorder *MockProductWriterMockRecorder
}

// MockProductWriterMockRecorder is the mock recorder for MockProductWriter.
type MockProductWriterMockRecorder struct {
	mock *MockProductWriter
}

// NewMockProductWriter creates a new mock instance.
func NewMockProductWriter(ctrl *gomock.Controller) *MockProductWriter {
	mock := &MockProductWriter{ctrl: ctrl}
	mock.recorder = &MockProductWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductWriter) EXPECT() *MockProductWriterMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockProductWriter) Save(product entities.ProductInterface) (entities.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(entities.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockProductWriterMockRecorder) Save(product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductWriter)(nil).Save), product)
}

// MockProductPersitenceInterface is a mock of ProductPersitenceInterface interface.
type MockProductPersitenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductPersitenceInterfaceMockRecorder
}

// MockProductPersitenceInterfaceMockRecorder is the mock recorder for MockProductPersitenceInterface.
type MockProductPersitenceInterfaceMockRecorder struct {
	mock *MockProductPersitenceInterface
}

// NewMockProductPersitenceInterface creates a new mock instance.
func NewMockProductPersitenceInterface(ctrl *gomock.Controller) *MockProductPersitenceInterface {
	mock := &MockProductPersitenceInterface{ctrl: ctrl}
	mock.recorder = &MockProductPersitenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductPersitenceInterface) EXPECT() *MockProductPersitenceInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProductPersitenceInterface) Get(id string) (entities.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(entities.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductPersitenceInterfaceMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductPersitenceInterface)(nil).Get), id)
}

// Save mocks base method.
func (m *MockProductPersitenceInterface) Save(product entities.ProductInterface) (entities.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(entities.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockProductPersitenceInterfaceMockRecorder) Save(product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductPersitenceInterface)(nil).Save), product)
}
