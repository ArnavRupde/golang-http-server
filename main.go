package main

import (
	"fmt"
	"io"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Status: OK")
}
func main() {
	fmt.Println("Starting HTTP Server on port 8000")
	http.HandleFunc("/health", healthCheckHandler)
	http.ListenAndServe(":8000", nil)
}
