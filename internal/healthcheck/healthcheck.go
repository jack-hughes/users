package healthcheck

import (
	"context"
	"github.com/jack-hughes/users/internal/storage"
	"github.com/jack-hughes/users/pkg/api/health"
	"go.uber.org/zap"
)

type HealthChecker struct {
	health.HealthServer

	log *zap.Logger
	db  storage.Postgres
}

func NewHealthChecker(log *zap.Logger, db storage.Postgres) *HealthChecker {
	return &HealthChecker{
		log: log.With(zap.String("component", "health")),
		db:  db,
	}
}

func (h HealthChecker) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	h.log.Debug("performing health check")

	if err := h.db.Ping(ctx); err != nil {
		return &health.HealthCheckResponse{
			Status: health.HealthCheckResponse_NOT_SERVING,
		}, nil
	}

	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthChecker) Watch(req *health.HealthCheckRequest, server health.Health_WatchServer) error {
	h.log.Debug("performing watch health check")
	if err := h.db.Ping(context.TODO()); err != nil {
		return server.Send(&health.HealthCheckResponse{
			Status: health.HealthCheckResponse_NOT_SERVING,
		})
	}

	return server.Send(&health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	})
}
