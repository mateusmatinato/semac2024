package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	googleResp := make(chan int)
	bingResp := make(chan int)
	go getUrl("https://google.com", googleResp)
	go getUrl("https://bing.com", bingResp)

	select {
	case googleSC := <-googleResp:
		println("Google status code", googleSC)
	case bingSC := <-bingResp:
		println("Bing status code", bingSC)
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s", elapsed.String())
}

func getUrl(s string, resp chan int) {
	httpResp, err := http.Get(s)
	if err == nil {
		resp <- httpResp.StatusCode
		return
	}
}
