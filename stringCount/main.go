package main

import (
	"fmt"
	"strings"
)



func StringCount(s string) map[string]int {

	arr:=strings.Fields(s)
	mp := make(map[string]int)
	for _, v := range arr {
		
		mp[v]++
	}

	return mp
}

func runesToString(runes []rune) (outString string) {
    
    for _, v := range runes {
        outString = outString+" "+ string(v)
    }
    return 
}




func main() {

	input:=[]rune{'c', 'd', 'b', 'd', 'a', 'a', 'd', 'b', 'd', 'c', 'e', 'c', 'h', 'e', 'b', 'c', 'f', 'a', 'e', 'd', 'b', 'b', 'b', 'h', 'f', 'f', 'b', 'h', 'g', 'e', 'g', 'g', 'b', 'c', 'b', 'c', 'h','z', 'b', 'h', 'a', 'b', 'e', 'a', 'c', 'e', 'd', 'h', 'c', 'b', 'h', 'h', 'c', 'b', 'f', 'h', 'e', 'c', 'c', 'a', 'h', 'g', 'a', 'g', 'a', 'e', 'c', 'f', 'h','q' ,'c', 'd', 'b', 'f', 'g', 'c', 'e', 'f', 'a', 'e', 'f', 'f', 'a', 'g', 'g', 'g', 'g', 'h','o', 'a', 'a', 'e', 'b', 'g', 'f', 'c', 'g', 'a', 'a', 'h', 'a', 'f', 'f'}
	inputSTR:=runesToString(input)
	fmt.Println("StringCount program")
	// mp := make(map[string]int)
	// mp=StringCount(inputSTR)
	//fmt.Println(inputSTR)

	(StringCount(inputSTR))
	fmt.Println((StringCount(inputSTR)))



}