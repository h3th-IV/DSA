package main

import "fmt"

var (
	my_arr   [5]int //an array
	my_slice []int  //a slice; has an underlying array
)

func Arrayer() {
	our_arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("our_arr has a lentgh of %v", len(our_arr))
}
func Slicer() {
	our_slice := []int{}
	our_slice = append(our_slice, 5)
	fmt.Printf("our_slice has a length of %v and a capacity of %v\n", len(our_slice), cap(our_slice))

	//a slice is copied from an unerlying array
	our_arr := [10]int{}
	new_slice := our_arr[:8]
	//len will the number of element in the slice and while the cap will the be len(number of element of the underlying array)
	fmt.Printf("length of our_arr is %v", len(our_arr))
	fmt.Printf("length of new_slice is %v and the cap of new_slice is %v", len(new_slice), cap(new_slice))
}
