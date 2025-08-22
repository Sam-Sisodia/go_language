package main

import "fmt"



// whne function return inner function
func counter() func() int {
	var count int = 0 
	return  func() int {
		count +=1
		return  count
	}
}


func main(){
	fmt.Println("Closures in Go ")
	increment := counter()
	fmt.Println("This is closure : ", increment(),increment())



	


}