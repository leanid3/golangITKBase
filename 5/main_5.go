package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasInter, common := Intersect(a, b)
	fmt.Printf("Есть пересечение: %v\n", hasInter)
	fmt.Printf("Общие значения: %v\n", common)

}

func Intersect(slice1, slice2 []int) (bool, []int) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return false, []int{}
	}

	set2 := make(map[int]struct{}, len(slice2))
	for _, v := range slice2 {
		set2[v] = struct{}{}
	}

	seen := make(map[int]struct{})
	var result []int

	for _, v := range slice1 {
		if _, exists := set2[v]; exists {
			if _, alreadyAdded := seen[v]; !alreadyAdded {
				result = append(result, v)
				seen[v] = struct{}{}
			}
		}
	}

	hasIntersection := len(result) > 0
	return hasIntersection, result
}
