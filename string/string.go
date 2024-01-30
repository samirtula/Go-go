package main

import "fmt"

func main() {
	q := `
		SELECT * 
		FROM person
		WHERE age > 18
	`
	fmt.Println(q)

	h := "Hello" + " " + "world"
	fmt.Println(h)
	
	fmt.Println("привет" == "привет")
	fmt.Println("golang" > "go")
	fmt.Println("golang" > "lang")
	fmt.Println("go" > "foobar")
	
	fmt.Println(len("go"))
	
	// Количество байт
	fmt.Println(len("го"))

	username := "Alesha"
	greeting := fmt.Sprintf("hello %s", username)

	fmt.Println(greeting)
}