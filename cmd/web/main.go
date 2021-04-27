package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	mux := routes()
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Println("Error starting application on port:", port, err)
		return
	}
	log.Println("Starting application on port", port)

}
