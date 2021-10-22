package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>welcome to course</h2>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting the server on port 3000")
	http.ListenAndServe(":3000", nil)
}
