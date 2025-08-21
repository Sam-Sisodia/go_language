package main

import (
	"fmt"

)

func main () {
	fmt.Println("Maps in go ")

	// creating maps
	m := make(map[string]string)  // first string is key  data type and  the second is value data type
	fmt.Println(m)


	// setting an element in  

	m["sajal"] = "sisodia"
	m["jack"] = "john"

	fmt.Println(m)

	// id key valus is does not exits in its return o  or empty 


	fmt.Println(len(m))

	//delete element from  hash 

	delete(m,"jack")
	fmt.Println("After delete ", m)
	clear(m) // it clear the  map
  

	fmt.Println("clear ", m)

	// without make  function 
  
	user := map[string]int {"price": 40, "phone":9008}


	fmt.Println("User details ", user)



	// check the perticualr  key exits  or not 

	v, ok := user["phone"]
	fmt.Println(v)

	if ok{ 
		fmt.Println("Key found")

	}else {
		fmt.Println("not  found")
	}


	//check 2 maps are equal or not      

	// fmt.Prinrlb(maps.Equal(m1,m2))



}