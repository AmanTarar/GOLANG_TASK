package main

import (
	"fmt"
	"net/http"
	"io"
)

type hotdog int
func main() {

	var h hotdog=2
	
	// http.HandleFunc("/",foo)
	http.ListenAndServe(":8080",h)
}

func (h hotdog)ServeHTTP(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w,"hjjhbjh",h)
		//w.Write([]byte(h))
		fmt.Println("vale of hotdog is",h)
}

func dummyhandler(w http.ResponseWriter,r *http.Request){

	io.WriteString(w,"welcome to dummy handler mode")

}

