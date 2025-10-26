package main

import (
	"testing"
)

func TestGenerateRandomSlice(t *testing.T) {
	slice := GenerateRandomSlice()

	// Проверяем, что слайс содержит 10 элементов
	if len(slice) != 10 {
		t.Errorf("Ожидалась длина 10, получена %d", len(slice))
	}

	// Проверяем, что все элементы являются целыми числами
	for i, v := range slice {
		if v < 0 || v >= 100 {
			t.Errorf("Элемент %d имеет неожиданное значение: %d", i, v)
		}
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
		{[]int{1, 3, 5, 7}, []int{}},
		{[]int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := sliceExample(test.input)
		if !slicesEqual(result, test.expected) {
			t.Errorf("sliceExample(%v) = %v, ожидалось %v", test.input, result, test.expected)
		}
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		input    []int
		num      int
		expected []int
	}{
		{[]int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
		{[]int{}, 1, []int{1}},
		{[]int{5}, 10, []int{5, 10}},
	}

	for _, test := range tests {
		result := addElements(test.input, test.num)
		if !slicesEqual(result, test.expected) {
			t.Errorf("addElements(%v, %d) = %v, ожидалось %v", test.input, test.num, result, test.expected)
		}
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	copied := copySlice(original)

	// Проверяем, что копия идентична оригиналу
	if !slicesEqual(original, copied) {
		t.Errorf("copySlice(%v) = %v, ожидалось %v", original, copied, original)
	}

	// Проверяем независимость копии
	copied[0] = 999
	if original[0] == 999 {
		t.Errorf("Изменение копии повлияло на оригинал")
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input    []int
		index    int
		expected []int
		hasError bool
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}, false},
		{[]int{1, 2, 3}, 0, []int{2, 3}, false},
		{[]int{1, 2, 3}, 2, []int{1, 2}, false},
		{[]int{1, 2, 3}, -1, nil, true},
		{[]int{1, 2, 3}, 3, nil, true},
		{[]int{}, 0, nil, true},
	}

	for _, test := range tests {
		result, err := removeElement(test.input, test.index)

		if test.hasError {
			if err == nil {
				t.Errorf("removeElement(%v, %d) должен был вернуть ошибку", test.input, test.index)
			}
		} else {
			if err != nil {
				t.Errorf("removeElement(%v, %d) вернул неожиданную ошибку: %v", test.input, test.index, err)
			}
			if !slicesEqual(result, test.expected) {
				t.Errorf("removeElement(%v, %d) = %v, ожидалось %v", test.input, test.index, result, test.expected)
			}
		}
	}
}

func TestIndexOutOfBoundsError(t *testing.T) {
	err := &IndexOutOfBoundsError{index: 5, length: 3}
	expectedMsg := "индекс 5 выходит за пределы длины слайса 3"
	if err.Error() != expectedMsg {
		t.Errorf("IndexOutOfBoundsError.Error() = %q, ожидалось %q", err.Error(), expectedMsg)
	}
}

// Вспомогательная функция для сравнения слайсов
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
