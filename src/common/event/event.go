package event

import "reflect"

type (
	Event struct {
		ID        string `event:"hashkey"`
		Name      string
		data      interface{}
		CreatedAt int `event:"sortkey"`
	}

	EventOption func(e *Event)
)

func (e *Event) Data() interface{} {
	valueof := reflect.ValueOf(e.data)
	switch valueof.Kind() {
	case reflect.Ptr:
		return valueof.Elem().Interface()
	default:
		return valueof.Interface()
	}
}

func NewEvent(name string, options ...EventOption) *Event {
	e := &Event{Name: name}

	if len(options) == 0 {
		panic("no options configured")
	}

	defaults(e)
	for _, option := range options {
		option(e)
	}

	return e
}
