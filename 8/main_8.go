package main

import (
	"fmt"
	"time"
)

// #### Задание 8

// Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.

// * Напишите unit тесты к созданным функциям

func main() {
	wg := NewCustomWaitGroup()
	defer wg.Close()

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}

	fmt.Println("Ожидание завершения всех горутин...")
	wg.Wait()
	fmt.Println("Все горутины завершены!")
}
