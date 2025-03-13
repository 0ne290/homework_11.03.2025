package main

import (
	"fmt"
	"sync"
)

func merge(channels []chan int) chan int {
	var wg sync.WaitGroup
	var once sync.Once
	ret := make(chan int)

	wg.Add(len(channels))
	for _, channel := range channels {
		go func(channel chan int) {
			defer wg.Done()

			for number := range channel {
				ret <- number
			}

			once.Do(func() {
				for _, channel := range channels {
					defer func() {
						recover()
					}()

					close(channel)
				}
			})
		}(channel)
	}

	go func() {
		wg.Wait()

		close(ret)
	}()

	return ret
}

func main() {
	var wg sync.WaitGroup
	channels := []chan int{make(chan int), make(chan int), make(chan int)}
	
	wg.Add(len(channels))
	for i, channel := range channels {
		go func(channel chan int) {
			defer wg.Done()

			for _, number := range []int{i + 1, i + 4, i + 7} {
				channel <- number
			}
		}(channel)
	}
	
	// Закрываем только третий канал. Первый и второй закроются автоматически.
	go func() {
		wg.Wait()

		close(channels[2])
	}()

	mergedChannel := merge(channels)
	for v := range mergedChannel {
		fmt.Println(v)
	}
}
