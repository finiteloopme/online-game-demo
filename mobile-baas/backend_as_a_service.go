package mbaas

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/time", timeHandler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Time from ["+hostname+"] is: "+time.Now().String())
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
