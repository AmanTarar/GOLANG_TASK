package main

import "fmt"

func main() {

	fmt.Println("welcome to arrays and slices");

	highscores:=make([]int,0);

	scores:=highscores[:0]
	scores=append(scores,1)
	scores=append(scores,1)
	scores=append(scores,1)
	scores=append(scores,1)
	scores=append(scores,1)
	scores=append(scores,1)
	scores=append(scores,1)
	scores=append(scores,1)
	scores=append(scores,1)
	// scores2:=scores[:0]
	// scores2=append(scores2,1)
	// scores2=append(scores2,1)
	// scores2=append(scores2,1)
	// scores2=append(scores2,1)
	// scores2=append(scores2,1)
	// scores=append(scores,2)
	// scores=append(scores,3)

	//fmt.Printf("type of the new array %T \n",highscores);
	// fmt.Println(highscores[0]);
	

	//fmt.Println("length: ",len(highscores),"capacity: ",cap(highscores))
	fmt.Println("length: ",len(scores),"capacity: ",cap(scores))
	//fmt.Println("length: ",len(scores2),"capacity: ",cap(scores2))
}