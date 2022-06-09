package storage

import (
	"context"
	"fmt"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

var listScan = listScanFunc

//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=./storage.go -package=mocks -destination=../../test/mocks/postgres_mocks.go
type Postgres interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type Store struct {
	db  Postgres
	log *zap.Logger
}

func New(log *zap.Logger, db Postgres) Store {
	log = log.With(zap.String("component", "storage"))
	log.Debug("configuring database")

	return Store{
		db:  db,
		log: log,
	}
}

func (s Store) Create(ctx context.Context, usr types.User) (types.User, error) {
	sql := `
INSERT INTO users.users (first_name, last_name, nickname, password, email, country)
VALUES ($1, $2, $3, $4, $5, $6)
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

func (s Store) Update(ctx context.Context, usr types.User) (types.User, error) {
	sql := `
UPDATE users.users SET
    first_name = $1,
    last_name = $2,
    nickname = $3,
    password = $4,
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

func (s Store) Delete(ctx context.Context, usr types.User) error {
	sql := "DELETE FROM users.users WHERE id = $1;"
	_, err := s.db.Exec(ctx, sql, usr.Id)
	if err != nil {
		return err
	}

	return nil
}

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
