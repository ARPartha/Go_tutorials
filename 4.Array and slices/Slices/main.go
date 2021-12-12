package main

import (
	"fmt"
)

func main(){
	//slices (use arrys behind)
	var variable =[]int {1,2,3,4}
	fmt.Println(variable,len(variable))

	//append value is slice
	var vari =[]int{} // shortform of declaring slice
	vari= append(vari,4,5,6,7)
	fmt.Println(vari,len(vari))

	//slice ranges
	rangeOne:=vari[0:3]
	fmt.Println(rangeOne,len(rangeOne))

	rangeTwo:=vari[2:]
	fmt.Println(rangeTwo, len(rangeTwo))

	rangeThree:=vari[:3]
	fmt.Println(rangeThree,len(rangeThree))

}