package main

import (
	"errors"
	"fmt"
)

type TimeoutErr struct {
	msg string
}

// структура TimeoutErr реализует интерфейс error
// и может быть использована как обычная ошибка
func (e *TimeoutErr) Error() string {
	return e.msg
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}

	if len([]rune(name)) > 100 {
		return errors.New("name is too long")
	}

	return nil
}

func main() {

	n := ""
	err := validateName(n)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("name is valid")
	}
}
