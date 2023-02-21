package main

import (
	"fmt"
	//"runtime"
)

func main() {


	c:=make (chan int)
	
	go func() {
		
		c<-43
		
	}()
	// runtime.Gosched()
	// fmt.Println("middlework")
	fmt.Printf("%d",cap(c))
	fmt.Println("value exchange from channel",<-c)
	//fmt.Println("value exchange from channel",<-c)

}

