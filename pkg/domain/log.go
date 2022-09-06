package domain

import "time"

type LogInput struct {
	Action    string    `bson:"action,omitempty"`
	Entity    *string   `bson:"entity,omitempty"`
	EntityId  *int64    `bson:"entity_id,omitempty"`
	UserId    int64     `bson:"user_id,omitempty"`
	Timestamp time.Time `bson:"timestamp,omitempty"`
}
