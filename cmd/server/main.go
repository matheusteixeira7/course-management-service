package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "3333"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test")
	})

	log.Printf("Listening server on port %v", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Error on listening server: %v", err)
	}
}
