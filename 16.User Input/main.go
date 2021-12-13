package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func getInput(message string, r *bufio.Reader) (string,error){
	fmt.Print(message)
	input,err:=r.ReadString('\n')

	return strings.TrimSpace(input),err
}


func createBill() bill{
	reader:= bufio.NewReader(os.Stdin)

	// fmt.Print("Create a New Bill name: ")
	
	// name=strings.TrimSpace(name)
	name,_:=getInput("Create a New Bill name: ", reader)
	b:=newBill(name)
	fmt.Println("New Bill Created for: ",b.name)
	// fmt.Println("New Bill Created ")
	return b
}


func promptOptions(b bill){
	reader:= bufio.NewReader(os.Stdin)
	opt,_:=getInput("Choose Options (a-add bill ,s-save item , t-add tip): ", reader)
	
	fmt.Println(opt)

}


func main() {
	mybill:= createBill()
	promptOptions(mybill)

}