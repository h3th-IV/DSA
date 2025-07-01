package main

//find the product of every element in array except itself: i.e
//array is [e, t, r, d], return [trd, erd, etd, etr]

func product_except_self(nums []int) []int {
	n := len(nums)
	output := make([]int, n)

	// Initialize all values to 1
	for i := 0; i < n; i++ {
		output[i] = 1
	}

	// Left pass
	leftProduct := 1
	for i := 0; i < n; i++ {
		output[i] = leftProduct
		leftProduct *= nums[i]
	}

	// Right pass
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return output
}
