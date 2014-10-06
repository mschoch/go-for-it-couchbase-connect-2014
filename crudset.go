package main

import (
	"encoding/json"
	"fmt"

	"github.com/couchbaselabs/go-couchbase"
)

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

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	bucket, err := couchbase.GetBucket("http://localhost:8091/", "default", "demo") // HL
	handleError(err)

	// START OMIT
	event := NewEvent("Couchbase Connect")
	err = bucket.Set("cc2014", 0, event) // HL
	handleError(err)

	event = NewEvent("GopherCon India")
	err = bucket.Set("gci2015", 0, event) // HL
	handleError(err)

	fmt.Printf("Saved Events\n")
	// END OMIT
}
