package main

import "fmt"

func main(){

	menu := map[string]float64{
		"soup": 4.99,
		"pie":7.99,
		"salad":6.99,
		"coffee pudding": 3.55,
	}
	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	//looping in map

	for key,value:= range menu{
		fmt.Printf("key:%q value: %v \n",key,value)
	}
}