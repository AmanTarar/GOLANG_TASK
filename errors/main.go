package main

import (
	"fmt"
	"log"
	//"io"
	"os"
	//"strings"
)

func main() {

 file,err:=os.Create("log2.txt")

if err!=nil{

	fmt.Println(err)
    //os.Exit(1)

}
log.SetOutput(file)

defer file.Close()

file2,err:=os.Open("no files.txt")
if err!=nil{

    log.Println("error happened in opening no file",err)
}

defer file2.Close()

 

}

