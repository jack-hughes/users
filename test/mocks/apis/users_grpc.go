// Code generated by MockGen. DO NOT EDIT.
// Source: ../../pkg/api/users/users_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	users "github.com/jack-hughes/users/pkg/api/users"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockUsersClient is a mock of UsersClient interface.
type MockUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsersClientMockRecorder
}

// MockUsersClientMockRecorder is the mock recorder for MockUsersClient.
type MockUsersClientMockRecorder struct {
	mock *MockUsersClient
}

// NewMockUsersClient creates a new mock instance.
func NewMockUsersClient(ctrl *gomock.Controller) *MockUsersClient {
	mock := &MockUsersClient{ctrl: ctrl}
	mock.recorder = &MockUsersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersClient) EXPECT() *MockUsersClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsersClient) Create(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*users.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersClient)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockUsersClient) Delete(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUsersClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsersClient)(nil).Delete), varargs...)
}

// List mocks base method.
func (m *MockUsersClient) List(ctx context.Context, in *users.ListUsersRequest, opts ...grpc.CallOption) (users.Users_ListClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(users.Users_ListClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUsersClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsersClient)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockUsersClient) Update(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*users.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUsersClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsersClient)(nil).Update), varargs...)
}

// MockUsers_ListClient is a mock of Users_ListClient interface.
type MockUsers_ListClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsers_ListClientMockRecorder
}

// MockUsers_ListClientMockRecorder is the mock recorder for MockUsers_ListClient.
type MockUsers_ListClientMockRecorder struct {
	mock *MockUsers_ListClient
}

// NewMockUsers_ListClient creates a new mock instance.
func NewMockUsers_ListClient(ctrl *gomock.Controller) *MockUsers_ListClient {
	mock := &MockUsers_ListClient{ctrl: ctrl}
	mock.recorder = &MockUsers_ListClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers_ListClient) EXPECT() *MockUsers_ListClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockUsers_ListClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockUsers_ListClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockUsers_ListClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockUsers_ListClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUsers_ListClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUsers_ListClient)(nil).Context))
}

// Header mocks base method.
func (m *MockUsers_ListClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockUsers_ListClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockUsers_ListClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockUsers_ListClient) Recv() (*users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockUsers_ListClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockUsers_ListClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockUsers_ListClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUsers_ListClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUsers_ListClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockUsers_ListClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUsers_ListClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUsers_ListClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockUsers_ListClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockUsers_ListClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockUsers_ListClient)(nil).Trailer))
}

// MockUsersServer is a mock of UsersServer interface.
type MockUsersServer struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServerMockRecorder
}

// MockUsersServerMockRecorder is the mock recorder for MockUsersServer.
type MockUsersServerMockRecorder struct {
	mock *MockUsersServer
}

// NewMockUsersServer creates a new mock instance.
func NewMockUsersServer(ctrl *gomock.Controller) *MockUsersServer {
	mock := &MockUsersServer{ctrl: ctrl}
	mock.recorder = &MockUsersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersServer) EXPECT() *MockUsersServerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsersServer) Create(arg0 context.Context, arg1 *users.User) (*users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockUsersServer) Delete(arg0 context.Context, arg1 *users.User) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUsersServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsersServer)(nil).Delete), arg0, arg1)
}

// List mocks base method.
func (m *MockUsersServer) List(arg0 *users.ListUsersRequest, arg1 users.Users_ListServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockUsersServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsersServer)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockUsersServer) Update(arg0 context.Context, arg1 *users.User) (*users.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*users.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUsersServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsersServer)(nil).Update), arg0, arg1)
}

// mustEmbedUnimplementedUsersServer mocks base method.
func (m *MockUsersServer) mustEmbedUnimplementedUsersServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUsersServer")
}

// mustEmbedUnimplementedUsersServer indicates an expected call of mustEmbedUnimplementedUsersServer.
func (mr *MockUsersServerMockRecorder) mustEmbedUnimplementedUsersServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUsersServer", reflect.TypeOf((*MockUsersServer)(nil).mustEmbedUnimplementedUsersServer))
}

// MockUnsafeUsersServer is a mock of UnsafeUsersServer interface.
type MockUnsafeUsersServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUsersServerMockRecorder
}

// MockUnsafeUsersServerMockRecorder is the mock recorder for MockUnsafeUsersServer.
type MockUnsafeUsersServerMockRecorder struct {
	mock *MockUnsafeUsersServer
}

// NewMockUnsafeUsersServer creates a new mock instance.
func NewMockUnsafeUsersServer(ctrl *gomock.Controller) *MockUnsafeUsersServer {
	mock := &MockUnsafeUsersServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUsersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUsersServer) EXPECT() *MockUnsafeUsersServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUsersServer mocks base method.
func (m *MockUnsafeUsersServer) mustEmbedUnimplementedUsersServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUsersServer")
}

// mustEmbedUnimplementedUsersServer indicates an expected call of mustEmbedUnimplementedUsersServer.
func (mr *MockUnsafeUsersServerMockRecorder) mustEmbedUnimplementedUsersServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUsersServer", reflect.TypeOf((*MockUnsafeUsersServer)(nil).mustEmbedUnimplementedUsersServer))
}

// MockUsers_ListServer is a mock of Users_ListServer interface.
type MockUsers_ListServer struct {
	ctrl     *gomock.Controller
	recorder *MockUsers_ListServerMockRecorder
}

// MockUsers_ListServerMockRecorder is the mock recorder for MockUsers_ListServer.
type MockUsers_ListServerMockRecorder struct {
	mock *MockUsers_ListServer
}

// NewMockUsers_ListServer creates a new mock instance.
func NewMockUsers_ListServer(ctrl *gomock.Controller) *MockUsers_ListServer {
	mock := &MockUsers_ListServer{ctrl: ctrl}
	mock.recorder = &MockUsers_ListServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers_ListServer) EXPECT() *MockUsers_ListServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockUsers_ListServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUsers_ListServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUsers_ListServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockUsers_ListServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUsers_ListServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUsers_ListServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockUsers_ListServer) Send(arg0 *users.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockUsers_ListServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockUsers_ListServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockUsers_ListServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockUsers_ListServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockUsers_ListServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockUsers_ListServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUsers_ListServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUsers_ListServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockUsers_ListServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockUsers_ListServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockUsers_ListServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockUsers_ListServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockUsers_ListServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockUsers_ListServer)(nil).SetTrailer), arg0)
}
