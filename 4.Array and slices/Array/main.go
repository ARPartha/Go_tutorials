package main

import "fmt"

func main(){

	//declaring array
	var array [3]int =[3]int {1,2,3}
	fmt.Println(array, len(array))

	// shortform  of declaring array
	var arrayTwo =[3]int{4,5,6}
	fmt.Println(arrayTwo, len(arrayTwo))

	// function variable shorthand declaration
	arrayThree:=[3]int{7,8,9}
	fmt.Println(arrayThree, len(arrayThree))

	
}