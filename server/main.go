package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// Request Context
	ctx := r.Context()

	log.Println("Request started")
	defer log.Println("Request ended")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request successfuly processed")
		w.Write([]byte("Request successfuly processed"))
	case <-ctx.Done():
		log.Println("Request cancelled")
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	}
}
