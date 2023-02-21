package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome:="welcome to userinput"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("enter the eating for the order:")

	//comma ok// err err

	input,_:=reader.ReadString('\n')
	fmt.Println("thanks for the rating",input)
	


}