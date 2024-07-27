package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	nums[2] = 33
	//массив массивов
	matrix := [3][3]int{}
	fmt.Println(matrix)
	fmt.Println(len(matrix[1]))

	slice := []int{6, 3, 4, 4, 2}
	fmt.Println(slice[:3])
	fmt.Println(slice[3:])

	wordArr := []string{"hello"}
	wordArr = append(wordArr, "world")

	numsMake := make([]int, 3)
	fmt.Println(numsMake)
	modifySlice(numsMake)
	fmt.Println(numsMake)

	// массивы  передается по ссылке копирование слайсов осуществляется через copy
	// важно передавать len массива 2 пармаетром copy
	numsCp := make([]int, len(slice))
	copy(numsCp, slice)

	fmt.Println(numsCp)
	sortSlice(numsCp)
	fmt.Println(numsCp)
}

func modifySlice(nums []int) {
	nums[2] = 1
	nums = append(nums, 9)
}

func sortSlice(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}
