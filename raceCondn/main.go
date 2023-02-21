package main

import "fmt"
import "runtime"
import "sync"



func main()  {

	fmt.Println("welcome to raceCondition")
	fmt.Println("no. of CPUs",runtime.NumCPU())

	counter:=0
	const gs=100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i:=0;i<gs;i++ {
		go func() {

                v:=counter
				v++
				counter=v
                fmt.Println(v)
				wg.Done()
        }()
		fmt.Println("no. of goroutines",runtime.NumGoroutine())
	}
	wg.Wait()
	

	
}