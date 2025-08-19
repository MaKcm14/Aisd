package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 2, 3, 4, 6, 7, 8, 8, 9, 10}

	fmt.Println(len(slice))
	fmt.Println(slices.BinarySearch(slice, 11))
	fmt.Println(slices.BinarySearch(slice, 8))
}
