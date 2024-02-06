package eb

// EventType represents the type of an event.
type EventType string

// EventOpt is a function type that represents an option for configuring an Event.
type EventOpt func(*Event)

// Type of the event
// Headers associated with the event
// Data payload of the event
type Event struct {
	Type    EventType
	Headers map[string]string
	Data    interface{}
}

// WithData is a function that returns an EventOpt function which sets the data field of an Event struct.
// The data parameter is the value to be set for the data field.
func WithData(data interface{}) EventOpt {
	return func(e *Event) {
		e.Data = data
	}
}

// WithHeaders is a function that returns an EventOpt function which sets the headers of an Event.
// The headers parameter is a map of string keys to string values representing the headers to be set.
// The returned EventOpt function sets the headers of the given Event to the provided headers.
// Example usage:
//   event := NewEvent(WithHeaders(map[string]string{"Content-Type": "application/json"}))
func WithHeaders(headers map[string]string) EventOpt {
	return func(e *Event) {
		e.Headers = headers
	}
}

// NewEvent creates a new event with the specified type and optional parameters.
// It returns a pointer to the created event.
func NewEvent(t EventType, opts ...EventOpt) *Event {
	e := &Event{
		Type: t,
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}
