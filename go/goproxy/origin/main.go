package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)

	fmt.Println("starting server on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received http request for", r.URL.Path)

	w.WriteHeader(200)
	fmt.Fprint(w, "Hello from Foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received http request for", r.URL.Path)

	w.WriteHeader(200)
	fmt.Fprint(w, "Hello from Bar")
}
