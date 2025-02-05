// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/pkg/proto/control_plane (interfaces: ChainRenewalServiceServer,TrustMaterialServiceServer)

// Package mock_control_plane is a generated GoMock package.
package mock_control_plane

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	control_plane "github.com/scionproto/scion/pkg/proto/control_plane"
)

// MockChainRenewalServiceServer is a mock of ChainRenewalServiceServer interface.
type MockChainRenewalServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockChainRenewalServiceServerMockRecorder
}

// MockChainRenewalServiceServerMockRecorder is the mock recorder for MockChainRenewalServiceServer.
type MockChainRenewalServiceServerMockRecorder struct {
	mock *MockChainRenewalServiceServer
}

// NewMockChainRenewalServiceServer creates a new mock instance.
func NewMockChainRenewalServiceServer(ctrl *gomock.Controller) *MockChainRenewalServiceServer {
	mock := &MockChainRenewalServiceServer{ctrl: ctrl}
	mock.recorder = &MockChainRenewalServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainRenewalServiceServer) EXPECT() *MockChainRenewalServiceServerMockRecorder {
	return m.recorder
}

// ChainRenewal mocks base method.
func (m *MockChainRenewalServiceServer) ChainRenewal(arg0 context.Context, arg1 *control_plane.ChainRenewalRequest) (*control_plane.ChainRenewalResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainRenewal", arg0, arg1)
	ret0, _ := ret[0].(*control_plane.ChainRenewalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainRenewal indicates an expected call of ChainRenewal.
func (mr *MockChainRenewalServiceServerMockRecorder) ChainRenewal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainRenewal", reflect.TypeOf((*MockChainRenewalServiceServer)(nil).ChainRenewal), arg0, arg1)
}

// MockTrustMaterialServiceServer is a mock of TrustMaterialServiceServer interface.
type MockTrustMaterialServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTrustMaterialServiceServerMockRecorder
}

// MockTrustMaterialServiceServerMockRecorder is the mock recorder for MockTrustMaterialServiceServer.
type MockTrustMaterialServiceServerMockRecorder struct {
	mock *MockTrustMaterialServiceServer
}

// NewMockTrustMaterialServiceServer creates a new mock instance.
func NewMockTrustMaterialServiceServer(ctrl *gomock.Controller) *MockTrustMaterialServiceServer {
	mock := &MockTrustMaterialServiceServer{ctrl: ctrl}
	mock.recorder = &MockTrustMaterialServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrustMaterialServiceServer) EXPECT() *MockTrustMaterialServiceServerMockRecorder {
	return m.recorder
}

// Chains mocks base method.
func (m *MockTrustMaterialServiceServer) Chains(arg0 context.Context, arg1 *control_plane.ChainsRequest) (*control_plane.ChainsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chains", arg0, arg1)
	ret0, _ := ret[0].(*control_plane.ChainsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chains indicates an expected call of Chains.
func (mr *MockTrustMaterialServiceServerMockRecorder) Chains(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chains", reflect.TypeOf((*MockTrustMaterialServiceServer)(nil).Chains), arg0, arg1)
}

// TRC mocks base method.
func (m *MockTrustMaterialServiceServer) TRC(arg0 context.Context, arg1 *control_plane.TRCRequest) (*control_plane.TRCResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TRC", arg0, arg1)
	ret0, _ := ret[0].(*control_plane.TRCResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TRC indicates an expected call of TRC.
func (mr *MockTrustMaterialServiceServerMockRecorder) TRC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TRC", reflect.TypeOf((*MockTrustMaterialServiceServer)(nil).TRC), arg0, arg1)
}
