package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

// new bill

func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}
	return b
}

// receiver function (formatting bill)

func (b *bill) format() string{
	fs:= "Bill Breakdown: \n"
	var total float64=0
	
	//list items 

	for key,value:= range b.item{
		fs+=fmt.Sprintf("%-25v ...$%v \n",key+":",value)
		total+=value
	}
	//tip
	fs+=fmt.Sprintf("%-25v ...$%0.1f\n","Tip:",b.tip)
	//total
	fs+=fmt.Sprintf("%-25v ...$%0.1f","Total:",total+b.tip)
	return fs
}
// add tip to the bill
func (b *bill) updateTip(tip float64){
	  b.tip=tip
}
// add item to the bill
func (b *bill) addItem(item string,price float64){
	b.item[item]=price
}
func (b *bill) saveBill(){
	data:= []byte(b.format())

	err:=os.WriteFile("/"+b.name+".txt",data,0644)

	if err!=nil{
		panic(err)
	}else{
		fmt.Println("bil is saved")
	}
}