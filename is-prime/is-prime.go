package isprime

import (
	"fmt"
	"time"
)

func Print(n int) {
	if n == 1 {
		time.Sleep(1 * time.Second)
		fmt.Println("1 is not prime")
		return
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d is not prime\n", n)
			return
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("%d is prime\n", n)
	return
}
