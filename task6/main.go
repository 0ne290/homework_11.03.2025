package main

import (
	"fmt"
	"math/rand/v2"
)

/* Неправильный вариант потому что невозможно протестировать
func main() {
	Строки неизменяемы
	str := "Привет"
	str[2] = 'e'
	fmt.Println(str)
}*/

// Правильный вариант

type iCharGenerator interface {
	generate() rune
}

// "Боевая" реализация
const alowedRunes [72]rune = [72]rune {
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
	'v', 'w', 'x', 'y', 'z',

	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
	'V', 'W', 'X', 'Y', 'Z',

	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}
type realCharGenerator struct { }
func (generator *realCharGenerator) generate() rune {
	return alowedRunes[rand.IntN(72)]
}

// "Моковая" реализация для тестов
type realCharGenerator struct {

}
func 

func generatePassword(length int, charGenerator iCharGenerator) {
	str := []rune("Привет")
	str[2] = 'e'
	fmt.Println(string(str))
}
