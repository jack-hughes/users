package storage

import (
	"context"
	"github.com/jack-hughes/users/pkg/api/users"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type Storage interface {
	Create(ctx context.Context, req *users.User) (User, error)
	Update(ctx context.Context, req *users.User) (User, error)
	Delete(ctx context.Context, req *users.User) error
	List(ctx context.Context, req *users.ListUsersRequest) ([]User, error)
}

type Postgres interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type Store struct {
	db  Postgres
	log *zap.Logger
}

func New(log *zap.Logger, db Postgres) Storage {
	log = log.With(zap.String("component", "storage"))
	log.Debug("configuring database")

	return &Store{
		db:  db,
		log: log,
	}
}

func (s *Store) Create(ctx context.Context, usr *users.User) (User, error) {
	sql := `
INSERT INTO users.users (first_name, lastname, nickname, password, email, country)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
`
	var dbUsr User
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
		return User{}, err
	}

	return dbUsr, nil
}

func (s *Store) Update(ctx context.Context, req *users.User) (User, error) {
	return User{
		Id: req.Id,
	}, nil
}

func (s *Store) Delete(ctx context.Context, req *users.User) error {
	return nil
}

func (s *Store) List(ctx context.Context, req *users.ListUsersRequest) ([]User, error) {
	return []User{
		{
			Id: req.Id,
		},
		{
			Id: req.Id,
		},
	}, nil
}

func usrToArgs(in *users.User) []interface{} {
	return []interface{}{
		in.FirstName,
		in.LastName,
		in.Nickname,
		in.Password,
		in.Email,
		in.Country,
	}
}
