package main

import (
	"fmt"
)

/* Исходный вариант
func main() {
	str := "Привет"
	str[2] = 'e'// Строки неизменяемы
	fmt.Println(str)
}*/

// Исправленный вариант
func main() {
	str := []rune("Привет")
	str[2] = 'e'
	fmt.Println(string(str))
}
