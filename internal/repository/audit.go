package repository

import (
	"context"
	"log"

	"github.com/Viquad/crud-audit-service/pkg/domain/audit"
)

type AuditRepository struct {
	// db
}

func NewAuditRepository() *AuditRepository {
	return &AuditRepository{}
}

func (s *AuditRepository) Log(ctx context.Context, request *audit.LogInput) error {
	log.Println(request)
	return nil
}
