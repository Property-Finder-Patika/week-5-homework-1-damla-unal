package main

import (
	"fmt"
	"time"
)

/*

Race Condition occurs when multiple threads(two or more goroutines) access and modify a shared memory area at the same time.
In race condition we get inconsistent results.
A critical section is a block of code where a goroutine attempts to write a shared variable.

*/

const limit = 5000

var counter = 0 //counter which is what all goroutines will try to modify simultaneously

// increment increases the counter variable up to the limit
func increment() {
	fmt.Printf("Counter before increment: %d\n", counter)
	for i := 0; i < limit; i++ { // each routine is expected to increase the counter 5k
		counter += 1
	}
	fmt.Printf("Counter after increment: %d\n", counter)
}

func main() {
	//two goroutine are started
	// we cannot have a race condition if we only have a single goroutine
	go increment() //Each goroutine reads the value from counter, increments it,
	go increment()

	time.Sleep(1 * time.Millisecond)
	fmt.Printf("Counter is expected to be %d but found as %d", 2*limit, counter)

	/* Expected Output:
	Counter before increment: 0
	Counter before increment: 0
	Counter after increment: 5000
	Counter after increment: 9938
	Counter is expected to be 10000 but found as 9938
	*/
}

/*
If you run the code with the built-in race condition checker, the go compiler will complain about the problem.

  `A small note about the Golang race condition checker: if your code occasionally accesses shared variables,
   it might not be able to detect the race condition.
   To detect it, the code should run in heavy load, and race conditions must be occurring.`

  >> go run -race ./race_condition

	You can see the output of the race condition checker:

	==================
	WARNING: DATA RACE
	Write at 0x0001028c58e0 by goroutine 7:
	  main.increment()
		  /Users/damlaunal/Projects/Patika/week-5-homework-1-damla-unal/race_condition/race_condition.go:24 +0xbc

	Previous read at 0x0001028c58e0 by goroutine 8:
	  main.increment()
		  /Users/damlaunal/Projects/Patika/week-5-homework-1-damla-unal/race_condition/race_condition.go:22 +0x30

	Goroutine 7 (running) created at:
	  main.main()
		  /Users/damlaunal/Projects/Patika/week-5-homework-1-damla-unal/race_condition/race_condition.go:32 +0x2c

	Goroutine 8 (running) created at:
	  main.main()
		  /Users/damlaunal/Projects/Patika/week-5-homework-1-damla-unal/race_condition/race_condition.go:33 +0x38
	==================

*/
