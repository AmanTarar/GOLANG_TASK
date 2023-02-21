package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"]="javascript";
	languages["RB"]="ruby";
	languages["PY"]="python";
	title,ok:=languages["JSx"]
	fmt.Println(title,"  ",ok)



	fmt.Println("languages map is",languages);
	// fmt.Println(languages[0]);
}