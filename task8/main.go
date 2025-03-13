package main

/* Исходный вариант
func main() {
	ch := make(chan bool)
	ch <- true// Дедлок
	go func() {
		<-ch
	}()
	ch <- true
}*/

// Исправленный вариант
func main() {
	ch := make(chan bool)
	go func() {
		<-ch
	}()
	ch <- true
	close(ch)
}
