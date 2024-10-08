package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaverickPallav/go_mongoapi/router"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the MongoDB API!")
}

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	
	fmt.Println("Server is getting started...")

	r.HandleFunc("/", WelcomeHandler)

	log.Fatal(http.ListenAndServe(":4001", r))
}
