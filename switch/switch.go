package main

import "fmt"

func main() {
	x := 10

	switch x {
	case 10:
		fmt.Println("x is 10")
		// ключевое слово fallthrough позволяет нам выполнить следующий блок case без проверки его условия.
		fallthrough
	case 20:
		fmt.Println("x is 20")
	default:
		fmt.Println("x is default")
	}
}
