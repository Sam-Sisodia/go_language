package main


import "fmt"


var name = "sajal"

func main(){
	fmt.Println("CONSTANTS, only  use fixed values")
	const age  = 30
	

 

	fmt.Println(age,name)
  
	// multiple constants
	const (
		port = 5000
		host = "localhost"
		 
	)

  fmt.Println(port,host)
}