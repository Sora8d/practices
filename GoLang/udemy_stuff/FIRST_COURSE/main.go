package main

import "fmt"

//So here i want to see the realtionship between slices/arrays and pointers. With arrays you get a copy, with slices you get the pointer; changing from array to slice in the function doesnt affect behaviour
type coolArray []int

func main() {
	array := [2]int{1, 2}
	fmt.Println(array)
	change_without_pointer(array)
	fmt.Println(array)
	//	change_with_pointer(&array)
	//	fmt.Println(array)

	fmt.Println(" ")

	freshArray := coolArray{1, 2}
	fmt.Println(freshArray)
	freshArray.change_with_receiver_without_pointer()
	fmt.Println(freshArray)
	//	freshArray.change_with_receiver_with_pointer()
	//	fmt.Println(freshArray)
}

func change_without_pointer(arr [2]int) {
	sliceOfArr := arr[1:2]
	sliceOfArr[0] = 12
	fmt.Println(sliceOfArr)
}

/*
func change_with_pointer(arr *[]int) {
	arr[1] = 12
}
*/
func (CA coolArray) change_with_receiver_without_pointer() {
	CA[1] = 12
}

/*
func (CA *coolArray) change_with_receiver_with_pointer() {
	CA[1] = 12
}
*/
