package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

const (
	connStringFmt = "postgresql://%s:%s@%s/%s"
)

// NewDatabase returns a new database connection pool.
func NewDatabase(ctx context.Context, username, password, hostname, port, db string) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(connStringFmt, username, password, net.JoinHostPort(hostname, port), db)
	return pgxpool.Connect(ctx, connString)
}

func SanitiseError(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return status.Error(codes.AlreadyExists, err.Error())
		default:
			return status.Error(codes.Internal, err.Error())
		}
	}

	return status.Error(codes.Internal, err.Error())
}
