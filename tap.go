package main

import (
	"fmt"
	"time"

	"github.com/couchbase/gomemcached/client"
	"github.com/couchbaselabs/go-couchbase"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	bucket, err := couchbase.GetBucket("http://localhost:8091/", "default", "demo") // HL
	handleError(err)

	// START OMIT
	args := memcached.DefaultTapArguments() // HL
	feed, err := bucket.StartTapFeed(&args) // HL
	handleError(err)

	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 5; i++ {
			bucket.SetRaw(fmt.Sprintf("tap-%d", i), 0, []byte("x"))
		}
	}()

	fmt.Printf("Listening to TAP:\n")
	for op := range feed.C {
		fmt.Printf("Received %s\n", op.String())
		if len(op.Value) > 0 && len(op.Value) < 500 {
			fmt.Printf("\tValue: %s\n", op.Value)
		}
	}
	// END OMIT
}
