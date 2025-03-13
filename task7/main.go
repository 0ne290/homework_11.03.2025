package main

import (
	"fmt"
	"sync"
)

/* Исходный вариант
func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v// Первая же горутина будет бесконечно ждать.
		}(i)
	}

	// Т. к. дополнительные горутины бесконечно ждут, то и главная будет бесконечно ждать их завершения.
	wg.Wait()

	close(ch)
	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}*/

// Исправленный вариант
func main() {
	ch := make(chan int, 3)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		// Гонок данных не будет, т. к. i не замыкается, а передается в горутину в качестве параметра.
		// P. S. в последних версиях Go гонок данных не было бы, даже если бы i замыкалась, т. к., похоже, на
		// каждую итерацию цикла выделяется новая ячейка памяти под i (в этом можно убедиться с помощью
		// fmt.Println(&i))
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	wg.Wait()
	close(ch)
	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)// result: 5
}