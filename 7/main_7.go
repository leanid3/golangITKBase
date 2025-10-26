package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 <- 3
	ch2 <- 4
	close(ch2)

	ch3 <- 5
	close(ch3)

	// Сливаем
	merged := MergeChannels(ch1, ch2, ch3)

	fmt.Println("Слияние каналов:")
	for val := range merged {
		fmt.Println(val)
	}
}

func MergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(channels))
	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
