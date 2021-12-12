package main

import "fmt"

func main(){
	name:="Partha"
	mother:="Manju"
	wife:= "Tonni"
	//print function
	fmt.Print("Hello")
	fmt.Print("World \n")
	//println function
	fmt.Println("hello")
	fmt.Println("world")
	// printf formatted string %v-> direct value %q -> " value " %T -> gives the time of variable

	fmt.Printf("my name is %v and my Mother name is %v and My wife name is %v \n",name,mother,wife)
	fmt.Printf("my name is %q and my Mother name is %q and My wife name is %q \n",name,mother,wife)

	//splintf function (saves in a variable)

	var data= fmt.Sprintf("my name is %v and my Mother name is %v and My wife name is %v \n",name,mother,wife)
	
	fmt.Print(data)
} 