package main

import (
	"fmt"
	isprime "semac-2024/is-prime"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	maxNumber := 1000
	wg := sync.WaitGroup{}
	wg.Add(maxNumber)
	for i := 1; i <= maxNumber; i++ {
		go func(number int) {
			isprime.Print(number)
			wg.Done()
		}(i)
	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s", elapsed.String())
}
