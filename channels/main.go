package main

import (
	"fmt"
	"time"
	//"runtime"
)

func main() {


	c:=make (chan int)
	
	
		
		c<-43
		c<-44
		c<-45
		
	
	// runtime.Gosched()
	// fmt.Println("middlework")
	//fmt.Printf("cap of channel c is: %d \n",cap(c))
	//fmt.Println("value exchange from channel",<-c)
	
	fmt.Println("value exchange from channel",<-c)
	time.Sleep(5*time.Second)
	fmt.Println("5 second hogya")
	fmt.Println("value exchange from channel",<-c)
	time.Sleep(1*time.Second)
	fmt.Println("1 second hogya")
	fmt.Println("value exchange from channel",<-c)
	
	

}

