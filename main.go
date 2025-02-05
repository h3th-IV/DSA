package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	channel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		PrintNumber(channel)
		wg.Done()
	}()

	fmt.Println("Running the main function...")
	for num := range channel {
		fmt.Println(num)
	}
	fmt.Println("running other go routines...")
	wg.Wait()
	fmt.Println("Done with main function")
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

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}

// 12 hr conversion to 24 hr
func timeConversion(time string) string {
	timex := strings.Split(time, ":")
	if len(timex) != 3 {
		log.Println("wrong time format")
	}
	hr, err := strconv.Atoi(timex[0])
	if err != nil {
		log.Println("err converting hr to int")
	}
	min := timex[1]
	secMer := timex[2]
	re := regexp.MustCompile(`[a-zA-Z]+`)
	sec := re.ReplaceAllString(secMer, "")
	if strings.Contains(secMer, "PM") && hr != 12 {
		hr += 12
	} else if strings.Contains(secMer, "AM") && hr == 12 {
		hr = 0
	}
	return fmt.Sprintf("%02d:%v:%v", hr, min, sec)
}

func compareTriplets(a []int32, b []int32) []int32 {
	results := [2]int32{}
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			results[0] += 1
		} else if a[i] < b[i] {
			results[1] += 1
		}
	}
	fmt.Println(results)
	return results[:]
}

func getDiagonalDifference(arr [][]int) uint {
	leng := len(arr)
	var (
		l_r int
		r_l int
	)
	for _, v := range arr {
		if len(v) != leng {
			log.Printf("the input matrix is not a square matrix")
		}
	}
	for i := 0; i < len(arr); i++ {
		l_r += arr[i][i]
		r_l += arr[i][len(arr)-1-i]
	}
	fmt.Println(uint(l_r - r_l))
	return uint(l_r - r_l)
}

func PrintStair(thread int) {
	for i := 1; i <= thread; i++ {
		for j := 1; j <= thread-i; j++ {
			fmt.Printf("  ")
		}
		for k := 1; k <= i; k++ {
			fmt.Printf("# ")
		}
		fmt.Println()
	}
}

func bDayCandles(arr []int) {
	var (
		total int
		high  int
	)
	for _, v := range arr {
		if v > high {
			high = v
		}
	}

	for _, v := range arr {
		if v == high {
			total += 1
		}
	}
	fmt.Println(total)
}

func repeatedString(s string, n int64) int64 {
	var acount int64
	for _, v := range s[:n] {
		if v == 'a' {
			acount += 1
		}
	}
	fullRep := n / int64(len(s))
	rem := n % int64(len(s))
	count := fullRep * acount

	for _, v := range s[:rem] {
		if v == 'a' {
			count++
		}
	}
	fmt.Println(count)
	return count
}

func jumpCloud(clouds []int) int {
	jumps := 0
	current := 0
	for current < len(clouds)-1 {
		if current+2 < len(clouds) && clouds[current+2] == 0 {
			current += 2
		} else {
			current += 1
		}
		jumps++
	}
	return jumps
}

func sortCards(cards []interface{}) []interface{} {
	cardOrder := map[string]int{
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"10":    10,
		"Jack":  11,
		"Queen": 12,
		"King":  13,
	}
}
