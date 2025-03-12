package main

import (
	"fmt"
)

/* Исходный вариант
func main() {
	Строки неизменяемы
	str := "Привет"
	str[2] = 'e'
	fmt.Println(str)
}*/

// Исправленный вариант
func main() {
	str := []rune("Привет")
	str[2] = 'e'
	fmt.Println(string(str))
}
