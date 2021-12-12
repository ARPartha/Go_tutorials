package main

import (
	"fmt"
	"sort"
)

func main(){

	ages:=[]int {1,25,3,343,6,7,1,2,35,77,5}
	//sort package
	sort.Ints(ages)
	fmt.Println(ages)
	index:=sort.SearchInts(ages,343)
	fmt.Println(index)
}