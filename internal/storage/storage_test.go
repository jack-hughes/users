package storage

import (
	"context"
	"errors"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jack-hughes/users/internal/utils"
	"github.com/jack-hughes/users/test/mocks"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"time"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen --build_flags=--mod=mod -package=mocks -destination=../../test/mocks/rows_mocks.go github.com/jackc/pgx/v4 Row,Rows

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

var (
	tm      = time.Now()
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

func TestStore_Update(t *testing.T) {
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
			name:     "successfully update a user",
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
						queryUpdate(),
						[]interface{}{
							firstName,
							lastName,
							nickname,
							password,
							email,
							country,
							userID,
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
						queryUpdate(),
						[]interface{}{
							firstName,
							lastName,
							nickname,
							password,
							email,
							country,
							userID,
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
			ex, err := store.Update(tt.ctx, tt.request)
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

func TestStore_Delete(t *testing.T) {
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
			name:     "successfully update a user",
			ctx:      context.TODO(),
			request:  getDbUser(),
			expected: types.User{},
			err:      nil,
			expectedCalls: func(tt tt, ctrl *gomock.Controller, mp *mocks.MockPostgres) {
				mp.
					EXPECT().
					Exec(
						gomock.Any(),
						queryDelete(),
						userID,
					).
					Return(nil, nil)
			},
		},
		{
			name:     "error on user delete",
			ctx:      context.TODO(),
			request:  getDbUser(),
			expected: types.User{},
			err:      someErr,
			expectedCalls: func(tt tt, ctrl *gomock.Controller, mp *mocks.MockPostgres) {
				mp.
					EXPECT().
					Exec(
						gomock.Any(),
						queryDelete(),
						userID,
					).
					Return(nil, someErr)
			},
		},
		{
			name:     "no rows affected",
			ctx:      context.TODO(),
			request:  getDbUser(),
			expected: types.User{},
			err:      utils.NotFoundError{},
			expectedCalls: func(tt tt, ctrl *gomock.Controller, mp *mocks.MockPostgres) {
				mp.
					EXPECT().
					Exec(
						gomock.Any(),
						queryDelete(),
						userID,
					).
					Return(pgconn.CommandTag{}, utils.NotFoundError{})
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
			err := store.Delete(tt.ctx, tt.request)
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			}
			ctrl.Finish()
		})
	}
}

func TestStore_List(t *testing.T) {
	type tt struct {
		name     string
		ctx      context.Context
		scanFunc func(rows pgx.Rows) ([]types.User, error)
		res      []types.User
		expCalls func(tt tt, db *mocks.MockPostgres, mr *mocks.MockRows)
		err      error
	}
	tests := []tt{
		{
			name: "successfully list users",
			ctx:  context.TODO(),
			scanFunc: func(rows pgx.Rows) ([]types.User, error) {
				return getListUsers(), nil
			},
			res: getListUsers(),
			expCalls: func(tt tt, db *mocks.MockPostgres, mr *mocks.MockRows) {
				db.
					EXPECT().
					Query(tt.ctx, queryList(), countryFilter).
					Return(mr, nil)
			},
			err: nil,
		},
		{
			name: "scan err",
			ctx:  context.TODO(),
			scanFunc: func(rows pgx.Rows) ([]types.User, error) {
				return nil, someErr
			},
			res: nil,
			expCalls: func(tt tt, db *mocks.MockPostgres, mr *mocks.MockRows) {
				db.
					EXPECT().
					Query(tt.ctx, queryList(), countryFilter).
					Return(mr, nil)
			},
			err: someErr,
		},
		{
			name: "query err",
			ctx:  context.TODO(),
			res:  nil,
			expCalls: func(tt tt, db *mocks.MockPostgres, mr *mocks.MockRows) {
				db.
					EXPECT().
					Query(tt.ctx, queryList(), countryFilter).
					Return(nil, someErr)
			},
			err: someErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mdb := mocks.NewMockPostgres(ctrl)
			mr := mocks.NewMockRows(ctrl)

			if tt.expCalls != nil {
				tt.expCalls(tt, mdb, mr)
			}

			if tt.scanFunc != nil {
				listScan = tt.scanFunc
			}

			store := Store{
				db:  mdb,
				log: zap.NewNop(),
			}
			ex, err := store.List(tt.ctx, countryFilter)

			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.res, ex)
			}

			ctrl.Finish()
		})
	}
}

func TestStore_ListScanFunc(t *testing.T) {
	type tt struct {
		name string
		rows pgx.Rows
		res  []types.User
		err  error
	}
	tests := []tt{
		{
			name: "successfully scan users",
			rows: pgxpoolmock.NewRows([]string{
				"id",
				"first_name",
				"last_name",
				"nickname",
				"password",
				"email",
				"country",
				"created_at",
				"updated_at",
			}).AddRow(
				userID,
				firstName,
				lastName,
				nickname,
				password,
				email,
				country,
				tm,
				tm,
			).ToPgxRows(),
			res: getListUsers(),
			err: nil,
		},
		{
			name: "fail to scan users",
			rows: pgxpoolmock.NewRows([]string{
				"test",
			}).AddRow(
				123,
			).ToPgxRows(),
			res: getListUsers(),
			err: errors.New("scanning database row into struct: Incorrect argument number 9 for columns 1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			list, err := listScanFunc(tt.rows)

			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.res, list)
			}

			ctrl.Finish()
		})
	}
}

func queryCreate() string {
	return `
INSERT INTO users.users (first_name, last_name, nickname, password, email, country)
VALUES ($1, $2, $3, crypt($4, gen_salt('bf', 8)), $5, $6)
RETURNING *;
`
}

func queryUpdate() string {
	return `
UPDATE users.users SET
    first_name = $1,
    last_name = $2,
    nickname = $3,
    password = crypt($4, gen_salt('bf', 8)),
    email = $5,
    country = $6
    WHERE id = $7
RETURNING *;
`
}

func queryList() string {
	return `
SELECT id, first_name, last_name, nickname, password, email, country, created_at, updated_at
FROM users.users
WHERE country = $1;
`
}

func queryDelete() string {
	return "DELETE FROM users.users WHERE id = $1;"
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
		CreatedAt: tm,
		UpdatedAt: tm,
	}
}
func getListUsers() []types.User {
	return []types.User{
		{
			Id:        userID,
			FirstName: firstName,
			LastName:  lastName,
			Nickname:  nickname,
			Password:  password,
			Email:     email,
			Country:   country,
			CreatedAt: tm,
			UpdatedAt: tm,
		},
	}
}
