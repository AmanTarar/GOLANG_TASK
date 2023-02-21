package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}

func validate(username, password string) bool {
	// You can implement your own validation logic here
	if len(password)<8 {
		fmt.Println("password is too short")
		return false
		
	}
	return true
}

func main() {
	http.HandleFunc("/", BasicAuth(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	}))
	http.ListenAndServe(":8080", nil)
}



