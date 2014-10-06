package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/couchbase/gomemcached"
	"github.com/couchbaselabs/go-couchbase"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// START DEF OMIT
type viewMarker struct {
	Version   int       `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

// Q OMIT
const ddocKey = "/@ddocVersion"
const ddocVersion = 1
const designDoc = `
{
  "views": {
    "likes": {
      "map": "function (doc, meta) { if(doc.type === 'event') { emit(doc.likes, null);} }"
    }
  }
}`

// END DEF OMIT

func main() {
	bucket, err := couchbase.GetBucket("http://localhost:8091/", "default", "demo")
	handleError(err)

	updateDesignDocs(bucket) // HL
}

// R OMIT

func updateDesignDocs(bucket *couchbase.Bucket) {
	// X OMIT
	marker := viewMarker{}
	err := bucket.Get(ddocKey, &marker) // HL
	if err != nil && !gomemcached.IsNotFound(err) {
		handleError(err)
	}
	if marker.Version < ddocVersion { // HL
		fmt.Printf("Installing new version of views (old version=%v)\n",
			marker.Version)
		doc := json.RawMessage([]byte(designDoc))
		err = bucket.PutDDoc("ddoc", &doc) // HL
		handleError(err)
		marker.Version = ddocVersion
		marker.Timestamp = time.Now().UTC()
		marker.Type = "ddocmarker"

		bucket.Set(ddocKey, 0, &marker) // HL
	} else {
		fmt.Printf("Version %v already installed\n", marker.Version)
	}
	// Y OMIT
}
