package storage

import (
	"context"
	"fmt"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jack-hughes/users/internal/utils"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

// listScan is defined in order to mock in unit tests
var listScan = listScanFunc

// Postgres defines operations available on the database connection pool, the interface allows us to easily mock
// database functionality for testing
//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=./storage.go -package=mocks -destination=../../test/mocks/postgres_mocks.go
type Postgres interface {
	Ping(ctx context.Context) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

// Store contains a database connection pool and logger
type Store struct {
	db  Postgres
	log *zap.Logger
}

// New instantiates a new storage object with a database connection pool and logger
func New(log *zap.Logger, db Postgres) Store {
	log = log.With(zap.String("component", "storage"))
	log.Debug("configuring database")

	return Store{
		db:  db,
		log: log,
	}
}

// Create a user in Postgres, returning inserted fields
func (s Store) Create(ctx context.Context, usr types.User) (types.User, error) {
	sql := `
INSERT INTO users.users (first_name, last_name, nickname, password, email, country)
VALUES ($1, $2, $3, crypt($4, gen_salt('bf', 8)), $5, $6)
RETURNING *;
`
	var dbUsr types.User
	err := s.db.QueryRow(ctx, sql, insertUserArgs(usr)...).Scan(
		&dbUsr.Id,
		&dbUsr.FirstName,
		&dbUsr.LastName,
		&dbUsr.Nickname,
		&dbUsr.Password,
		&dbUsr.Email,
		&dbUsr.Country,
		&dbUsr.CreatedAt,
		&dbUsr.UpdatedAt,
	)
	if err != nil {
		return types.User{}, err
	}

	return dbUsr, nil
}

// Update a user in Postgres based on their id, returning all user columns
func (s Store) Update(ctx context.Context, usr types.User) (types.User, error) {
	sql := `
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
	var dbUsr types.User
	err := s.db.QueryRow(ctx, sql, updateUserArgs(usr)...).Scan(
		&dbUsr.Id,
		&dbUsr.FirstName,
		&dbUsr.LastName,
		&dbUsr.Nickname,
		&dbUsr.Password,
		&dbUsr.Email,
		&dbUsr.Country,
		&dbUsr.CreatedAt,
		&dbUsr.UpdatedAt,
	)
	if err != nil {
		return types.User{}, err
	}

	return dbUsr, nil
}

// Delete a user in Postgres
func (s Store) Delete(ctx context.Context, usr types.User) error {
	sql := "DELETE FROM users.users WHERE id = $1;"
	tag, err := s.db.Exec(ctx, sql, usr.Id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return utils.NotFoundError{}
	}

	return nil
}

// List all users in Postgres where a country filter applies
func (s Store) List(ctx context.Context, countryFilter string) ([]types.User, error) {
	sql := `
SELECT id, first_name, last_name, nickname, password, email, country, created_at, updated_at
FROM users.users
WHERE country = $1;
`
	rows, err := s.db.Query(ctx, sql, countryFilter)
	if err != nil {
		return nil, err
	}

	res, err := listScan(rows)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// insertUserArgs generates an interface slice for prettier args
func insertUserArgs(in types.User) []interface{} {
	return []interface{}{
		in.FirstName,
		in.LastName,
		in.Nickname,
		in.Password,
		in.Email,
		in.Country,
	}
}

// updateUserArgs generates an interface slice for prettier args
func updateUserArgs(in types.User) []interface{} {
	return []interface{}{
		in.FirstName,
		in.LastName,
		in.Nickname,
		in.Password,
		in.Email,
		in.Country,
		in.Id,
	}
}

// listScanFunc scans rows into the user slice, it is its own function to enable easier testing
func listScanFunc(rows pgx.Rows) ([]types.User, error) {
	var list []types.User
	for rows.Next() {
		var u types.User
		if err := rows.Scan(
			&u.Id,
			&u.FirstName,
			&u.LastName,
			&u.Nickname,
			&u.Password,
			&u.Email,
			&u.Country,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scanning database row into struct: %v", err)
		}
		list = append(list, u)
	}

	return list, nil
}
