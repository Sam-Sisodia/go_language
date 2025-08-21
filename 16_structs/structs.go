package main

import (
	"fmt"
	"time"
)



type order  struct {

	id string
	amount float32
	status string
	createdAt  time.Time // neno second
}

func main () {

	fmt.Println("Structs in Go")
  
	 
	myorder := order {
		id : "1",
		amount: 50.0,
		status: "received",

	}
	myorder.createdAt = time.Now()

	fmt.Println("myorder stuct ",myorder)

	//get the fields 
	fmt.Println("mystore amount  ",myorder.amount)


	myorder2 := order {
		id : "2",
		amount: 10.0,
		status: "received",

	}

	fmt.Println("mystore amount  ",myorder2)


 

}  