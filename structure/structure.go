package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string // N поле публично
	age    int    // a поле приватно: можно обращаться только внутри текущего пакета
	wallet wallet
}

// w структура приватна: можно инициализировать только внутри текущего пакета
type wallet struct {
	id          string
	moneyAmount float64
}

// User json используется для названий полей при сериализации/десериализации структуры в json и обратно
type User struct {
	Id        int64  `json:"id" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
}

func main() {
	p := Person{Name: "John", age: 25, wallet: wallet{id: "2", moneyAmount: 100}}
	fmt.Println(p.Name)
	fmt.Println(p.age)
	fmt.Println(p.wallet)

	x := Person{}
	x.Name = "Jane"
	fmt.Println(x)

	w := wallet{id: "123", moneyAmount: 100}
	fmt.Println(w)

	u := User{Id: 1, Email: "jane@gmail.com", FirstName: "Jane"}
	// {"id":1,"email":"jane@gmail.com","first_name":"Jane"}
	userSerialize(u)
}

func userSerialize(u User) {
	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
}
