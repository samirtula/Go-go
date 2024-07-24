package main

import "fmt"

const okStatus = 200

const (
	notFoundStatus = 404
	redirectStatus = 302
)

const (
	num     = 20
	str     = "hey"
	isValid = true
)

const (
	InfoStatus  = 100 // I публичная константа, которую можно использовать во внешних пакетах
	clientError = 404 // c приватная константа, доступная только в рамках текущего пакета
	serverError = 500
)

// нельзя объявить структуру как константу

const (
	// Для последовательных числовых констант следует использовать идентификатор
	//iota, который присвоит для списка чисел значения от его текущей позиции
	a = iota
	b = 42
	c = iota
	d
)

func main() {
	fmt.Println(okStatus)
	fmt.Println(notFoundStatus)
	fmt.Println(redirectStatus)
}
