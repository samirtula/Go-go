package main

import (
	"errors"
	"fmt"
)

func multiply(x int, y int) (res int) {
	res  = x * y
	return
}

func Divide(x, y int)  (int, error) {
	if y == 0 {
		return 0, errors.New("Can't devide to zero")
	}

	return x / y, nil
}

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
	
}