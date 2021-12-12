package main

import "fmt"

func main(){
	age:=45
	name:=[]string{"a","b","c","d","e","f","g"}
	fmt.Println(age>=30)
	fmt.Println(age<=40)
	fmt.Println(age==45)
	fmt.Println(age!=45)
	
	if age>=30{
		age=40
	}
	// fmt.Println(age)

	//continue
	for index, value:= range name{
		if index==1{
			fmt.Println(index)
			continue
			
			
		}
		fmt.Println("continue Statement: ",value)
	}
	//break
	for index, value:= range name{
		if index==1{
			fmt.Println(index)
			break
			
		}
		fmt.Println("break statement: ",value)
	}




}