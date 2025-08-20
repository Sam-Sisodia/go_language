package main


import "fmt"

func main () {

	fmt.Println("Conditions in GO")
	age := 16

	if age >=18{
		fmt.Println("person is an adult")
	} else if age <= 16 {
		fmt.Println("Person is child ")
	}else{
		fmt.Println("Person is old")
	}


	// we can use logical opreator as well

	var  role = "admin"
	var hasPermissions = true 

	if role=="admin"  && hasPermissions{
		fmt.Println("Permission granted ")
	}else{
		fmt.Println("Acess denied ")
	}


	// we can define variable inside if constuct 

	if age:=20 ;  age>=18{
		fmt.Println("persopn is an adult",age)
	} else if age >= 12 {
		fmt.Println("Person is teenager",age)
	}


	// we dont have ternaru operator in Go 

}