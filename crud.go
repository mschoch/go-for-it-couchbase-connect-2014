package main

import (
	"fmt" // OMIT
	// OMIT
	"github.com/couchbaselabs/go-couchbase"
)

func main() {

	bucket, err := couchbase.GetBucket("http://localhost:8091/", "default", "demo") // HL
	if err != nil {
		fmt.Println(err)
		return
	}

	// START SET OMIT
	sampleDoc := map[string]interface{}{
		"name": "Couchbase Connect",
	}

	err = bucket.Set("key", 0, sampleDoc) // HL
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Created Document: 'key'\n")
	// END SET OMIT

	var doc map[string]interface{}
	err = bucket.Get("key", &doc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Read Document: 'key'\n")
	// END OMIT
}
