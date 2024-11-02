package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100000000 //100 milli
var totalPrimeNums int32 = 0
var currentNum int32 = 2

var CONCURRENCY = 10 //10 threads
func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}

	//Doing atomic increment, for correctness since multiple threads will be incrementing it
	atomic.AddInt32(&totalPrimeNums, 1)
}

// So here basically all the 10 thread will run this function
// Inside it, an infinite loop will run until 100 million,
// and all thread will keep picking, checking and updating the current number
func doWork(name string, wg *sync.WaitGroup) {
	start := time.Now()

	defer wg.Done()

	for {
		x := atomic.AddInt32(&currentNum, 1)

		if x > int32(MAX_INT) {
			break
		}

		checkPrime(int(x))

	}

	fmt.Printf("Thread %s completed in %s\n", name, time.Since(start))
}

func main() {

	start := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)

		go doWork(strconv.Itoa(i), &wg)
	}

	wg.Wait()

	fmt.Println("Checking till", MAX_INT, "found", totalPrimeNums+1, "prime numbers took", time.Since(start))

}
