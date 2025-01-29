package main

import "fmt"

//using make in go to create slices, maps and channels

func Maker() {
	//slice with lenght 5 and capacity of 10
	new_slice := make([]int, 5, 10)
	//map
	my_map := make(map[int]string, 10) //maps are unordered collection of key value pair

	//also define maps like this
	ano_map := map[int]string{}
	ano_map[2] = "threader"

	//or
	var text_mapper = map[string]interface{}{}
	text_mapper["map"] = my_map
	text_mapper["slice"] = new_slice
	fmt.Println(text_mapper)
}
