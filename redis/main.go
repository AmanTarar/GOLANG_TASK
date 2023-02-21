package main

import (
	"fmt"
	"log"
	"github.com/gomodule/redigo/redis"
)

func checkError(err error){
	if err!=nil{
	log.Fatal(err)
	}
}
func main() {

conn,err:=redis.Dial("tcp","localhost:6379")
checkError(err)
defer conn.Close()

_,err=conn.Do(
	"HMSET",
	"podcast:1",
	"title",
	"tech over tea",
)
checkError(err)
title,err:=redis.String(conn.Do("HGET","podcast:1","title"))
checkError(err)
fmt.Println("title of podcast:",title)

}