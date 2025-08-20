package main

import "fmt"

// slices
// most use in go
// dynamic in go
func main (){

	// ininitilized nil

	var nums[]int

	fmt.Println(nums==nil)


	// slice can take  3 parameter  

	var slicenum = make([]int, 2,5)    // 2 is not a fixed size its initial size , and 5 is capacity  initialy they we change becuase  its dynamic automatically 

	fmt.Println(slicenum)
	fmt.Println(cap(slicenum))

	fmt.Println("After add elements ")

	slicenum = append(slicenum,1)
	slicenum = append(slicenum,2)
	slicenum = append(slicenum,3)
	slicenum = append(slicenum,4)
	slicenum = append(slicenum,5)
	slicenum = append(slicenum,6)
 

	fmt.Println(slicenum)
	fmt.Println(cap(slicenum))

	// another way to define 

	newslice := []int{}

	fmt.Println("-------------",newslice, cap(newslice), len(newslice))

	newslice = append(newslice,1)
	newslice = append(newslice,2)


	fmt.Println("-------------",newslice, cap(newslice), len(newslice))


	copyslice := []int{}
	anoterslice := make([]int, 2) // allocate space for 2 elements

	copyslice = append(copyslice, 10)
	copyslice = append(copyslice, 20)

	fmt.Println("This is copy slice:", copyslice)

	newtxt := copy(anoterslice, copyslice) // number of elements copied
	fmt.Println("Copied elements count:", newtxt)
	fmt.Println("Another slice:", anoterslice)



}