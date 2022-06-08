package storage

import (
	"context"
	"github.com/jack-hughes/users/internal/storage/types"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=./storage.go -package=mocks -destination=../../test/mocks/postgres_mocks.go
type Postgres interface {
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
INSERT INTO users.users (first_name, lastname, nickname, password, email, country)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
`
	var dbUsr types.User
	err := s.db.QueryRow(ctx, sql, usrToArgs(usr)...).Scan(
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

func (s Store) Update(ctx context.Context, req types.User) (types.User, error) {
	return types.User{
		Id: req.Id,
	}, nil
}

func (s Store) Delete(ctx context.Context, req types.User) error {
	return nil
}

func (s Store) List(ctx context.Context, countryFilter string) ([]types.User, error) {
	return []types.User{
		{
			Id: "fake-id",
		},
		{
			Id: "fake-id",
		},
	}, nil
}

func usrToArgs(in types.User) []interface{} {
	return []interface{}{
		in.FirstName,
		in.LastName,
		in.Nickname,
		in.Password,
		in.Email,
		in.Country,
	}
}
