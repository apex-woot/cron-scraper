package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hello from handler</h1>`)
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.ListenAndServe(":80", nil)
	fmt.Printf("Start listening on port 3000...\n")
}
