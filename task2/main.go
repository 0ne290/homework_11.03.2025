package main

import (
	"fmt"
)

// Тип реализует стандартный интерфейс ошибок.
type MyError struct{}
func (MyError) Error() string {
	return "MyError!"
}

// Функция принимает любую структуру, реализующую стандартный интерфейс ошибок.
func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	// Т. к. значение nil является допустимым для любого типа интерфейсов, его можно передать в функции,
	// принимающие интерфейсы - именно это тут и происходит. Вывод: Error: <nil>
	var err *MyError
	errorHandler(err)

	// Передаем нашу кастомную ошибку. Вывод: Error: MyError!
	err = &MyError{}
	errorHandler(err)
}
	