package main

import "fmt"

type Printer interface {
	Print()
}

type User struct {
	email string
}

// Print структура User имеет метод Print, как в интерфейсе Printer
// Следовательно, во время компиляции запишется связь между User и Printer
func (u *User) Print() {
	fmt.Println("My email is", u.email)
}

// TestPrint функция принимает как аргумент интерфейс Printer
func TestPrint(p Printer) {
	p.Print()
}

func main() {
	TestPrint(&User{email: "go@gmail.com"})
}
