package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Lorem ipsum dolor sit amet"
	fmt.Println(strings.Contains(str, "L"))

	splitStr := strings.Split(str, "")
	fmt.Println(splitStr)

	joinedStr := strings.Join([]string{"Lorem", "Lorem"}, "")
	fmt.Println(joinedStr)

	// Билдер
	sb := strings.Builder{}
	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("World")

	strRes := sb.String()
	fmt.Println(strRes)
}
