package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://globo.com")
	if err != nil {
		panic(fmt.Sprintf("error on http request: %v", err))
	}

	println("Status code", resp.StatusCode)
}
