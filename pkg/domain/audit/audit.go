package audit

import "github.com/Viquad/crud-audit-service/pkg/domain"

func (r *LogRequest) LogInput() *domain.LogInput {
	input := &domain.LogInput{
		Action:    r.Action.String(),
		UserId:    r.UserId,
		Timestamp: r.GetTimestamp().AsTime(),
	}

	if r.Entity != nil {
		v := (*r.Entity).String()
		input.Entity = &v
	}

	if r.EntityId != nil {
		v := *r.EntityId
		input.EntityId = &v
	}

	return input
}
