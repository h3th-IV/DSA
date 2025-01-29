package main

import (
	"fmt"
)

func main() {
	// test_arr := []float64{1, 2, 3, 4, 5, 5, 6, 7, 56, 9, 0, 89, 9, 19, 99, 101, 0, 1, 2, 3, 7, 3, 2, 8, 1, 5, 150, 0, 2, 4, 2, 6, 3, 100, 500, 9, 3, 6, 2, 6, 2, 9, -2}
	// max_num := maxNumArr(test_arr)
	// fmt.Printf("the hightest number in the provided array is: %v\n", max_num)

	arr := []int32{-4, 3, -9, 0, 4, 1}
	// plusMinus(arr)
	test_ll := new(LinkedList)
	fmt.Println("here")
	for _, v := range arr {
		test_ll.InsertNode(v)
	}
	test_ll.Print()
}

/*
Given an array of signed integers find the max integer
*/
func maxNumArr(num_arr []float64) float64 {
	max_num := num_arr[0]
	for _, i := range num_arr {
		if i > max_num {
			max_num = i
		}
	}
	return max_num
}

/*
Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with places after the decimal.
*/
func plusMinus(arr []int32) {
	// Write your code here
	arr_len := len(arr)
	posCount, negCount, zeroCount := 0, 0, 0
	for _, i := range arr {
		if i > 0 {
			posCount += 1
		} else if i < 0 {
			negCount += 1
		} else {
			zeroCount += 1
		}
	}
	posRatio := float64(posCount) / float64(arr_len)
	negRatio := float64(negCount) / float64(arr_len)
	zeroRatio := float64(zeroCount) / float64(arr_len)

	fmt.Printf("%.6f\n", posRatio)
	fmt.Printf("%.6f\n", negRatio)
	fmt.Printf("%.6f\n", zeroRatio)
}

/*
Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.
*/
func minMaxSum(arr []int32) {
}
