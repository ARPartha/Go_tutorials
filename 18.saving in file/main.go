package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	//switch case 
	switch strings.ToLower(opt){
	case "a":
		name,_:=getInput("What's the item: ",reader)
		price,_:=getInput("Item price: ",reader)

		p,err:=strconv.ParseFloat(price,64)

		if err!=nil{
			fmt.Println("price must be a number!")
			promptOptions(b)
		}else{
			b.addItem(name,p)
			promptOptions(b)
		}
		fmt.Println(name,price)
	case "s":
		b.saveBill()
		fmt.Println("You saved the file",b.name)
	case "t":
		tip,_:= getInput("Enter tip ammount:",reader)
		t,err:=strconv.ParseFloat(tip,64)

		if err!=nil{
			fmt.Println("Tip must be a number!")
			promptOptions(b)
		}else{
			b.updateTip(t)
			promptOptions(b)
		}
		fmt.Println(tip)
	default:
		fmt.Println("Invalid input")
		promptOptions(b)
	}

}


func main() {
	mybill:= createBill()
	promptOptions(mybill)

}