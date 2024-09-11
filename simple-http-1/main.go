package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!\n"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("error on http server")
	}
}
