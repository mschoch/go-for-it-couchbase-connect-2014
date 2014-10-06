package main

import (
	"encoding/json"
)

// START OMIT
type Event struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Likes int    `json:"likes"`
}

func NewEvent(name string) *Event {
	return &Event{"event", name, 0}
}

func NewEventJSON(jsonbytes []byte) (event *Event) {
	err := json.Unmarshal(jsonbytes, &event)
	handleError(err)
	return
}

func (e *Event) String() string {
	return fmt.Sprintf("Event '%s', Likes: %d", e.Name, e.Likes)
}

// END OMIT
