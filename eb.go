// Package eb provides an event bus implementation for subscribing to and publishing events.
package eb

import (
	"sync"
)

// EventBus manages subscriptions and publishes events.
type EventBus struct {
	rw       *sync.RWMutex
	resolver map[EventType][]Handler // Map to store event handlers for each event type
}

// NewEventBus creates and returns a new instance of EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		rw:       &sync.RWMutex{},
		resolver: make(map[EventType][]Handler),
	}
}

// Subscribe adds a new event handler for the specified event type.
func (eb *EventBus) Subscribe(event EventType, handler Handler) {
	eb.rw.Lock()
	defer eb.rw.Unlock()
	eb.resolver[event] = append(eb.resolver[event], handler)
}

// Publish sends the specified event to all subscribed handlers.
func (eb *EventBus) Publish(event *Event) {
	eb.rw.Lock()
	defer eb.rw.Unlock()

	if handlers, ok := eb.resolver[event.Type]; ok {
		for _, handler := range handlers {
			go handler(event)
		}
	}
}
