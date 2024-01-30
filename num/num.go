package main

import "fmt"

var x int = 10
var y int = 5

func main() {
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x / y)
	fmt.Println(x * y)

	// Любые операции осуществляются только над числами одинакового типа
	x := 5.05
 	y := 10
	// invalid operation: x + y (mismatched types float64 and int)
	// fmt.Println(x + y)

	fmt.Println(x + float64(y))
	// Отбрасываются значения после запятой
	fmt.Println(int64(x))
	// cannot convert -5 (untyped int constant) to type uint
	// fmt.Println(uint(-5))
}
