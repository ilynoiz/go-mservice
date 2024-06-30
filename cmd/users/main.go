package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"context"
	"time"
	
	"../../internal/users/middleware"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello!\n")
}

// func login(w http.ResponseWriter, r *http.Request) {
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
// 	fmt.Printf("name = %s, password = %s", username, password)
// 	//TODO: send username and password to server by grpc
// }

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/sayHello", sayHello)
	mux.HandleFunc("/login", AuthMiddleware)

	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), mux))

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	// http.ListenAndServe("localhost:8090", nil)
}
