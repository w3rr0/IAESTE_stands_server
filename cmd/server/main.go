package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Ready")
	})

	log.Println("Server is running at :8080")
	http.ListenAndServe(":8080", nil)
}
