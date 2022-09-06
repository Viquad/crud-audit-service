package service

import (
	"context"

	"github.com/Viquad/crud-audit-service/pkg/domain"
)

type IAuditRepository interface {
	Log(context.Context, *domain.LogInput) error
}

type AuditService struct {
	repo IAuditRepository
}

func NewAuditService(repo IAuditRepository) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

func (s *AuditService) Log(ctx context.Context, input *domain.LogInput) error {
	return s.repo.Log(ctx, input)
}
