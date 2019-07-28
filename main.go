package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get the port from the environment, either set this or if deploying to heroku it will be available
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set, defaulting to 8080")
		port = "8080"
	}

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":"+port, nil)
}

// HelloServer Sample response
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s! And don't forget, Go Joey!!", r.URL.Path[1:])
}
