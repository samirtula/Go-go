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
	// удаление элемента
	delete(engToRus, "world")
	fmt.Println(engToRus)

	// обход map
	idToName := map[int64]string{1: "Alex", 2: "Bob", 3: "Chris"}

	for id, name := range idToName {
		fmt.Println("id: ", id, "name: ", name)
	}

	numExistence := make(map[int]bool)

	// порядок ключей в мапе рандомизирован
	for i := 0; i < 10; i++ {
		numExistence[i] = true
	}

	//8 1 2 3 6 7 9 ...
	for num := range numExistence {
		fmt.Println(num)
	}
}
