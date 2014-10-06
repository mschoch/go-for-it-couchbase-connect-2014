package main

func main() {
	var val interface{}

	// START OMIT

	// write to channel
	channel <- val

	// read from channel
	val = <-channel

	//END OMIT

}
