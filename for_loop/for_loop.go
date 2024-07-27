package main

import "fmt"

func main() {
	nums := make([]int, 0)

	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}

	fmt.Println(nums)

	// Аналог while
	j := 0
	numsWhile := make([]int, 0)

	for j < 11 {
		numsWhile = append(numsWhile, j)
		j++
	}

	fmt.Println(numsWhile)

	numsEven := make([]int, 0)
	for k := 0; k < 12; k++ {
		if k%2 != 0 {
			continue
		}
		numsEven = append(numsEven, k)
	}

	fmt.Println(numsEven)

	names := []string{"John", "Paul", "George", "Ringo"}
	// ranges аналог foreach
	for i, name := range names {
		fmt.Println("Hello ", name, " at index ", i)
	}

	for i := range names {
		fmt.Println("index = ", i)
	}
	// Можно пропустить первую переменную, это можно сделать с помощью _
	for _, name := range names {
		fmt.Println("Hello ", name)
	}
}
