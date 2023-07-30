package main

import (
	"fmt"
	"sync"
)

// This is a small nice coding challenge for dealing with concurrency and thread synchronization and atomic variable
// manipulation between thread which implies dealing with race-condition avoidance mechanism. The task is to have a
// dedicated thread printing odd number and another printing even number until reaching a target common value. The
// printing at scren should be into ascending order from 1 to Target number. Each thread should have its name in front
// of its printing. See below an expected output for target value 20. If the target value is 0, nothing should be displayed.
// Thread denotes here to goroutine. Use -race flag when running the code to make sure it is race free or not.

// A technique is to use two goroutines (pseudo threads) and provide them

/*

Thread-1 :  1
Thread-2 :  2
Thread-1 :  3

<snip>

Thread-2 :  18
Thread-1 :  19
Thread-2 :  20

*/

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 09 August 2021

// printOddNumbers is considered as thread 1 to print odd numbers with Thread-1 as prefix.
func printOddNumbers(max int, currentValuePtr *int, mutex *sync.Mutex, waitgroup *sync.WaitGroup) {
	nextValue := 0
	for {
		// acquire the lock so other goroutine will wait.
		mutex.Lock()
		if *currentValuePtr < max {
			// increment current val and check if odd.
			nextValue = *currentValuePtr + 1
			// check if the new value is odd - if so print the incremented
			// value and assign it as the new current value. Otherwise, ignore
			// and release the lock so other goroutine will print that value.
			if nextValue%2 != 0 && nextValue <= max {
				fmt.Println("Thread-1 : ", nextValue)
				*currentValuePtr = nextValue
			}
			// release the lock.
			mutex.Unlock()
			continue
		} else {
			// release the lock and quit.
			mutex.Unlock()
			break
		}
	}
	// outside loop so notify end of the goroutine operation.
	waitgroup.Done()
}

// printEvenNumbers is considered as thread 2 to print even numbers with Thread-2 as prefix.
func printEvenNumbers(max int, currentValuePtr *int, mutex *sync.Mutex, waitgroup *sync.WaitGroup) {
	nextValue := 0
	for {
		// acquire the lock so other goroutine will wait.
		mutex.Lock()
		if *currentValuePtr < max {
			// temporarily increment the current value.
			nextValue = *currentValuePtr + 1
			// check if the new value is even - if so print the incremented
			// value and assign it as the new current value. Otherwise, ignore
			// and release the lock so other goroutine will print that value.
			if nextValue%2 == 0 && nextValue <= max {
				fmt.Println("Thread-2 : ", nextValue)
				*currentValuePtr = nextValue
			}
			// release the lock.
			mutex.Unlock()
			continue
		} else {
			// reached end of target number so release the lock and quit the loop.
			mutex.Unlock()
			break
		}
	}
	// outside loop so notify end of the goroutine operation.
	waitgroup.Done()
}

func main() {
	// mutex for the goroutines synchronization
	mutex := &sync.Mutex{}
	// helps block main func until goroutines done.
	var waitgroup sync.WaitGroup
	// we have two goroutines here.
	waitgroup.Add(2)
	// store the value to be displayed at a given time.
	var currentValue int

	// maximum positive value for testing purpose. Increase if you want.
	max := 10

	// start thread-1 responsible for printing odd numbers.
	go printOddNumbers(max, &currentValue, mutex, &waitgroup)
	// start thread-2 responsible for printing even numbers.
	go printEvenNumbers(max, &currentValue, mutex, &waitgroup)
	// wait until above two goroutines are done.
	waitgroup.Wait()
}

// go run -race goroutines-race-free-even-odd-numbers-printing.go
// Output:
// Thread-1 :  1
// Thread-2 :  2
// Thread-1 :  3
// Thread-2 :  4
// Thread-1 :  5
// Thread-2 :  6
// Thread-1 :  7
// Thread-2 :  8
// Thread-1 :  9
// Thread-2 :  10
