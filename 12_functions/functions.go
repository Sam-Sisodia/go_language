package main

import "fmt"



func add(a int,b int ) int {

	return a+b
}




func  get_languages ()(string,string,string)  {
	return "go", "python", "JS"
}


//  function as parameter

func process (fn func(a int)int  ){

	fn(1)



}




func main(){
	result := add(3,6)
	fmt.Println("This is no", result)
	lan1,lan2,lan3 := get_languages()
	fmt.Println(lan1,lan2,lan3)




  fn:= func(a int) int {
		return 2
	}

	process(fn)
 







}























// fmt.Println("Functions in Go ")

	// num :=121
	// rev := 0

	// for num > 0 {
	// 	digit := num % 10

	// 	rev = rev * 10 + digit
	// 	num = num / 10

	// }
	// fmt.Println(rev)

// }