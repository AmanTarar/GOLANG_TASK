package main

import (
	"fmt"
	//"time"
)

func foo(c chan int,somevalue int){

    //time.Sleep(5*time.Second)
    c<-somevalue *5
    

}

func main() {

fmt.Println("working with channels")

fooval:=make(chan int)

go foo(fooval,5)
go foo(fooval,8)

v1:=<-fooval
v2:=<-fooval
fmt.Println("v1:",v1)
fmt.Println("v2:",v2)

fmt.Println("program finished")

}