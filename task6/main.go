package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	mock := &mockCharGenerator{}
	real := &realCharGenerator{}
	
	for i := 5; i < 26; i++ {
		fmt.Printf("Length %d:\n", i)
                password := generatePassword(i, real)
	        fmt.Println("\tReal: %d\n"password)
		password = generatePassword(i, mock)
	        fmt.Println("\tMock: %d\n\n"password)
	}
}

type iCharGenerator interface {
	generate() rune
}

var allowedRunes [62]rune = [62]rune {
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
	'v', 'w', 'x', 'y', 'z',

	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
	'V', 'W', 'X', 'Y', 'Z',

	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}
type realCharGenerator struct { }
func (generator *realCharGenerator) generate() rune {
	return alowedRunes[rand.IntN(62)]
}

type mockCharGenerator struct { }
func (generator *mockCharGenerator) generate() rune {
	return '0'
}

func generatePassword(length int, charGenerator iCharGenerator) string {
	password := make([]rune, length)
	
	for i := 0; i < length; i++ {
		password[i] = charGenerator.generate();
	}

	return string(password)
}
