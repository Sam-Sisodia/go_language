package main

import "fmt"


func main(){
	fmt.Println("Variables in GO")
	var name  string= "Go"   // we must use this variable  otherwise it give error 

	fmt.Println("My name is",name)

	var adult bool = true
	fmt.Println(adult)

	// another way 
	o := "sajal"           // we cant override the value of name variable like python 
	fmt.Println("My name is",o)




}