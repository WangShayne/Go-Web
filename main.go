package main

import (
	"net/http"
	_ "net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
