package main

import (
	"fmt"

	
)

func main() {

	fmt.Println("welcome to mystructs")
	aman:=User{"aman","aman@01",true,19}
	fmt.Println("user details in println:",aman)
	fmt.Printf("user details in printf %v",aman);


}

type User struct{

	Name	string
	email	string
	status	bool
	age		int
}