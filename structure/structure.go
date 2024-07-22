package main

import "fmt"

type Person struct {
	// N поле публично
	Name string
	// a поле приватно: можно обращаться только внутри текущего пакета
	age int
}

// w структура приватна: можно инициализировать только внутри текущего пакета
type wallet struct {
	id          string
	moneyAmount float64
}

func main() {
	p := Person{Name: "John", age: 25}
	fmt.Println(p.Name)

	x := Person{}
	x.Name = "Jane"
	fmt.Println(x)

	y := Person{}
	fmt.Println(y)

	w := wallet{id: "123", moneyAmount: 100}
	fmt.Println(w)
}
