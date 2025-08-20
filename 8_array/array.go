package main

import "fmt"


func main() {
	var nums [4] int
	fmt.Println("the  len of array is ", len((nums)))


	// enter the values in array 

	var new_array[4] int 
	new_array[0] = 10 

	fmt.Println("This is the array element ", new_array[0],new_array)


	// it adds  by defult values in array depends  on  data type 

	var strarray[4] string

	strarray[0] = "Go"
	fmt.Println("This is dtring default array ", strarray[0],strarray)



	// assign element at  declare time

	decarray := [3]int{3,4,5}
	fmt.Println(decarray)

  // 2 d array
	newarray := [2][2]int{{3,4,},{5,6}}
	fmt.Println(newarray)



	//    use if fixed size , memory optimize  

	



	
}