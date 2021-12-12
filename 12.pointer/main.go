package main

import "fmt"


func updateAge(x *int)int{
	//dereffercence of printing value
	fmt.Println("pointer pass age:",*x)
	*x =43434
	return *x
}

func main(){

	name:= "partha"
	age:=25
	m:=&age
	*m=34
	//passing pointer in function
	n:= updateAge(m)
	fmt.Println(n)
	fmt.Println(&name)
	fmt.Println(age)

}