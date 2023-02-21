package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza aap between 1 and 5");

	reader:=bufio.NewReader(os.Stdin)

	fmt.Println("enter the rating")

	input,_:=reader.ReadString('\n');

	numrating,err:=strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err!=nil {
	
		fmt.Println(err)
	} else
	{
		fmt.Println("added 1 to your rating",numrating+1)
	}

}