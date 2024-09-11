package main

import (
	"fmt"
	isprime "semac-2024/is-prime"
	"time"
)

func main() {
	start := time.Now()

	maxNumber := 10
	done := make(chan bool)

	for i := 1; i <= maxNumber; i++ {
		go func(number int) {
			isprime.Print(number)
			done <- true
		}(i)
	}

	for i := 1; i <= maxNumber; i++ {
		<-done
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s", elapsed.String())
}
