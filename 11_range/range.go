package main

import "fmt"



func  main() {
	fmt.Println("Range in Go we use  range in go  for iterating on stuctures ")
	nums := []int {10,20,30,}

	for i:=0 ; i< len(nums) ; i++ {
		fmt.Println("This is the i values ",nums[i])
	}

	// using range 
	sum :=0

	for index, num := range nums{
		fmt.Println("This using by  range :  ",index,num)       // _ is index 
		sum = sum + num

	}


	fmt.Println("This is  the sun ",sum)



	// range in maps

	m := map[string]string{"firstname": "Sajal", "last_name":"Sisodia"}

	for key , val := range m{

		fmt.Println("This is info",key,val)

	}
	 


}