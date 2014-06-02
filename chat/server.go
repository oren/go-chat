package main

import (
	"fmt"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "logged in!")
}

func main() {
	http.HandleFunc("/auth", Auth)
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
