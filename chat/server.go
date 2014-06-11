package main

import (
	"fmt"
	"log"
	"net/http"
        "encoding/json"
)

type User struct {
        Login string
        Password string
}

func Auth(w http.ResponseWriter, r *http.Request) {
        var user User
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
                http.Error(w, err.Error(), 500)
        }

	fmt.Fprintf(w, "logged in %s!", user.Login)
}

func main() {
	http.HandleFunc("/auth", Auth)
	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
