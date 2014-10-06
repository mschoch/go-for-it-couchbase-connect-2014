package main

import (
	"fmt"

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
	args := map[string]interface{}{
		"stale":      false,
		"descending": true,
	}

	res, err := bucket.View("ddoc", "likes", args) // HL
	handleError(err)

	for _, r := range res.Rows {
		fmt.Printf("Key: %v - DocID: '%s'\n", r.Key, r.ID)
	}
	// END OMIT
}
