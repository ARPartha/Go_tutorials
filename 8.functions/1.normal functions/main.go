package main

import "fmt"

//single value
func sayGreeting(n string){
	fmt.Printf("hello %v \n",n)
}

//multi value
func multiValue(n[]string, f func(string)){
	for _,value:=range n{
		f(value)
	}
}
func main(){

	array:=[]string{"Partha","Manju","Tonni"}
	multiValue(array,sayGreeting)
	// sayGreeting("Partha")
	fmt.Println()
}

