package main

import (
	"math"
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	ch := RandomGenerator()

	count := 10
	received := make([]int, 0, count)

	for i := 0; i < count; i++ {
		val, ok := <-ch
		if !ok {
			t.Fatal("Канал закрыт раньше времени")
		}
		if val < math.MinInt || val > math.MaxInt {
			t.Errorf("Получено недопустимое значение: %d", val)
		}
		received = append(received, val)
	}

	select {
	case _, ok := <-ch:
		if !ok {
			t.Error("Канал неожиданно закрыт")
		}
	default:
	}

	if len(received) != count {
		t.Errorf("Ожидалось %d значений, получено %d", count, len(received))
	}
}

func TestRandomGenerator_Diversity(t *testing.T) {
	ch := RandomGenerator()
	values := make(map[int]bool)

	for i := 0; i < 100; i++ {
		val, ok := <-ch
		if !ok {
			t.Fatal("Канал закрыт")
		}
		if values[val] {
			t.Logf("Найден дубликат: %d (нормально при большом объёме)", val)
		}
		values[val] = true
	}

	if len(values) < 90 {
		t.Errorf("Слишком много дубликатов: уникальных = %d из 100", len(values))
	}
}
