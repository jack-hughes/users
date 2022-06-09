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
	// connStringFmt is the postgres database connection string
	connStringFmt = "postgresql://%s:%s@%s/%s"
)

// NotFoundError is a bespoke error for db executions where the command tag affects 0 rows
type NotFoundError struct{}

// Error returns the default not found error
func (n NotFoundError) Error() string {
	return "record not found"
}

// NewDatabase returns a new database connection pool
func NewDatabase(ctx context.Context, username, password, hostname, port, db string) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(connStringFmt, username, password, net.JoinHostPort(hostname, port), db)
	return pgxpool.Connect(ctx, connString)
}

// SanitiseError enables us to return standardised gRPC status codes for specific errors
func SanitiseError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return status.Error(codes.AlreadyExists, err.Error())
		default:
			return status.Error(codes.Internal, err.Error())
		}
	}

	if errors.As(err, &NotFoundError{}) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
