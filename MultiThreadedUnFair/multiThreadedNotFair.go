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

func doBatch(name string, wg *sync.WaitGroup, nstart int, nend int) {
	defer wg.Done()

	start := time.Now()

	for i := nstart; i <= nend; i++ {
		checkPrime(i)
	}

	fmt.Printf("Thread %s [%d, %d] completed in %s\n", name, nstart, nend, time.Since(start))
}

func main() {

	start := time.Now()

	var wg sync.WaitGroup
	nstart := 3
	batchSize := int(float64(MAX_INT) / float64(CONCURRENCY))

	for i := 0; i < CONCURRENCY-1; i++ {
		wg.Add(1)

		go doBatch(strconv.Itoa(i), &wg, nstart, nstart+batchSize)
		nstart += batchSize
	}

	wg.Add(1)

	go doBatch(strconv.Itoa(CONCURRENCY-1), &wg, nstart, MAX_INT)

	wg.Wait()

	fmt.Println("Checking till", MAX_INT, "found", totalPrimeNums+1, "prime numbers took", time.Since(start))

}
