package storage

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jack-hughes/users/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"time"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen --build_flags=--mod=mod -package=mocks -destination=../../test/mocks/rows_mocks.go github.com/jackc/pgx/v4 Row,Rows

const (
	userID    = "2ef3ee83-f536-4df4-a171-90426dc9199b"
	firstName = "test-first-name"
	lastName  = "test-last-name"
	nickname  = "test-nickname"
	password  = "test-password"
	email     = "test-email"
	country   = "test-country"
)

var (
	t       = time.Now()
	someErr = errors.New("some-err")
)

func TestStore_New(t *testing.T) {
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

			New(zap.NewNop(), db)
		})
	}
}

func TestStore_Create(t *testing.T) {
	type tt struct {
		name          string
		ctx           context.Context
		request       types.User
		expected      types.User
		err           error
		expectedCalls func(tt tt, ctrl *gomock.Controller, mp *mocks.MockPostgres)
	}
	tests := []tt{
		{
			name:     "successfully create a user",
			ctx:      context.TODO(),
			request:  getDbUser(),
			expected: types.User{},
			err:      nil,
			expectedCalls: func(tt tt, ctrl *gomock.Controller, mp *mocks.MockPostgres) {
				row := mocks.NewMockRow(ctrl)
				mp.
					EXPECT().
					QueryRow(
						gomock.Any(),
						queryCreate(),
						[]interface{}{
							firstName,
							lastName,
							nickname,
							password,
							email,
							country,
						},
					).
					Return(row)
				row.EXPECT().Scan(
					&tt.expected.Id,
					&tt.expected.FirstName,
					&tt.expected.LastName,
					&tt.expected.Nickname,
					&tt.expected.Password,
					&tt.expected.Email,
					&tt.expected.Country,
					&tt.expected.CreatedAt,
					&tt.expected.UpdatedAt,
				).Return(nil)
			},
		},
		{
			name:     "fail to scan rows",
			ctx:      context.TODO(),
			request:  getDbUser(),
			expected: types.User{},
			err:      someErr,
			expectedCalls: func(tt tt, ctrl *gomock.Controller, mp *mocks.MockPostgres) {
				row := mocks.NewMockRow(ctrl)
				mp.
					EXPECT().
					QueryRow(
						gomock.Any(),
						queryCreate(),
						[]interface{}{
							firstName,
							lastName,
							nickname,
							password,
							email,
							country,
						},
					).
					Return(row)
				row.EXPECT().Scan(
					&tt.expected.Id,
					&tt.expected.FirstName,
					&tt.expected.LastName,
					&tt.expected.Nickname,
					&tt.expected.Password,
					&tt.expected.Email,
					&tt.expected.Country,
					&tt.expected.CreatedAt,
					&tt.expected.UpdatedAt,
				).Return(someErr)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mdb := mocks.NewMockPostgres(ctrl)

			if tt.expectedCalls != nil {
				tt.expectedCalls(tt, ctrl, mdb)
			}
			store := Store{
				log: zap.NewNop(),
				db:  mdb,
			}
			ex, err := store.Create(tt.ctx, tt.request)
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, ex)
			}
			ctrl.Finish()
		})
	}
}

func queryCreate() string {
	return `
INSERT INTO users.users (first_name, last_name, nickname, password, email, country)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
`
}

func getDbUser() types.User {
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
