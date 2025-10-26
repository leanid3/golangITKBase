// slice_utils.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomSlice генерирует слайс из 10 случайных целых чисел
func GenerateRandomSlice() []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(100) // числа от 0 до 99
	}
	return slice
}

// sliceExample возвращает новый слайс, содержащий только чётные числа
func sliceExample(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// addElements добавляет число в конец слайса и возвращает новый слайс
func addElements(slice []int, num int) []int {
	newSlice := make([]int, len(slice)+1)
	copy(newSlice, slice)
	newSlice[len(slice)] = num
	return newSlice
}

// copySlice возвращает независимую копию слайса
func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// removeElement удаляет элемент по индексу и возвращает новый слайс
func removeElement(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return nil, &IndexOutOfBoundsError{index: index, length: len(slice)}
	}
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice, nil
}

// Пользовательская ошибка для выхода за границы
type IndexOutOfBoundsError struct {
	index  int
	length int
}

func (e *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("индекс %d выходит за пределы длины слайса %d", e.index, e.length)
}
