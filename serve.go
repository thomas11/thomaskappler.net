package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	dir := os.Args[1]
	port := ":9999"

	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("Serving %v on %v.", dir, port)
	http.ListenAndServe(port, nil)
}
