package utils

import (
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_SanitiseError(t *testing.T) {
	type tt struct {
		name     string
		error    error
		expected error
	}
	tests := []tt{
		{
			name:     "no error",
			error:    nil,
			expected: nil,
		},
		{
			name:     "unique constraint violation error",
			error:    &pgconn.PgError{Code: pgerrcode.UniqueViolation, Severity: "test-sev", Message: "test-message"},
			expected: status.Error(codes.AlreadyExists, "test-sev: test-message (SQLSTATE 23505)"),
		},
		{
			name:     "default fallthrough",
			error:    &pgconn.PgError{Code: pgerrcode.IOError, Severity: "test-sev", Message: "test-message"},
			expected: status.Error(codes.Internal, "test-sev: test-message (SQLSTATE 58030)"),
		},
		{
			name:     "non-pg error",
			error:    errors.New("bang"),
			expected: status.Error(codes.Internal, "bang"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SanitiseError(tt.error)
			if tt.expected != nil {
				require.Error(t, err)
				assert.Equal(t, tt.expected, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
