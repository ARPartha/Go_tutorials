package main

import "fmt"


func main(){
	var name string= "partha" // one type of declaration
	var nameOne= "tonni" // another type of declaration
	var nameTwo string // another type of declaration
	nameThree:="In Sha ALLAH sob hobe" //short form of declaring variable  
	var check bool = true
	nameTwo="Manju"
	
	if(check==true){
		fmt.Println(nameThree)
		fmt.Println(nameTwo)
		fmt.Println(nameOne)
		fmt.Println(name)
	}

	var age int =20 
	fmt.Println(age)


	//bits and memory 

	var numOne int8=25 //8 bit int 
	var numTwo int16= 25543 // 16 bit int 
	var numThree int32 =343434343 // 32 bit int
	var numFour int64=3343434343434343434 //64 bit int 
	var unsigned uint= 121323232 // all positive
	fmt.Println(numOne, numTwo, numThree,numFour,unsigned)

	//floats

	var float float64= 34.343434343434343

	fmt.Println(float)

}
