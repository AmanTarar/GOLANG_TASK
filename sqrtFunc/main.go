package main

import "fmt"

func main() {

	fmt.Println("welcome to sqrt function in golang")

	



}

func sqrt(input int) float64  {

	prime:=[]int{2,3,5,7,11,13,17,19}

	for _,p:=range prime{
		curr:=input/p

		if(input%p)==0 && curr!=1{
			
		}


	}
	
}