// Automatically generated by MockGen. DO NOT EDIT!
// Source: utils_test.go

package tusd_test

import (
	io "io"

	gomock "github.com/golang/mock/gomock"
	tusd "github.com/nagae-memooff/tusd"
)

// Mock of FullDataStore interface
type MockFullDataStore struct {
	ctrl     *gomock.Controller
	recorder *_MockFullDataStoreRecorder
}

// Recorder for MockFullDataStore (not exported)
type _MockFullDataStoreRecorder struct {
	mock *MockFullDataStore
}

func NewMockFullDataStore(ctrl *gomock.Controller) *MockFullDataStore {
	mock := &MockFullDataStore{ctrl: ctrl}
	mock.recorder = &_MockFullDataStoreRecorder{mock}
	return mock
}

func (_m *MockFullDataStore) EXPECT() *_MockFullDataStoreRecorder {
	return _m.recorder
}

func (_m *MockFullDataStore) NewUpload(info tusd.FileInfo) (string, error) {
	ret := _m.ctrl.Call(_m, "NewUpload", info)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFullDataStoreRecorder) NewUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewUpload", arg0)
}

func (_m *MockFullDataStore) WriteChunk(id string, offset int64, src io.Reader) (int64, error) {
	ret := _m.ctrl.Call(_m, "WriteChunk", id, offset, src)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFullDataStoreRecorder) WriteChunk(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteChunk", arg0, arg1, arg2)
}

func (_m *MockFullDataStore) GetInfo(id string) (tusd.FileInfo, error) {
	ret := _m.ctrl.Call(_m, "GetInfo", id)
	ret0, _ := ret[0].(tusd.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFullDataStoreRecorder) GetInfo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInfo", arg0)
}

func (_m *MockFullDataStore) Terminate(id string) error {
	ret := _m.ctrl.Call(_m, "Terminate", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFullDataStoreRecorder) Terminate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Terminate", arg0)
}

func (_m *MockFullDataStore) ConcatUploads(destination string, partialUploads []string) error {
	ret := _m.ctrl.Call(_m, "ConcatUploads", destination, partialUploads)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFullDataStoreRecorder) ConcatUploads(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConcatUploads", arg0, arg1)
}

func (_m *MockFullDataStore) GetReader(id string) (io.Reader, error) {
	ret := _m.ctrl.Call(_m, "GetReader", id)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFullDataStoreRecorder) GetReader(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetReader", arg0)
}

func (_m *MockFullDataStore) FinishUpload(id string) error {
	ret := _m.ctrl.Call(_m, "FinishUpload", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFullDataStoreRecorder) FinishUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FinishUpload", arg0)
}

// Mock of Locker interface
type MockLocker struct {
	ctrl     *gomock.Controller
	recorder *_MockLockerRecorder
}

// Recorder for MockLocker (not exported)
type _MockLockerRecorder struct {
	mock *MockLocker
}

func NewMockLocker(ctrl *gomock.Controller) *MockLocker {
	mock := &MockLocker{ctrl: ctrl}
	mock.recorder = &_MockLockerRecorder{mock}
	return mock
}

func (_m *MockLocker) EXPECT() *_MockLockerRecorder {
	return _m.recorder
}

func (_m *MockLocker) LockUpload(id string) error {
	ret := _m.ctrl.Call(_m, "LockUpload", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLockerRecorder) LockUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LockUpload", arg0)
}

func (_m *MockLocker) UnlockUpload(id string) error {
	ret := _m.ctrl.Call(_m, "UnlockUpload", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLockerRecorder) UnlockUpload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnlockUpload", arg0)
}
