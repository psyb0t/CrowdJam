package events

type Event struct {
	Name string
}

func NewEvent(name string) *Event {
	return &Event{Name: name}
}
