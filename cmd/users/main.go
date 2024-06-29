package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello!\n")
}

func main() {
	http.HandleFunc("/sayHello", sayHello)

	http.ListenAndServe("localhost:8090", nil)
}
