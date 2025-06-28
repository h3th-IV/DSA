package main


func check_duplicate(nums []int) bool{
	empty_array := []int{}
	for _, element := range nums{
		for i:=0; i<len(empty_array); i++{
			if element == empty_array[i]{
				return true
			}
		}
		empty_array = append(empty_array, element)
	}
	return false
}

func map_check_duplicate(nums []int) bool{
	exist_map := make(map[int]bool)
	for _, element := range nums{
		if exist_map[element] {
			return true
		}
		exist_map[element] = true
	}
	return false
}