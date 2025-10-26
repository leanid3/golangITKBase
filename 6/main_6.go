package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := RandomGenerator()

	fmt.Println("Первые 5 случайных чисел:")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func RandomGenerator() <-chan int {
	ch := make(chan int)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		defer close(ch)
		for {
			ch <- r.Int()
		}
	}()

	return ch
}
