package main

import (
	"fmt"
	"math"
	"time"
)

var MAX_INT = 100000000 //100 milli
var totalPrimeNums int32 = 0

// Sequential Approach
func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	totalPrimeNums++
}

func main() {
	start := time.Now()
	for i := 3; i < MAX_INT; i++ {
		checkPrime(i)
	}
	end := time.Now()
	fmt.Printf("Sequential Approach: Total Prime Numbers = %d, Time taken = %v\n", totalPrimeNums+1, end.Sub(start))
}
