package main

import "time"

func PrintNumber(channel chan int) {
	for i := 0; i <= 10; i++ {
		channel <- i
		time.Sleep(500 * time.Microsecond)
	}
	close(channel)
}
