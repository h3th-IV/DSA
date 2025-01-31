package main

import (
	"fmt"
	"math"
)

func main() {
	// test_arr := []float64{1, 2, 3, 4, 5, 5, 6, 7, 56, 9, 0, 89, 9, 19, 99, 101, 0, 1, 2, 3, 7, 3, 2, 8, 1, 5, 150, 0, 2, 4, 2, 6, 3, 100, 500, 9, 3, 6, 2, 6, 2, 9, -2}
	// max_num := maxNumArr(test_arr)
	// fmt.Printf("the hightest number in the provided array is: %v\n", max_num)

	// arr := []int32{-4, 3, -9, 0, 4, 1}
	// pos_arr := []int32{1, 2, 3, 4, 5}
	// plusMinus(arr)
	// minMaxSum(pos_arr)

	// ans := checkPalindrome("kayak")
	// fmt.Println(ans)

	// fmt.Println(reverseNum(90001))
	fmt.Println(isPalindrome("kayak"))
}

// reverse a string
func ReverseString(data string) {
	for i := 0; i < len(data); i++ {
		fmt.Printf("%c", data[len(data)-i-1])
	}
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
	max, min := math.Inf(-1), math.Inf(1)
	var totalSum float64
	for _, v := range arr {
		totalSum += float64(v)
	}

	for _, v := range arr {
		sum_without_current_value := totalSum - float64(v)
		if sum_without_current_value < min {
			min = sum_without_current_value
		}
		if sum_without_current_value > max {
			max = sum_without_current_value
		}
	}
	fmt.Println(min, max)
}

func reversal(data string) string {
	lengh := len(data)
	reversedData := ""
	fmt.Println("")
	for i, _ := range data {
		reversedData += string(data[lengh-i-1])
	}
	return reversedData
}

// palindrome a word that can
func checkPalindrome(data string) bool {
	reversedString := reversal(data)
	word_len := len(data)
	for i := 0; i < word_len; i++ {
		if data[i] != reversedString[i] {
			return false
		}
	}
	return true
}

// reverse an integer
func reverseNum(num int) int {
	reversed_num := 0
	for num != 0 {
		last_digit := num % 10
		reversed_num = reversed_num*10 + last_digit
		num = num / 10
	}
	return reversed_num
}

func isPalindrome(data string) bool {
	word_len := len(data)
	var start, end = 0, word_len - 1
	for start < end {
		if data[start] != data[end] {
			return false
		}
		start++
		end--
	}
	return true
}
