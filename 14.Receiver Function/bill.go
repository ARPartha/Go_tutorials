package main

import "fmt"

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

// new bill

func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{
			"Tea":43.3333,
			"Coffee":75.423,
			"Cappacinno":75.5233,
		},
		tip:  0,
	}
	return b
}

// receiver function (formatting bill)

func (b bill) format() string{
	fs:= "Bill Breakdown: \n"
	var total float64=0

	//list items 

	for key,value:= range b.item{
		fs+=fmt.Sprintf("%v ...$%v \n",key+":",value)
		total+=value
	}
	//total
	fs+=fmt.Sprintf("%v ....$%0.3f","total:",total)
	return fs
}