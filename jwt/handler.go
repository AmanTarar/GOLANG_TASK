package main

import (
	"encoding/json"
	"net/http"
)


var jwtKey=[]byte("secret key")

var users=map[string]string{

	"user1": "pass1",
	"user2": "pass2",
}

type Credentials struct {

Username string `json:"username"`
Password string `json:"password"`

}

type Claims struct{
	Username string `json:"username"`
	jwt.StandardClaims
}
func Login(w http.ResponseWriter, r *http.Request){

	var credentials Credentials
	err:=json.NewDecoder(r.Body).Decode(&credentials)
	if err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		return
		
	}

	expectedPassword,ok :=users[credentials.Username]

}

func Home(w http.ResponseWriter, r *http.Request){

	
}

func Refresh(w http.ResponseWriter, r *http.Request){

	
}