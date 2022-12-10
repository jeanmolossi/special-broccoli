package event

import (
	"time"

	"github.com/google/uuid"
)

func CreatedAt(createdAt time.Time) EventOption {
	return func(e *Event) {
		e.CreatedAt = int(createdAt.Unix())
	}
}

func Data(data interface{}) EventOption {
	return func(e *Event) {
		e.data = data
	}
}

func ID(id string) EventOption {
	return func(e *Event) {
		e.ID = id
	}
}

func defaults(e *Event) {
	ID(uuid.NewString())(e)
	CreatedAt(time.Now())(e)
}
