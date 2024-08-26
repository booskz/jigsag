package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
