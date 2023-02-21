package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
    // Create a new cache with a default expiration time of 5 minutes and a cleanup interval of 10 minutes
    // c := cache.New(5*time.Minute,10*time.Minute)

    // // Set a value in the cache with a key "mykey" and an expiration time of 1 hour
    // c.Set("mykey", "myvalue", 5*time.Second)

    // // Get a value from the cache with a key "mykey"
    // //time.Sleep(6*time.Second)
    // go func ()  {
    //     val,ok :=c.Get("mykey")
    //     if ok {
    //         fmt.Println("from go func",val)
    //     }else{

    //         fmt.Println("val not found in cache")
    //     }
        
    // }()
    // go func ()  {
    //     val,ok :=c.Get("mykey")
    //     if ok {
    //         fmt.Println("from go func2",val)
    //     }else{

    //         fmt.Println("val not found in cache")
    //     }
        
    // }()

    // value, found := c.Get("mykey")
    // if found {
    //     fmt.Println("Value found:", value)
    // } else {
    //     fmt.Println("Value not found")
    // }

    // cnew:=c.Items()
    // fmt.Printf("%T\n",cnew)

    // fmt.Println("map value of c cache",cnew)
    // fmt.Println("")

    // err:=c.SaveFile("newCache")
    // if err!=nil{
    //     fmt.Println("data not send to cache")
    // }
    // time.Sleep(2*time.Second)
    // Delete a value from the cache with a key "mykey"
    // c.Delete("mykey")
    it:=map[string]cache.Item{

        "key1":cache.Item{
            Object:1,
            Expiration:time.Now().Add(5*time.Minute).UnixNano(),
        },
    }
    cacheNew:=cache.NewFrom(5*time.Minute,10*time.Minute,it)
    b,rr:=cacheNew.IncrementInt64("key1",5)
    if rr!=nil{
        fmt.Println("error")
    }
    fmt.Println("b value from cache new",b)
    // value,ok:=cacheNew.Get("key1")
    // if ok {
    //     fmt.Println("Value found:", value)
    // } else {
    //     fmt.Println("Value not found")
    // }
   
}

