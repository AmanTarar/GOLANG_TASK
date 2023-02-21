package main

import (
    "fmt"
    "runtime"
)
//why it is not able to manage 2 Goroutines  
func say(s string) {

	
    for i := 0; i < 5; i++ {
		//runtime.Gosched()
        
        fmt.Println(s)
    }
}

func main() {
    
	fmt.Println("num of cpus:", runtime.NumCPU())
    go say("world")
	
    say("hello")
	fmt.Println("num of goroutines:", runtime.NumGoroutine())
}
