package main

import (
	"fmt"
)

func doSomething() error {
	return nil
}

// START OMIT
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := doSomething()
	handleError(err)
	fmt.Printf("Success")
}

// END OMIT
