// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/owodunni/hano-scraper/internal/server (interfaces: Logger,Metrics,Processor)

// Package server is a generated GoMock package.
package server

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/owodunni/hano-scraper/internal/models"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Error mocks base method.
func (m *MockLogger) Error(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", arg0)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), arg0)
}

// Infof mocks base method.
func (m *MockLogger) Infof(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// InflightRequestsGaugeAdd mocks base method.
func (m *MockMetrics) InflightRequestsGaugeAdd(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InflightRequestsGaugeAdd", arg0)
}

// InflightRequestsGaugeAdd indicates an expected call of InflightRequestsGaugeAdd.
func (mr *MockMetricsMockRecorder) InflightRequestsGaugeAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InflightRequestsGaugeAdd", reflect.TypeOf((*MockMetrics)(nil).InflightRequestsGaugeAdd), arg0)
}

// RequestCountInc mocks base method.
func (m *MockMetrics) RequestCountInc(arg0 string, arg1 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequestCountInc", arg0, arg1)
}

// RequestCountInc indicates an expected call of RequestCountInc.
func (mr *MockMetricsMockRecorder) RequestCountInc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCountInc", reflect.TypeOf((*MockMetrics)(nil).RequestCountInc), arg0, arg1)
}

// ResponseBytesCountAdd mocks base method.
func (m *MockMetrics) ResponseBytesCountAdd(arg0 string, arg1, arg2 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResponseBytesCountAdd", arg0, arg1, arg2)
}

// ResponseBytesCountAdd indicates an expected call of ResponseBytesCountAdd.
func (mr *MockMetricsMockRecorder) ResponseBytesCountAdd(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseBytesCountAdd", reflect.TypeOf((*MockMetrics)(nil).ResponseBytesCountAdd), arg0, arg1, arg2)
}

// ResponseTimeHistogramObserve mocks base method.
func (m *MockMetrics) ResponseTimeHistogramObserve(arg0 string, arg1 int, arg2 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResponseTimeHistogramObserve", arg0, arg1, arg2)
}

// ResponseTimeHistogramObserve indicates an expected call of ResponseTimeHistogramObserve.
func (mr *MockMetricsMockRecorder) ResponseTimeHistogramObserve(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseTimeHistogramObserve", reflect.TypeOf((*MockMetrics)(nil).ResponseTimeHistogramObserve), arg0, arg1, arg2)
}

// MockProcessor is a mock of Processor interface.
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor.
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance.
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockProcessor) CreateUser(arg0 context.Context, arg1 models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockProcessorMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockProcessor)(nil).CreateUser), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockProcessor) GetUserByID(arg0 context.Context, arg1 uint64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockProcessorMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockProcessor)(nil).GetUserByID), arg0, arg1)
}
