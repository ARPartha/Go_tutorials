package main

import (
	"fmt"
	"strings"
)

// return multiple values
func getInitials(n string)(string,string){  
	init:=[]string{}
	s:= strings.Split(n," ")
	for _,val:=range s{
		init=append(init,val)
	}
	if len(init)>1{
		return init[0],init[1]
	}
	return init[0],""
}

func main(){

	// multi value check 
	f1,f2:=getInitials("Ar Partha") //function calling
	fmt.Println(f1,f2)
	f3,f4:=getInitials("Marium Begum") //function calling
	fmt.Println(f3,f4)
	f5,f6:=getInitials("Nafisa Tonni") //function calling
	fmt.Println(f5,f6)
// single value check
	f7,f8:=getInitials("Partha")//function calling
	fmt.Println(f7,f8)
	f9,f10:=getInitials("Marium")//function calling
	fmt.Println(f9,f10)
	f11,f12:=getInitials("Tonni")//function calling
	fmt.Println(f11,f12)

}