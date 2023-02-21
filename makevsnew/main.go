package main

import "fmt"

func main() {

	s := new([]int)
	*s = append(*s, 1)
	//s=append(*s[],1)

	fmt.Println("the cap is ",cap(*s))
	

}
