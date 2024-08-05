package main

import "fmt"

// Dog Структура в которую добавляется метод
type Dog struct {
	isBarked bool
}

// Bark указатель * чтобы можно было изменить парметр
func (d *Dog) Bark() {
	fmt.Println("Meow!")
	d.isBarked = true

}

func main() {
	d := Dog{}
	d.Bark()
	fmt.Println(d.isBarked)
}
