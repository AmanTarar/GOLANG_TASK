package main

import "fmt"

func main() {


	c:=make(chan int)

	go foo(c)
	fmt.Printf(" cap of c--> %d",cap(c))

	//bar(c)
	func (c <-chan int)  {

		for v:=range c{
			fmt.Println(v)
		}
		
	
	}(c)
	fmt.Printf(" cap of c--> %T\n",c)
	fmt.Printf(" len of c--> %d",cap(c))





}

func foo(c chan<- int)  {

	for i := 0; i < 5; i++ {
		c<-i
	}
	close(c)
	
}



