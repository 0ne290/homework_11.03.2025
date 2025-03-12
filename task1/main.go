package main

import (
	"fmt"
)

/* Исходный вариант
func printNumber(ptrToNumber interface{}) {
	if ptrToNumber != nil {
		fmt.Println(*ptrToNumber.(*int))
	} else {
		fmt.Println("nil")
	}
}*/

// Исправленный вариант
func printNumber(ptrToNumber interface{}) {
	if ptrToNumber == nil {
		fmt.Println("nil")

		return
	}

	ptrToInt, ok := ptrToNumber.(*int)
	if !ok {
		fmt.Println("not pointer to int")

		return
	}

	fmt.Println(*ptrToInt)
}

func main() {
	number := 34
	printNumber(nil)
	printNumber(4)
	printNumber(&number)
}
