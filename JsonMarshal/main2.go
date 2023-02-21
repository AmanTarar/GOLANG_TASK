package main

import (
	"encoding/json"
	"fmt"
)

type person struct{

	name string  
	Age int
	Creditcard string 

}
defer func A init(){}
func main() {

	fmt.Println("welcome to marshal 2.0")

	p:=person{name:"aman",Creditcard:"super secret"}

	pbytes,err:=json.Marshal(p)

	if err!=nil{


		fmt.Println("error aagya ji")
	}
	// log.Print(err)

	// log.Print(string(pbytes))

	fmt.Println(string(pbytes))



}