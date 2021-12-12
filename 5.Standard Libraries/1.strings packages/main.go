package main

import (
	"fmt"
	"strings"
)

func main(){
	
	greetings:="hello"


	fmt.Println(strings.Contains(greetings,""))
	fmt.Println(strings.ReplaceAll(greetings, "he" ,"hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings,"o"))
	fmt.Println(strings.Split(greetings," "))
}