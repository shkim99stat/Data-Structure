package main

import "fmt"

func Bin_search(ar []int, len, target int) int {
	first := 0
	last := len

	for first <= last {
		mid := int((first + last) / 2)

		if target == ar[mid] {
			return mid
		} else if target < ar[mid] {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	idx := Bin_search(arr, len(arr), 7)

	if idx == -1 {
		fmt.Println("Fail")
	} else {
		fmt.Println("index :", idx)
	}
}
