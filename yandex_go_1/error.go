package main

import "fmt"

func fooBar() string {
	return "bar"
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("делитель равен 0") // Возвращение 2 значений одновременно
	}
	return a / b, nil // Возвращение 2 значений одновременно
}

func foo() {
	// Механизм похожий на исключение, но ее нельзя типизировать
	// По умолчанию паника будет идти вверх по стеку и завершать все функции, пока не завершит функцию main
	// Не нужно использовать панику там, где можно просто возвратить ошибку и затем обработать её
	// Функция recover, которая позволяет восстановить выполнение программы в случае паники
	panic("unexpected!")
}

func main() {
	// defer механизм похожий на деструктор
	defer func() {
		if r := recover(); r != nil {
			// Обработка паники, в переменной r будет лежать строка "unexpected"
			fmt.Println("Recovered in f", r)
		}
	}()

	d, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("d = %d", d)
	}

	foo()
}
