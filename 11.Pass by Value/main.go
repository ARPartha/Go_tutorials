package main

import "fmt"

//update the original value
func updateName(x string)string{
  	x="Abdur Rahman "
	fmt.Print(x)
	return x
}

//don't update the original value
func dontUpdatename(x string){
	x="Abdur Rahman \n"
  fmt.Print(x)
}

func main(){

	name:="partha"
	nameTwo:="Partha"
	dontUpdatename(nameTwo)
 	name=updateName(name)
	fmt.Println(name+"\n")
	fmt.Println(nameTwo)

}