package main

import (
	"encoding/json"
	"fmt"
	"sync"

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

// LIKE OMIT
func likeEvent(bucket *couchbase.Bucket, id string) {
	bucket.Update(id, 0, func(current []byte) ([]byte, error) { // HL
		event := NewEventJSON(current)
		event.Likes++
		return json.Marshal(event)
	})
}

// LIKE OMIT

func main() {

	bucket, err := couchbase.GetBucket("http://localhost:8091/", "default", "demo") // HL
	handleError(err)

	// START OMIT
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			likeEvent(bucket, "cc2014")
		}()
	}
	wg.Wait()

	var event Event
	err = bucket.Get("cc2014", &event)
	handleError(err)
	fmt.Println(&event)
	// END OMIT
}
