package main

import (
	"fmt"
	"log"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "logged in!")
}

func main() {
	http.HandleFunc("/auth", Auth)
	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
