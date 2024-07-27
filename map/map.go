package main

import "fmt"

func main() {
	// map хэш-таблица, словарь, ассоциативный массив
	m := map[int]string{1: "hello", 2: "world"}
	m[3] = "!"
	fmt.Println(m)

	elements := map[int64]bool{1: true, 2: true, 3: true}
	element, elementExists := elements[225]
	fmt.Println(element, elementExists)

	cache := make(map[string]struct{})
	cache["key"] = struct{}{}
	_, ok := cache["key"]
	fmt.Println(ok)

	engToRus := map[string]string{"hello": "привет", "world": "мир"}
	fmt.Println(engToRus)
	//удаление элемента
	delete(engToRus, "world")
	fmt.Println(engToRus)
}
