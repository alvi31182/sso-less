package main

import (
	"context"
	"sso-less/internal/config"

	log "github.com/alvi31182/log-go"
)

func main() {
	logger, err := log.NewLogger(log.LogLevelDebug, &log.ServiceInfo{
		Name:    "sso-service",
		Version: "1.0.0",
	})
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	log.SetDefaultLogger(logger)

	cfg := config.MustLoad()

	log.Infow(context.Background(), "Starting application",
		"env", cfg.Env,
		"storage_path", cfg.StoragePath,
		"token_ttl", cfg.TokenTtl,
		"grpc_port", cfg.GRPC.Port,
		"grpc_timeout", cfg.GRPC.Timeout,
	)
}
