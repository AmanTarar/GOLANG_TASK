package main

import "fmt"

func main()  {
	

 input:=[]int{1,2,3,4,5,6,7,8,9}

//  res:=sum(input...)
//  fmt.Println(res)

 evensum:=even(sum,input...)
 fmt.Println("sum of even numbers--->",evensum)
}

func sum(xi...int)int  {

	total:=0

	for _,v:=range xi {

		total+=v
	}


	return total
	
}

func even(f func(xi...int)int,yi...int)int  {

	evennum:=[]int{}

	for _,v:=range yi{

		if v%2==0{

			evennum=append(evennum,v)
		}
	}

	return f(evennum...)
		// res:=0
		// for _,v range evennum{

		// 	res+=v
		// }
		// return res
	// }
	
}