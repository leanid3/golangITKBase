// diff_test.go
package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "Обычный случай",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "Пустой slice2",
			slice1:   []string{"a", "b"},
			slice2:   []string{},
			expected: []string{"a", "b"},
		},
		{
			name:     "Пустой slice1",
			slice1:   []string{},
			slice2:   []string{"x", "y"},
			expected: []string{},
		},
		{
			name:     "Нет различий",
			slice1:   []string{"x", "y"},
			slice2:   []string{"x", "y", "z"},
			expected: []string{},
		},
		{
			name:     "Дубликаты в slice1",
			slice1:   []string{"a", "b", "a", "c"},
			slice2:   []string{"b"},
			expected: []string{"a", "a", "c"},
		},
		{
			name:     "Полное отсутствие пересечений",
			slice1:   []string{"1", "2"},
			slice2:   []string{"3", "4"},
			expected: []string{"1", "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Difference() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
