package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>welcome to course</h2>")
}

func welcomeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>welcome to course. You've hit the welcomeFunc.</h2>")
}

// missing a Content-Type header

func main() {
	// routing  HandleFunc(pattern, function to run for given path pattern)
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/welcome", welcomeFunc)

	fmt.Println("pausing 2 secs...")
	time.Sleep(2 * time.Second)

	fmt.Println("starting the server on port 3000")
	http.ListenAndServe(":3000", nil)
}
