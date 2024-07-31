package main

import (
	"errors"
	"fmt"
)

// Функции с маленькой буквы используются только в рамках текущего пакета
func main() {
	result := multiply(6, 6)
	fmt.Println(result)

	divideRes, err := Divide(6, 6)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(divideRes)
	}

	PrintNums(1, 2, 3, 4, 5)

	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10}

	res := append(nums1, nums2...)
	fmt.Println(res)
}

func multiply(x int, y int) (res int) {
	res = x * y
	return
}

func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Can't devide to zero")
	}

	return x / y, nil
}

func PrintNums(nums ...int) {
	for _, n := range nums {
		fmt.Println(n)
	}
}
