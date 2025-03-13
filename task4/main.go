package main

import (
	"fmt"
)

/* Исходный вариант
func main() {
	// Дефолтное значение типа map - nil. Соответственно, мапу всегда нужно инициализировать с помощью make (также
	// желательно выделять память под мапу с запасом, т. к. при расширении мапы и ее неравномерном заполнении
	// скорость мапы снижается
	var m map[string]int

	m = make(map[string]int, 128)
	for _, word := range []string{"hello", "world", "from", "the", "best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}*/

// Исправленный вариант
func main() {
	m := make(map[string]int, 128)

	for _, word := range []string{"hello", "world", "from", "the", "best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
