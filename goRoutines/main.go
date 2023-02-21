package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func greet(s string) {

	for i := 0; i < 5; i++ {

		fmt.Println(s)
		
	}
	defer wg.Done()
	
}
// func greetbro(s string) {

// 	for i := 0; i < 5; i++ {

// 		fmt.Println(s)
		
// 	}
// 	wg.Done()
// }

func main() {


	
	 wg.Add(1)
	 go greet("bro")
	 
	
	  wg.Wait()
}

