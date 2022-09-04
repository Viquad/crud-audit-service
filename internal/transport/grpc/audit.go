package grpc

import (
	"context"

	"github.com/Viquad/crud-audit-service/pkg/domain/audit"
)

type IAuditService interface {
	Log(context.Context, *audit.LogRequest) error
}

// AuditServer is a struct which implements audit.AuditServiceServer interface
type AuditServer struct {
	service IAuditService
	audit.UnimplementedAuditServiceServer
}

func NewAuditServer(service IAuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (s *AuditServer) Log(ctx context.Context, request *audit.LogRequest) (*audit.Empty, error) {
	return &audit.Empty{}, s.service.Log(ctx, request)
}
