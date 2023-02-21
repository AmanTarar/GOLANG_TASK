// package main

// import (
// 	"fmt"
	
// 	"runtime"
// 	"sync"
// )



// func main()  {

// 	fmt.Println("welcome to raceCondition")
// 	fmt.Println("no. of CPUs",runtime.NumCPU())

// 	counter:=0
// 	const gs=50
// 	var wg sync.WaitGroup
// 	wg.Add(gs)
// 	var mu sync.Mutex
// 	for i:=0;i<gs;i++ {
// 		go func() {
// 				mu.Lock()
//                 v:=counter
// 				runtime.Gosched()
// 				v++
				
// 				counter=v
//                 //fmt.Println(v)
				
// 				wg.Done()
// 				mu.Unlock()
//         }()
// 		fmt.Println("no. of goroutines",runtime.NumGoroutine())
// 	}
// 	wg.Wait()
// 	fmt.Println("counter",counter)
	

	
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	
	defer c.mu.Unlock()
	return c.v[key]
	
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}