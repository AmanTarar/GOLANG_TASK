package main

import (
	"encoding/json"
	"fmt"
)

type book struct{

	Author string  `json:"author"`
	Price Price      `json:"price"`
}

type Price struct{

	Amount int  `json:"amount"`
	Currency string  `json:"curr"`
}

func main() {

	fmt.Println("welcome to json in go")
	sfrom:=`{"author":"aman","price":{"amount":100,"curr":"dollar"}}`
	

	var book1 book

	err:=json.Unmarshal([]byte (sfrom), &book1)
	if err!=nil{
		fmt.Println("error aagya ji")
	}
	fmt.Println(book1)

	sto,err:=json.Marshal(book1)
	if err!=nil{
		fmt.Println("error aagya ji")
	}

	fmt.Println(string(sto))

	

	

	



}