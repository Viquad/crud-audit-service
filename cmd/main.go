package main

import (
	"github.com/Viquad/crud-audit-service/internal/repository"
	"github.com/Viquad/crud-audit-service/internal/service"
	"github.com/Viquad/crud-audit-service/internal/transport/grpc"
	_grpc "google.golang.org/grpc"
)

func main() {
	audit_repo := repository.NewAuditRepository()
	audit_service := service.NewAuditService(audit_repo)
	audit_server := grpc.NewAuditServer(audit_service)
	grpc_server := _grpc.NewServer()
	server := grpc.NewServer(audit_server, grpc_server)

	server.ListenAndServe(9000)
}
