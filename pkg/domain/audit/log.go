package audit

import "time"

type LogInput struct {
	Action    string    `bson:"action,omitempty"`
	Entity    *string   `bson:"entity,omitempty"`
	EntityId  *int64    `bson:"entity_id,omitempty"`
	UserId    int64     `bson:"user_id,omitempty"`
	Timestamp time.Time `bson:"timestamp,omitempty"`
}

func NewLogInput(request *LogRequest) *LogInput {
	input := &LogInput{
		Action:    request.Action.String(),
		UserId:    request.UserId,
		Timestamp: request.GetTimestamp().AsTime(),
	}

	if request.Entity != nil {
		v := (*request.Entity).String()
		input.Entity = &v
	}

	if request.EntityId != nil {
		v := *request.EntityId
		input.EntityId = &v
	}

	return input
}
