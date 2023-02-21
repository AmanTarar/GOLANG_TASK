package main

import "fmt"



func horizontal(arr [3][3]int)int  {
	
	i:=0
	j:=0
	var res int=-1

	for i=0;i<len(arr);i++{

		if arr[i][j]==arr[i][j+1] && arr[i][j+1]==arr[i][j+2]{
			res=arr[i][j]
		}

	}
	return res
}

func vertical(arr[3][3]int)int  {
	i:=0
	j:=0
	var res1 int=-1

	for j=0;j<len(arr[i]);j++{

		if arr[i][j]==arr[i+1][j]&& arr[i+1][j]==arr[i+2][j]{

			res1=arr[i][j]
		}

	}
	return res1
	
}

func diagonal(arr[3][3]int)int  {
	
	
	var res2 int =-1

	if (arr[0][0]==arr[1][1] && arr[1][1]==arr[2][2])||(arr[0][2]==arr[1][1]&&arr[1][1]==arr[2][0]){
		
		res2 = arr[1][1]
	}

	return res2
}

func winner(arr [3][3]int)int {

	var result int=-1
	if horizontal(arr)!=-1{
		result=horizontal(arr)
		fmt.Println("horizontal win by:",result)
		//return result
	}
	if vertical(arr)!=-1{
		result=vertical(arr)
		fmt.Println("vertical win by:",result)
		//return result
	}
	if diagonal(arr)!=-1{
		result=diagonal(arr)
		fmt.Println("diagonal win by:",result)
		//return result
	}

	return result



	
	
}

func main() {


	var arr =[3][3]int{{1,0,0},{1,0,1},{0,0,1}}
	//winner(arr)
	fmt.Println("winner of tictactoe game is: ",winner(arr))
}