package main

import (
	"context"
	"log"

	"github.com/Viquad/crud-audit-service/internal/repository"
	"github.com/Viquad/crud-audit-service/internal/service"
	"github.com/Viquad/crud-audit-service/internal/transport/grpc"
	"github.com/Viquad/crud-audit-service/pkg/config"
	"github.com/Viquad/crud-audit-service/pkg/database"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New("configs", "config")
	if err != nil {
		log.Fatalf("config: %s", err.Error())
	}

	db, err := database.NewMongoConnection(ctx, cfg.DB)
	if err != nil {
		log.Fatalf("db: %s", err.Error())
	}

	audit_repo := repository.NewAuditRepository(db)
	audit_service := service.NewAuditService(audit_repo)
	audit_server := grpc.NewAuditServer(audit_service)
	server := grpc.NewServer(audit_server)

	server.ListenAndServe(9000)
}
