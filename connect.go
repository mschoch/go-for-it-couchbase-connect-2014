package main

// START OMIT
import (
	"fmt"

	"github.com/couchbaselabs/go-couchbase" // HL
)

func main() {
	bucket, err := couchbase.GetBucket("http://localhost:8091/", "default", "demo") // HL
	handleError(err)
	defer bucket.Close()
	fmt.Printf("Connected to Couchbase Bucket '%s'\n", bucket.Name)
}

// END OMIT

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
