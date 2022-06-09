package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jack-hughes/users/pkg/api/users"
	"github.com/jack-hughes/users/test/mocks"
	umocks "github.com/jack-hughes/users/test/mocks/apis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=../../pkg/api/users/users_grpc.pb.go -package=mocks -destination=../../test/mocks/apis/users_grpc.go

var t = time.Now()
var someErr = status.Error(codes.Internal, "some-err")

const (
	userID        = "2ef3ee83-f536-4df4-a171-90426dc9199b"
	firstName     = "test-first-name"
	lastName      = "test-last-name"
	nickname      = "test-nickname"
	password      = "test-password"
	email         = "test-email"
	country       = "test-country"
	countryFilter = "UK"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "new is called",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mocks.NewMockPostgres(ctrl)
			store := mocks.NewMockStorage(ctrl)

			NewUserService(zap.NewNop(), store, db)
		})
	}
}

func Test_Create(t *testing.T) {
	type tt struct {
		name          string
		req           *users.User
		expected      *users.User
		err           error
		expectedCalls func(tt tt, m *mocks.MockStorage)
	}
	tests := []tt{
		{
			name: "successfully create a user",
			req: &users.User{
				FirstName: firstName,
				LastName:  lastName,
				Nickname:  nickname,
				Password:  password,
				Email:     email,
				Country:   country,
			},
			expected: getAPIUser(),
			expectedCalls: func(tt tt, m *mocks.MockStorage) {
				m.EXPECT().
					Create(
						gomock.Any(),
						types.User{
							FirstName: firstName,
							LastName:  lastName,
							Nickname:  nickname,
							Password:  password,
							Email:     email,
							Country:   country,
						},
					).
					Return(getDBUser(), nil)
			},
			err: nil,
		},
		{
			name: "database returns an error",
			req: &users.User{
				FirstName: firstName,
				LastName:  lastName,
				Nickname:  nickname,
				Password:  password,
				Email:     email,
				Country:   country,
			},
			expected: getAPIUser(),
			expectedCalls: func(tt tt, m *mocks.MockStorage) {
				m.EXPECT().
					Create(
						gomock.Any(),
						types.User{
							FirstName: firstName,
							LastName:  lastName,
							Nickname:  nickname,
							Password:  password,
							Email:     email,
							Country:   country,
						},
					).
					Return(getDBUser(), errors.New("some-err"))
			},
			err: someErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mocks.NewMockPostgres(ctrl)
			store := mocks.NewMockStorage(ctrl)
			svc := UserService{
				log:   zap.NewNop(),
				db:    db,
				store: store,
			}
			if tt.expectedCalls != nil {
				tt.expectedCalls(tt, store)
			}
			user, err := svc.Create(context.TODO(), tt.req)
			if tt.err != nil {
				require.Error(t, err)
				assert.Equal(t, tt.err, err)
			} else {
				assert.Equal(t, tt.expected, user)
			}
			ctrl.Finish()
		})
	}
}

func Test_Update(t *testing.T) {
	type tt struct {
		name          string
		req           *users.User
		expected      *users.User
		err           error
		expectedCalls func(tt tt, m *mocks.MockStorage)
	}
	tests := []tt{
		{
			name: "successfully update a user",
			req: &users.User{
				Id:        userID,
				FirstName: firstName,
				LastName:  lastName,
				Nickname:  nickname,
				Password:  password,
				Email:     email,
				Country:   country,
			},
			expected: getAPIUser(),
			expectedCalls: func(tt tt, m *mocks.MockStorage) {
				m.EXPECT().
					Update(
						gomock.Any(),
						types.User{
							Id:        userID,
							FirstName: firstName,
							LastName:  lastName,
							Nickname:  nickname,
							Password:  password,
							Email:     email,
							Country:   country,
						},
					).
					Return(getDBUser(), nil)
			},
			err: nil,
		},
		{
			name: "database returns an error",
			req: &users.User{
				Id:        userID,
				FirstName: firstName,
				LastName:  lastName,
				Nickname:  nickname,
				Password:  password,
				Email:     email,
				Country:   country,
			},
			expected: getAPIUser(),
			expectedCalls: func(tt tt, m *mocks.MockStorage) {
				m.EXPECT().
					Update(
						gomock.Any(),
						types.User{
							Id:        userID,
							FirstName: firstName,
							LastName:  lastName,
							Nickname:  nickname,
							Password:  password,
							Email:     email,
							Country:   country,
						},
					).
					Return(getDBUser(), errors.New("some-err"))
			},
			err: someErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mocks.NewMockPostgres(ctrl)
			store := mocks.NewMockStorage(ctrl)
			svc := UserService{
				log:   zap.NewNop(),
				db:    db,
				store: store,
			}
			if tt.expectedCalls != nil {
				tt.expectedCalls(tt, store)
			}
			user, err := svc.Update(context.TODO(), tt.req)
			if tt.err != nil {
				require.Error(t, err)
				assert.Equal(t, tt.err, err)
			} else {
				assert.Equal(t, tt.expected, user)
			}
			ctrl.Finish()
		})
	}
}

func Test_Delete(t *testing.T) {
	type tt struct {
		name          string
		req           *users.User
		expected      *empty.Empty
		err           error
		expectedCalls func(tt tt, m *mocks.MockStorage)
	}
	tests := []tt{
		{
			name: "successfully delete a user",
			req: &users.User{
				Id: userID,
			},
			expected: &empty.Empty{},
			expectedCalls: func(tt tt, m *mocks.MockStorage) {
				m.EXPECT().
					Delete(
						gomock.Any(),
						types.User{Id: userID},
					).
					Return(nil)
			},
			err: nil,
		},
		{
			name: "database returns an error",
			req: &users.User{
				Id: userID,
			},
			expected: &empty.Empty{},
			expectedCalls: func(tt tt, m *mocks.MockStorage) {
				m.EXPECT().
					Delete(
						gomock.Any(),
						types.User{Id: userID},
					).
					Return(errors.New("some-err"))
			},
			err: someErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mocks.NewMockPostgres(ctrl)
			store := mocks.NewMockStorage(ctrl)
			svc := UserService{
				log:   zap.NewNop(),
				db:    db,
				store: store,
			}
			if tt.expectedCalls != nil {
				tt.expectedCalls(tt, store)
			}
			user, err := svc.Delete(context.TODO(), tt.req)
			if tt.err != nil {
				require.Error(t, err)
				assert.Equal(t, tt.err, err)
			} else {
				assert.Equal(t, tt.expected, user)
			}
			ctrl.Finish()
		})
	}
}

func Test_List(t *testing.T) {
	type table struct {
		name     string
		ctx      context.Context
		expCalls func(tt table, m *mocks.MockStorage, src *umocks.MockUsers_ListServer)
		req      *users.ListUsersRequest
		err      error
	}
	tests := []table{
		{
			name: "successfully list users by country code",
			ctx:  context.TODO(),
			req:  &users.ListUsersRequest{Filter: countryFilter},
			expCalls: func(tt table, m *mocks.MockStorage, src *umocks.MockUsers_ListServer) {
				src.
					EXPECT().
					Context().
					Return(tt.ctx)
				m.
					EXPECT().
					List(tt.ctx, countryFilter).
					Return(getDBUserList(), nil)
				src.
					EXPECT().
					Send(getAPIUser()).Times(1)
			},
			err: nil,
		},
		{
			name: "fail to retrieve user list from db",
			ctx:  context.TODO(),
			req:  &users.ListUsersRequest{Filter: countryFilter},
			expCalls: func(tt table, m *mocks.MockStorage, src *umocks.MockUsers_ListServer) {
				src.
					EXPECT().
					Context().
					Return(tt.ctx)
				m.
					EXPECT().
					List(tt.ctx, countryFilter).
					Return(getDBUserList(), errors.New("some-err"))
			},
			err: someErr,
		},
		{
			name: "fail to send user to stream",
			ctx:  context.TODO(),
			req:  &users.ListUsersRequest{Filter: countryFilter},
			expCalls: func(tt table, m *mocks.MockStorage, src *umocks.MockUsers_ListServer) {
				src.
					EXPECT().
					Context().
					Return(tt.ctx)
				m.
					EXPECT().
					List(tt.ctx, countryFilter).
					Return(getDBUserList(), nil)
				src.
					EXPECT().
					Send(getAPIUser()).Return(errors.New("some-err"))
			},
			err: someErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			store := mocks.NewMockStorage(ctrl)
			srv := umocks.NewMockUsers_ListServer(ctrl)

			svc := UserService{
				log:   zap.NewNop(),
				store: store,
			}
			if tt.expCalls != nil {
				tt.expCalls(tt, store, srv)
			}

			err := svc.List(tt.req, srv)
			if tt.err != nil {
				require.Error(t, err)
				assert.Equal(t, tt.err, err)
			}

			ctrl.Finish()
		})
	}
}

func getAPIUser() *users.User {
	return &users.User{
		Id:        userID,
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Password:  password,
		Email:     email,
		Country:   country,
		CreatedAt: timestamppb.New(t),
		UpdatedAt: timestamppb.New(t),
	}
}

func getDBUser() types.User {
	return types.User{
		Id:        userID,
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Password:  password,
		Email:     email,
		Country:   country,
		CreatedAt: t,
		UpdatedAt: t,
	}
}

func getDBUserList() []types.User {
	return []types.User{
		{
			Id:        userID,
			FirstName: firstName,
			LastName:  lastName,
			Nickname:  nickname,
			Password:  password,
			Email:     email,
			Country:   country,
			CreatedAt: t,
			UpdatedAt: t,
		},
	}
}
