package main

import "fmt"



// for only we have for loop , we use for loop for while loop as well 
func main(){

	// while loop  in go 

	i:=1

	for i <=3 {
		fmt.Println(i)
		i = i+1
	}



	// infinite loop 
  
	// for {
	// 	fmt.Print("1")

	// }

	//real for loop in go 
	for i:=0 ; i <10 ; i++{
		fmt.Println(i)
	

	}
  
	fmt.Println("----------------------------------------------------------")


	// we can use continue and break like python  , it perform same  in go 

	//range in go 

	for i:= range 10{
		fmt.Println(i)
	}




}  