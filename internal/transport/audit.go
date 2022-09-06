package transport

import (
	"context"

	"github.com/Viquad/crud-audit-service/pkg/domain"
)

type IAuditService interface {
	Log(context.Context, *domain.LogInput) error
}
