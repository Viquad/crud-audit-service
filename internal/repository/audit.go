package repository

import (
	"context"

	"github.com/Viquad/crud-audit-service/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuditRepository struct {
	db *mongo.Database
}

func NewAuditRepository(db *mongo.Database) *AuditRepository {
	return &AuditRepository{db}
}

func (r *AuditRepository) Log(ctx context.Context, input *domain.LogInput) error {
	_, err := r.db.Collection("logs").InsertOne(ctx, input)
	return err
}
