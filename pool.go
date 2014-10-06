package main

// START OMIT
type connectionPool struct {
	host        string
	mkConn      func(host string, ah AuthHandler) (*memcached.Client, error) // OMIT
	auth        AuthHandler
	connections chan *memcached.Client
	createsem   chan bool
}

func newConnectionPool(host string, ah AuthHandler, poolSize, poolOverflow int) *connectionPool {
	return &connectionPool{
		host:        host,
		connections: make(chan *memcached.Client, poolSize), // HL
		createsem:   make(chan bool, poolSize+poolOverflow), // HL
		mkConn:      defaultMkConn,                          // OMIT
		auth:        ah,
	}
}

// END OMIT
