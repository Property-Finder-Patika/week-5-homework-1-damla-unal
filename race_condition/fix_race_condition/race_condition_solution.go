package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	To prevent the race condition we can use a Mutex.
	We can avoid this by locking access to the counter variable when one of the goroutines reads it,
	and then unlocking it when it is done writing the incremented value.
	This way, nobody else can use that variable when a particular goroutine is updating it.
	This is exactly what Mutex does.
*/

const limit2 = 5000

var counter2 = 0 //counter which is what all goroutines will try to modify simultaneously

var mu sync.Mutex

// increment increases the counter variable up to the limit
func incrementWithMutex() {
	mu.Lock()
	fmt.Printf("Counter before increment: %d\n", counter2)
	for i := 0; i < limit2; i++ { // each routine is expected to increase the counter 5k
		counter2 += 1
	}
	fmt.Printf("Counter after increment: %d\n", counter2)
	mu.Unlock()
}

func main() {
	//two goroutine are started
	// When the first call to mutex.Lock() is made, mutex state changes to Locked.
	// Any other calls to mutex.Lock() will block the goroutine until mutex.Unlock() is called
	// So, only one thread can access the critical section.
	go incrementWithMutex()
	go incrementWithMutex()

	time.Sleep(1 * time.Millisecond)
	mu.Lock()
	fmt.Printf("Counter is expected to be %d but found as %d", 2*limit2, counter2)
	mu.Unlock()

	/*
		Expected Output:
			Counter before increment: 0
			Counter after increment: 5000
			Counter before increment: 5000
			Counter after increment: 10000
			Counter is expected to be 10000 but found as 10000
	*/
}
