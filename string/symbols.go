package main

import "fmt"

func main() {
	emoji := []rune("Hello 😀")
	// rune — это алиас к int32
	// Каждая руна представляет собой код символа стандарта Юникод
	for i := 0; i < len(emoji); i++ {
		fmt.Println(emoji[i], string(emoji[i]))
	}

	//s := "hey 😀"
	//rs := []rune([]byte(s)) // cannot convert ([]byte)(s) (type []byte) to type []rune
	//bs := []byte([]rune(s)) // cannot convert ([]rune)(s) (type []rune) to type []byte

	for _, symbol := range emoji {
		fmt.Println(symbol, string(symbol))
	}
}
