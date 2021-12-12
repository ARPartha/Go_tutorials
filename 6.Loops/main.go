package main

import "fmt"

func main(){

	x:=0
	for x<5 {
		// fmt.Println(x)
		x++
	}

	fmt.Println()

	for i:=0 ;i<5;i++{
		// fmt.Println("I ",i)
	}

	names:=[]string{"a","b","c","d","e","f","g"}

	for i:=0;i<len(names);i++{
		// fmt.Print(names[i])
	}

	for index,value :=range names{
		fmt.Printf("index: %v value:%q \n",index,value)
	}
	for _,value :=range names{
		fmt.Printf("value:%q \n",value)
	}
}