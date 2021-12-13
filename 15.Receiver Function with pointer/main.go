package main

import "fmt"

func main() {
	mybill:= newBill("mario's bill")

	mybill.addItem("Cake",45.5)
	mybill.addItem("Coffee",25.5)
	mybill.updateTip(20)
	fmt.Println(mybill.format())


}