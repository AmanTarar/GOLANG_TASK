package main

import (
	"fmt"
	"sync"
)

func main() {

	c:=make(chan int)
	input:=make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	var val int

	go func(){

		input<-1

		//defer close(input)
		defer wg.Done()
	}()

	go func (){

		fmt.Println("value inserted inside c",<-c)
		defer wg.Done()
	}()
	 go func(){

        val=<-input
		fmt.Println("value taken out from input-->",val)
		c<-val

	}()

	wg.Wait()
	fmt.Println("program finished")
}

// func main(){
// 	c := make(chan int)
// 	i := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	var 
// 	go func(){
// 		fmt.Println(<-c)
// 	}()
// 	go func(){
// 		i<-1
// 		val<-i
// 		close(i)
// 		close(c)
// 		wg.Done()
// 	}()
// 	defer wg.Wait()
	
// }