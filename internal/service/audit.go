package service

import (
	"context"

	"github.com/Viquad/crud-audit-service/pkg/domain/audit"
)

type IAuditRepository interface {
	Log(context.Context, *audit.LogInput) error
}

type AuditService struct {
	repo IAuditRepository
}

func NewAuditService(repo IAuditRepository) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

func (s *AuditService) Log(ctx context.Context, request *audit.LogRequest) error {
	input := audit.NewLogInput(request)
	return s.repo.Log(ctx, input)
}
