package main

import (
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/invert", invert)
	http.HandleFunc("/sum", sum)
	http.HandleFunc("/multiply", multiply)
	http.HandleFunc("/flatten", flatten)
	http.ListenAndServe(":8080", nil)
}
