package main

import "fmt"

func main() {
	s := "hey"
	// Конвертация строки в слайс байт описывается в коде явно
	bs := []byte(s)

	fmt.Println(s[0], s[1], s[2])                         // 104 101 121
	fmt.Println(string(s[0]), string(s[1]), string(s[2])) // h e y

	fmt.Println([]byte(s))  // [104 101 121]
	fmt.Println(string(bs)) // hey

	l := "Lorem ipsum dolor sit amet"
	// Обход строки
	for i := 0; i < len(l); i++ {
		fmt.Println(string(l[i]))
	}
}
