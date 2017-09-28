package events

import "time"

type Conditioner func() bool
type Handler func(event *Event)
type EventListener struct {
	event       *Event
	conditioner Conditioner
	handler     Handler
}

func NewEventListener() *EventListener {
	return &EventListener{}
}

// event_listener.On(fn(){}).Trigger("SomeEvent").HandledBy(fn(){})

func (el *EventListener) On(conditioner Conditioner) *EventListener {
	el.conditioner = conditioner
	return el
}

func (el *EventListener) Trigger(eventName string) *EventListener {
	el.event = NewEvent(eventName)

	return el
}

func (el *EventListener) HandledBy(handler Handler) {
	el.handler = handler

	if el.conditioner != nil && el.event != nil {
		go func(conditioner Conditioner, handler Handler) {
			for {
				if conditioner() {
					handler(el.event)
				}

				time.Sleep(time.Microsecond)
			}
		}(el.conditioner, el.handler)

		return
	}

	panic("Can't handle! Conditioner and event name must be set!")
}
