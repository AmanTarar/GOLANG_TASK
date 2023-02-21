package main

import (
	"fmt"
	"sync/atomic"

	"runtime"
	"sync"
)



func main()  {

	fmt.Println("welcome to raceCondition")
	fmt.Println("no. of CPUs",runtime.NumCPU())

	var counter int64
	const gs=50
	var wg sync.WaitGroup
	wg.Add(gs)

	
	for i:=0;i<gs;i++ {
		go func() {
				atomic.AddInt64(&counter,1)
                
				runtime.Gosched()
				atomic.LoadInt64(&counter)
				
				wg.Done()
				
        }()
		fmt.Println("no. of goroutines",runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("counter",counter)
	

	
}