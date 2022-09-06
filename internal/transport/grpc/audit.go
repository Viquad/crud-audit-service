package grpc

import (
	"context"

	"github.com/Viquad/crud-audit-service/internal/transport"
	"github.com/Viquad/crud-audit-service/pkg/domain/audit"
)

// AuditServer is a struct which implements audit.AuditServiceServer interface
type AuditServer struct {
	audit.UnimplementedAuditServiceServer

	service transport.IAuditService
}

func NewAuditServer(service transport.IAuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (s *AuditServer) Log(ctx context.Context, request *audit.LogRequest) (*audit.Empty, error) {
	input := request.LogInput()
	return &audit.Empty{}, s.service.Log(ctx, input)
}
