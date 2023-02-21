package main

import (
    "sync"
	"fmt"
)

func main() {
    cache := sync.Map{}
    cache.Store("key1", "value1")
    cache.Store("key2", "value2")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		val,ok:=cache.Load("key2")

		if ok{
			fmt.Println(val)
		}
	}()
    value, ok := cache.Load("key1")
    if ok {
        println(value.(string))
    }

	wg.Wait()
	
}