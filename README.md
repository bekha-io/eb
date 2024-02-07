# eb - Simple Event Bus

## Overview

Event Bus is a Go package that provides a simple event bus implementation for subscribing to and publishing events within your Go applications. It allows components of your application to communicate asynchronously through events. That simple that even kids can understand.

## Installation

To use the Event Bus library in your Go project, you can simply import it using:

```go
import "github.com/bekha-io/eb"
```



And then install it using the `go get` command:

```bash
go get -u github.com/bekha-io/eb
```

## Usage

### Creating an EventBus instance

To create a new instance of the event bus, use the `NewEventBus` function:

```go
eb := eventbus.NewEventBus()
```

### Subscribing to Events

To subscribe to events, use the `Subscribe` method. You can register one or more event handlers for each event type:

```go

var UserLoggedInEvent eb.EventType = "user_logged_in"

eb.Subscribe(UserLoggedInEvent, func(e *eb.Event) {
    // Something's gonna happen here :) ...
})
```

### Publishing Events

To publish an event, create an instance of the `Event` struct and use the `Publish` method:

```go
event := eventbus.NewEvent(eventType)
eb.Publish(event)
```

### Event Options

The `NewEvent` function supports options to customize the event object. You can use `WithData` and `WithHeaders` options to set data and headers respectively:

```go
event := eventbus.NewEvent(eventType, eventbus.WithData(data), eventbus.WithHeaders(headers))
```

## Example

Here's a simple example demonstrating the usage of the Event Bus library:

```go
package main

import (
    "fmt"
    "github.com/bekha-io/eb"
)

func main() {
    eb := eventbus.NewEventBus()

    eb.Subscribe("userLoggedIn", func(event *eventbus.Event) {
        fmt.Println("User logged in")
    })

    event := eventbus.NewEvent("userLoggedIn")
    eb.Publish(event)
}
```
