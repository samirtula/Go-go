package main

import "fmt"

type User struct {
	email    string
	password string
}

// При объявлении указываем,
// Передача по ссылке
// Что переменная должна быть указателем.
// Для этого ставим звездочку * перед типом данных

func fillUserData(u *User, email string, pass string) {
	u.email = email
	u.password = pass
}

func fillMap(m map[string]int) {
	m["random"] = 1
}

func main() {
	u := User{}
	// передаем указатель с помощью амперсанда
	// & перед переменной
	fillUserData(&u, "go@gmail.com", "qwerty")
	fmt.Printf("points on func call %+v\n", u)

	up := &User{}
	fillUserData(up, "test@test.com", "qwerty")

	m := map[string]int{}
	fmt.Println(m)
	fillMap(m)
	fmt.Println(m)
}
