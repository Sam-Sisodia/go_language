package  main

import "fmt"

// variadic_func   take n number of argumnets 


func sum(nums... int) int {
	total :=0

	for _ , num := range nums {
		total = total + num
	}

	return  total

}


func digit_sum(digits... int)  int {
	t := 0 

	for _, x := range digits{
		t = t+x
	}

	return  t 

}

func main() {
	res := sum(2,3,4,5,66)
	fmt.Println("This is sum ", res)

	fmt.Println(1,2)

	digits :=[]int {3,4,5,6}

	// pass directly 

	res2 := digit_sum(digits...)
	fmt.Println(res2)


}


