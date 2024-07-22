package main

import (
	"fmt"
	"strings"
)

func main() {
	n := "Mr.Mickey"
	fmt.Println(n + " is a " + statusByName(n))

	n = "Mrs. Mini"
	fmt.Println(n + " is a " + statusByName(n))

	n = "Karl"
	fmt.Println(n + " is a " + statusByName(n))
}

func statusByName(name string) string {
	if strings.HasPrefix(name, "Mr.") {
		return "married man"
	} else if strings.HasPrefix(name, "Mrs.") {
		return "married woman"
	}

	return "single person"
}