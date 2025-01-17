package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yanisdib/tomodachi/server/config"
)

func main() {
	config.OpenDBConnectionPool()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("TOMODACHI API is running on http://localhost:5050")
	})

	err := http.ListenAndServe(":5050", mux)
	if err != nil {
		log.Fatalf("An error occured while starting the server.")
	}

	log.Println("Server is running")
}
