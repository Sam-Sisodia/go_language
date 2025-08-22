package main

import "fmt"


func changenum(num  *int ) {

  *num = 10
	fmt.Println("In change num ",*num)
	

}

func main() {
	fmt.Println("Pointer in  GO")
	num := 1 
	//changenum(num)  // pass by value the main num values not change 

	// pass by reference 
	changenum(&num)



	fmt.Println("After change  ", num)
}