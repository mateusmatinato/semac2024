package main

import (
	"fmt"
	isprime "semac-2024/is-prime"
	"time"
)

func main() {
	start := time.Now()

	maxNumber := 1000
	for i := 1; i <= maxNumber; i++ {
		isprime.Print(i)
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s", elapsed.String())
}
