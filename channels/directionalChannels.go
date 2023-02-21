package main

import (
	"fmt"
	"time"
)

func main() {


	c:=make(chan int)

	//receive
	  foo(c)

	//send
	go bar(c)

	 time.Sleep(time.Millisecond*30)
}

//receive
func foo(c chan<- int ){
	c<-23
}

//send

func bar(c <-chan int){

	fmt.Println("send",<-c)

}

