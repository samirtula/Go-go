package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	nums[2] = 33
	matrix := [3][3]int{} //массив массивов

	fmt.Println(nums)
	fmt.Println(nums[2])
	fmt.Println(nums[3])

	fmt.Println(matrix)
	fmt.Println(len(matrix[1]))
}
