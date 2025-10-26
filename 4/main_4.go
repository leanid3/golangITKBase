package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := Difference(slice1, slice2)
	fmt.Println("Элементы только в первом слайсе:", result)
}

func Difference(slice1, slice2 []string) []string {
	// Преобразуем slice2 в множество (map) для O(1) поиска
	set := make(map[string]struct{}, len(slice2))
	for _, s := range slice2 {
		set[s] = struct{}{}
	}

	var result []string
	for _, s := range slice1 {
		if _, exists := set[s]; !exists {
			result = append(result, s)
		}
	}
	return result
}
