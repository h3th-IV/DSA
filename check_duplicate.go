package main

import "fmt"

func check_duplicate(nums []int) bool{
	empty_array := []int{}
	fmt.Println(empty_array)
	for _, element := range nums{
		for i:=0; i<len(empty_array); i++{
			if element == empty_array[i]{
				return true
			}
		}
		empty_array = append(empty_array, element)
	fmt.Println(empty_array)
	}
	return false
}