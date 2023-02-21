package main

import "fmt"

func main() {
	a := [3][3]int{
		{1, 0, 1},
		{1, 1, 0},
		{0, 1, 0},
	}

	//right diagonal
	if a[0][0] == a[1][1] && a[1][1] == a[2][2] {
		if a[0][0] == 1 {
			fmt.Print("1 wins")
			return
		} else {
			fmt.Print("0 wins")
			return
		}
	}

	//left diagonal
	if a[0][2] == a[1][1] && a[1][1] == a[2][0] {
		if a[0][0] == 1 {
			fmt.Print("1 wins")
			return
		} else {
			fmt.Print("0 wins")
			return
		}
	}

	//rows change
	for i := 0; i < 3; i++ {
		if a[i][0] == a[i][1] && a[i][1] == a[i][2] && a[i][0] == 1 {
			if a[i][0] == 1 {
				fmt.Println("1 wins")
				return
			} else {
				fmt.Println("0 wins")
				return
			}
		}
	}

	// columns change
	for j := 0; j < 3; j++ {
		if a[0][j] == a[1][j] && a[1][j] == a[2][j] && a[0][j] == 1 {
			if a[0][j] == 1 {
				fmt.Println("1 wins")
				return
			} else {
				fmt.Println("0 wins")
				return
			}
		}
	}

	fmt.Println("Draw")
	return

}
