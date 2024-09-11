package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("GET method\n"))
		}

		if r.Method == "POST" {
			w.Write([]byte("POST method\n"))
		}

		w.Write([]byte("Hello, world!\n"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("error on http server")
	}
}
