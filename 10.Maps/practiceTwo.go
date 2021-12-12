package main

import "fmt"

func main(){

	menu:= map[string]int64{
		"Age":25,
		"Number":1234,
	} 

	for key,value:= range menu{
		fmt.Printf("key is %q value is %v\n",key,value)
	}
	fmt.Println(menu)
}