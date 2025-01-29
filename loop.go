package main

import "fmt"

func Looper() {
	//iteration loops
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//using range
	for i, v := range [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		fmt.Printf("element at position %d is %d\n", i, v)
	}

	//while loop
	j := 0
	for j <= 10 {
		fmt.Printf("%d", j)
	}
	//infinite loop
	for {
		fmt.Println("something")
		break
	}
}
