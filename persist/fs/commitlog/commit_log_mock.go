// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db/persist/fs/commitlog/types.go

package commitlog

import (
	"reflect"
	"time"

	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/persist/fs"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3x/context"
	"github.com/m3db/m3x/instrument"
	"github.com/m3db/m3x/pool"
	time0 "github.com/m3db/m3x/time"

	"github.com/golang/mock/gomock"
)

// MockCommitLog is a mock of CommitLog interface
type MockCommitLog struct {
	ctrl     *gomock.Controller
	recorder *MockCommitLogMockRecorder
}

// MockCommitLogMockRecorder is the mock recorder for MockCommitLog
type MockCommitLogMockRecorder struct {
	mock *MockCommitLog
}

// NewMockCommitLog creates a new mock instance
func NewMockCommitLog(ctrl *gomock.Controller) *MockCommitLog {
	mock := &MockCommitLog{ctrl: ctrl}
	mock.recorder = &MockCommitLogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCommitLog) EXPECT() *MockCommitLogMockRecorder {
	return _m.recorder
}

// Open mocks base method
func (_m *MockCommitLog) Open() error {
	ret := _m.ctrl.Call(_m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (_mr *MockCommitLogMockRecorder) Open() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Open", reflect.TypeOf((*MockCommitLog)(nil).Open))
}

// Write mocks base method
func (_m *MockCommitLog) Write(ctx context.Context, series Series, datapoint ts.Datapoint, unit time0.Unit, annotation ts.Annotation) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, series, datapoint, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (_mr *MockCommitLogMockRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Write", reflect.TypeOf((*MockCommitLog)(nil).Write), arg0, arg1, arg2, arg3, arg4)
}

// Close mocks base method
func (_m *MockCommitLog) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockCommitLogMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockCommitLog)(nil).Close))
}

// MockIterator is a mock of Iterator interface
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return _m.recorder
}

// Next mocks base method
func (_m *MockIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (_mr *MockIteratorMockRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next))
}

// Current mocks base method
func (_m *MockIterator) Current() (Series, ts.Datapoint, time0.Unit, ts.Annotation) {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(Series)
	ret1, _ := ret[1].(ts.Datapoint)
	ret2, _ := ret[2].(time0.Unit)
	ret3, _ := ret[3].(ts.Annotation)
	return ret0, ret1, ret2, ret3
}

// Current indicates an expected call of Current
func (_mr *MockIteratorMockRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Current", reflect.TypeOf((*MockIterator)(nil).Current))
}

// Err mocks base method
func (_m *MockIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (_mr *MockIteratorMockRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Err", reflect.TypeOf((*MockIterator)(nil).Err))
}

// Close mocks base method
func (_m *MockIterator) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockIteratorMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockIterator)(nil).Close))
}

// MockOptions is a mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsMockRecorder
}

// MockOptionsMockRecorder is the mock recorder for MockOptions
type MockOptionsMockRecorder struct {
	mock *MockOptions
}

// NewMockOptions creates a new mock instance
func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &MockOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockOptions) EXPECT() *MockOptionsMockRecorder {
	return _m.recorder
}

// Validate mocks base method
func (_m *MockOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (_mr *MockOptionsMockRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Validate", reflect.TypeOf((*MockOptions)(nil).Validate))
}

// SetClockOptions mocks base method
func (_m *MockOptions) SetClockOptions(value clock.Options) Options {
	ret := _m.ctrl.Call(_m, "SetClockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetClockOptions indicates an expected call of SetClockOptions
func (_mr *MockOptionsMockRecorder) SetClockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetClockOptions", reflect.TypeOf((*MockOptions)(nil).SetClockOptions), arg0)
}

// ClockOptions mocks base method
func (_m *MockOptions) ClockOptions() clock.Options {
	ret := _m.ctrl.Call(_m, "ClockOptions")
	ret0, _ := ret[0].(clock.Options)
	return ret0
}

// ClockOptions indicates an expected call of ClockOptions
func (_mr *MockOptionsMockRecorder) ClockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ClockOptions", reflect.TypeOf((*MockOptions)(nil).ClockOptions))
}

// SetInstrumentOptions mocks base method
func (_m *MockOptions) SetInstrumentOptions(value instrument.Options) Options {
	ret := _m.ctrl.Call(_m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetInstrumentOptions indicates an expected call of SetInstrumentOptions
func (_mr *MockOptionsMockRecorder) SetInstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetInstrumentOptions", reflect.TypeOf((*MockOptions)(nil).SetInstrumentOptions), arg0)
}

// InstrumentOptions mocks base method
func (_m *MockOptions) InstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

// InstrumentOptions indicates an expected call of InstrumentOptions
func (_mr *MockOptionsMockRecorder) InstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstrumentOptions", reflect.TypeOf((*MockOptions)(nil).InstrumentOptions))
}

// SetRetentionPeriod mocks base method
func (_m *MockOptions) SetRetentionPeriod(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetRetentionPeriod", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetRetentionPeriod indicates an expected call of SetRetentionPeriod
func (_mr *MockOptionsMockRecorder) SetRetentionPeriod(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRetentionPeriod", reflect.TypeOf((*MockOptions)(nil).SetRetentionPeriod), arg0)
}

// RetentionPeriod mocks base method
func (_m *MockOptions) RetentionPeriod() time.Duration {
	ret := _m.ctrl.Call(_m, "RetentionPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RetentionPeriod indicates an expected call of RetentionPeriod
func (_mr *MockOptionsMockRecorder) RetentionPeriod() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RetentionPeriod", reflect.TypeOf((*MockOptions)(nil).RetentionPeriod))
}

// SetBlockSize mocks base method
func (_m *MockOptions) SetBlockSize(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetBlockSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockSize indicates an expected call of SetBlockSize
func (_mr *MockOptionsMockRecorder) SetBlockSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBlockSize", reflect.TypeOf((*MockOptions)(nil).SetBlockSize), arg0)
}

// BlockSize mocks base method
func (_m *MockOptions) BlockSize() time.Duration {
	ret := _m.ctrl.Call(_m, "BlockSize")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockSize indicates an expected call of BlockSize
func (_mr *MockOptionsMockRecorder) BlockSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BlockSize", reflect.TypeOf((*MockOptions)(nil).BlockSize))
}

// SetFilesystemOptions mocks base method
func (_m *MockOptions) SetFilesystemOptions(value fs.Options) Options {
	ret := _m.ctrl.Call(_m, "SetFilesystemOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetFilesystemOptions indicates an expected call of SetFilesystemOptions
func (_mr *MockOptionsMockRecorder) SetFilesystemOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetFilesystemOptions", reflect.TypeOf((*MockOptions)(nil).SetFilesystemOptions), arg0)
}

// FilesystemOptions mocks base method
func (_m *MockOptions) FilesystemOptions() fs.Options {
	ret := _m.ctrl.Call(_m, "FilesystemOptions")
	ret0, _ := ret[0].(fs.Options)
	return ret0
}

// FilesystemOptions indicates an expected call of FilesystemOptions
func (_mr *MockOptionsMockRecorder) FilesystemOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FilesystemOptions", reflect.TypeOf((*MockOptions)(nil).FilesystemOptions))
}

// SetFlushSize mocks base method
func (_m *MockOptions) SetFlushSize(value int) Options {
	ret := _m.ctrl.Call(_m, "SetFlushSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetFlushSize indicates an expected call of SetFlushSize
func (_mr *MockOptionsMockRecorder) SetFlushSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetFlushSize", reflect.TypeOf((*MockOptions)(nil).SetFlushSize), arg0)
}

// FlushSize mocks base method
func (_m *MockOptions) FlushSize() int {
	ret := _m.ctrl.Call(_m, "FlushSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// FlushSize indicates an expected call of FlushSize
func (_mr *MockOptionsMockRecorder) FlushSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FlushSize", reflect.TypeOf((*MockOptions)(nil).FlushSize))
}

// SetStrategy mocks base method
func (_m *MockOptions) SetStrategy(value Strategy) Options {
	ret := _m.ctrl.Call(_m, "SetStrategy", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetStrategy indicates an expected call of SetStrategy
func (_mr *MockOptionsMockRecorder) SetStrategy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetStrategy", reflect.TypeOf((*MockOptions)(nil).SetStrategy), arg0)
}

// Strategy mocks base method
func (_m *MockOptions) Strategy() Strategy {
	ret := _m.ctrl.Call(_m, "Strategy")
	ret0, _ := ret[0].(Strategy)
	return ret0
}

// Strategy indicates an expected call of Strategy
func (_mr *MockOptionsMockRecorder) Strategy() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Strategy", reflect.TypeOf((*MockOptions)(nil).Strategy))
}

// SetFlushInterval mocks base method
func (_m *MockOptions) SetFlushInterval(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetFlushInterval", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetFlushInterval indicates an expected call of SetFlushInterval
func (_mr *MockOptionsMockRecorder) SetFlushInterval(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetFlushInterval", reflect.TypeOf((*MockOptions)(nil).SetFlushInterval), arg0)
}

// FlushInterval mocks base method
func (_m *MockOptions) FlushInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "FlushInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// FlushInterval indicates an expected call of FlushInterval
func (_mr *MockOptionsMockRecorder) FlushInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "FlushInterval", reflect.TypeOf((*MockOptions)(nil).FlushInterval))
}

// SetBacklogQueueSize mocks base method
func (_m *MockOptions) SetBacklogQueueSize(value int) Options {
	ret := _m.ctrl.Call(_m, "SetBacklogQueueSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBacklogQueueSize indicates an expected call of SetBacklogQueueSize
func (_mr *MockOptionsMockRecorder) SetBacklogQueueSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBacklogQueueSize", reflect.TypeOf((*MockOptions)(nil).SetBacklogQueueSize), arg0)
}

// BacklogQueueSize mocks base method
func (_m *MockOptions) BacklogQueueSize() int {
	ret := _m.ctrl.Call(_m, "BacklogQueueSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// BacklogQueueSize indicates an expected call of BacklogQueueSize
func (_mr *MockOptionsMockRecorder) BacklogQueueSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BacklogQueueSize", reflect.TypeOf((*MockOptions)(nil).BacklogQueueSize))
}

// SetBytesPool mocks base method
func (_m *MockOptions) SetBytesPool(value pool.CheckedBytesPool) Options {
	ret := _m.ctrl.Call(_m, "SetBytesPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBytesPool indicates an expected call of SetBytesPool
func (_mr *MockOptionsMockRecorder) SetBytesPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBytesPool", reflect.TypeOf((*MockOptions)(nil).SetBytesPool), arg0)
}

// BytesPool mocks base method
func (_m *MockOptions) BytesPool() pool.CheckedBytesPool {
	ret := _m.ctrl.Call(_m, "BytesPool")
	ret0, _ := ret[0].(pool.CheckedBytesPool)
	return ret0
}

// BytesPool indicates an expected call of BytesPool
func (_mr *MockOptionsMockRecorder) BytesPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BytesPool", reflect.TypeOf((*MockOptions)(nil).BytesPool))
}

// SetReadConcurrency mocks base method
func (_m *MockOptions) SetReadConcurrency(concurrency int) Options {
	ret := _m.ctrl.Call(_m, "SetReadConcurrency", concurrency)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetReadConcurrency indicates an expected call of SetReadConcurrency
func (_mr *MockOptionsMockRecorder) SetReadConcurrency(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetReadConcurrency", reflect.TypeOf((*MockOptions)(nil).SetReadConcurrency), arg0)
}

// ReadConcurrency mocks base method
func (_m *MockOptions) ReadConcurrency() int {
	ret := _m.ctrl.Call(_m, "ReadConcurrency")
	ret0, _ := ret[0].(int)
	return ret0
}

// ReadConcurrency indicates an expected call of ReadConcurrency
func (_mr *MockOptionsMockRecorder) ReadConcurrency() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadConcurrency", reflect.TypeOf((*MockOptions)(nil).ReadConcurrency))
}