package main

// START OMIT
func (cp *connectionPool) GetWithTimeout(d time.Duration) (rv *memcached.Client, err error) {

	// short-circuit available connetions
	select {
	case rv, isopen := <-cp.connections:
		if !isopen {
			return nil, errClosedPool
		}
		return rv, nil
	default:
	}

	// END OMIT
	// START P2 OMIT
	// create a very short timer, 1ms
	t := time.NewTimer(ConnPoolAvailWaitTime)
	defer t.Stop()

	select {
	case rv, isopen := <-cp.connections:
		// connection became available
		if !isopen {
			return nil, errClosedPool
		}
		return rv, nil
	case <-t.C:
		// waited 1ms
	}
}

// END P2 OMIT

func f() {
	select {
	case rv, isopen := <-cp.connections:
		// connection became available
		if !isopen {
			return nil, errClosedPool
		}
		return rv, nil
	case <-t.C:
		// START P3 OMIT
		t.Reset(d) // reuse original timer for full timeout
		select {
		case rv, isopen := <-cp.connections:

			// keep trying to get connection from main pool
			if !isopen {
				return nil, errClosedPool
			}
			return rv, nil

		case cp.createsem <- true:

			// create a new connection
			rv, err := cp.mkConn(cp.host, cp.auth)
			if err != nil {
				<-cp.createsem // buffer only allows poolSize + poolOverflow
			}
			return rv, err

		case <-t.C:

			// exceeded caller provided timeout
			return nil, ErrTimeout
			// END P3 OMIT
		}
	}
}
