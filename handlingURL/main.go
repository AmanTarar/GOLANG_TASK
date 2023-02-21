package main

import (
	"fmt"
	//"net/http"
	"net/url"
)


const myurl string="http://3000/ulta?Promax=1233454545&yyy=nnn"

func main() {


	

	fmt.Println(myurl)

	result,_:=url.Parse(myurl)

	fmt.Printf("type of result from url: %T\n", result)
	fmt.Println(result.Query())
	fmt.Println(result.RawQuery)
	qparams:=result.Query()

	for _,v:=range qparams{

		fmt.Println("values inside qparams",v)
	}

	partsofURL :=&url.URL{
		Scheme:"http",
        Host:   "lco.dev",
        Path:   "/tutcss",
		RawPath: "user=hitesh",
    }
	
	anotherURL :=partsofURL.String()

	fmt.Println(anotherURL)


}

