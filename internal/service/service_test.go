package service

import (
	"context"
	"errors"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jack-hughes/users/pkg/api/users"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jack-hughes/users/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var t = time.Now()
var someErr = errors.New("some-err")

const (
	userID    = "2ef3ee83-f536-4df4-a171-90426dc9199b"
	firstName = "test-first-name"
	lastName  = "test-last-name"
	nickname  = "test-nickname"
	password  = "test-password"
	email     = "test-email"
	country   = "test-country"
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
			_, err := svc.Create(context.TODO(), tt.req)
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
