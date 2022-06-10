package main

import (
	"context"
	"fmt"
	"github.com/jack-hughes/users/cmd/server/config"
	"github.com/jack-hughes/users/internal/healthcheck"
	"github.com/jack-hughes/users/internal/logger"
	"github.com/jack-hughes/users/internal/service"
	"github.com/jack-hughes/users/internal/storage"
	"github.com/jack-hughes/users/internal/utils"
	"github.com/jack-hughes/users/pkg/api/health"
	"github.com/jack-hughes/users/pkg/api/users"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	log := logger.NewZapLogger("users-service", zap.DebugLevel)
	log.Debug("initialising users-service...")

	// Load app configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("couldn't load config: %v", err))
	}
	log.Debug("config loaded")
	// Establish database connection pool
	db, err := utils.NewDatabase(
		context.TODO(),
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("couldn't initialise database pool: %v", err))
	}
	if err := db.Ping(context.TODO()); err != nil {
		log.Fatal(fmt.Sprintf("couldn't ping database: %v", err))
	}
	log.Debug("database successfully pinged")

	// Bootstrap and launch service
	store := storage.New(log, db)
	addr := fmt.Sprintf(":%s", cfg.GRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(fmt.Sprintf("couldn't load config: %v", err))
	}
	log.Debug(fmt.Sprintf("listening on port: %v", cfg.GRPCPort))

	srv := grpc.NewServer()
	// Register services
	svc := service.NewUserService(log, store, db)
	hc := healthcheck.NewHealthChecker(log, db)
	users.RegisterUsersServer(srv, svc)
	health.RegisterHealthServer(srv, hc)
	if err := srv.Serve(lis); err != nil {
		log.Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}
