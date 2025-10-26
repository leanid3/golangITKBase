package main

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name         string
		slice1       []int
		slice2       []int
		wantHasInter bool
		wantResult   []int
	}{
		{
			name:         "Пример из задания",
			slice1:       []int{65, 3, 58, 678, 64},
			slice2:       []int{64, 2, 3, 43},
			wantHasInter: true,
			wantResult:   []int{3, 64},
		},
		{
			name:         "Нет пересечений",
			slice1:       []int{1, 2, 3},
			slice2:       []int{4, 5, 6},
			wantHasInter: false,
			wantResult:   []int{},
		},
		{
			name:         "Полное совпадение",
			slice1:       []int{10, 20},
			slice2:       []int{20, 10},
			wantHasInter: true,
			wantResult:   []int{10, 20},
		},
		{
			name:         "Один общий элемент",
			slice1:       []int{5},
			slice2:       []int{5},
			wantHasInter: true,
			wantResult:   []int{5},
		},
		{
			name:         "Пустой первый слайс",
			slice1:       []int{},
			slice2:       []int{1, 2},
			wantHasInter: false,
			wantResult:   []int{},
		},
		{
			name:         "Пустой второй слайс",
			slice1:       []int{1, 2},
			slice2:       []int{},
			wantHasInter: false,
			wantResult:   []int{},
		},
		{
			name:         "Оба пустых",
			slice1:       []int{},
			slice2:       []int{},
			wantHasInter: false,
			wantResult:   []int{},
		},
		{
			name:         "Дубликаты в slice1",
			slice1:       []int{3, 5, 3, 7, 5},
			slice2:       []int{5, 9},
			wantHasInter: true,
			wantResult:   []int{5},
		},
		{
			name:         "Дубликаты в обоих",
			slice1:       []int{1, 1, 2, 2},
			slice2:       []int{2, 2, 3},
			wantHasInter: true,
			wantResult:   []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasInter, result := Intersect(tt.slice1, tt.slice2)
			if hasInter != tt.wantHasInter {
				t.Errorf("hasIntersection = %v, want %v", hasInter, tt.wantHasInter)
			}
			if !reflect.DeepEqual(result, tt.wantResult) {
				t.Errorf("result = %v, want %v", result, tt.wantResult)
			}
		})
	}
}
