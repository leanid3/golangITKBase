// merge_test.go
package main

import (
	"reflect"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	// Создаём каналы и заполняем их
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

	merged := MergeChannels(ch1, ch2, ch3)

	var result []int
	for val := range merged {
		result = append(result, val)
	}

	expectedSet := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	if len(result) != 5 {
		t.Fatalf("Ожидалось 5 значений, получено %d", len(result))
	}

	for _, v := range result {
		if !expectedSet[v] {
			t.Errorf("Получено неожиданное значение: %d", v)
		}
		delete(expectedSet, v) // защита от дубликатов
	}

	if len(expectedSet) != 0 {
		t.Errorf("Не все ожидаемые значения получены. Остались: %v", expectedSet)
	}
}

func TestMergeChannels_Empty(t *testing.T) {
	merged := MergeChannels()

	_, ok := <-merged
	if ok {
		t.Error("Ожидался закрытый канал, но получено значение")
	}
}

func TestMergeChannels_OneChannel(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)

	merged := MergeChannels(ch)

	want := []int{10, 20, 30}
	got := []int{}
	for v := range merged {
		got = append(got, v)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Получено %v, ожидалось %v", got, want)
	}
}

func TestMergeChannels_SomeEmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	close(ch1)

	ch2 := make(chan int, 2)
	ch2 <- 100
	ch2 <- 200
	close(ch2)

	ch3 := make(chan int)
	close(ch3)

	merged := MergeChannels(ch1, ch2, ch3)

	got := []int{}
	for v := range merged {
		got = append(got, v)
	}

	if len(got) != 2 {
		t.Fatalf("Ожидалось 2 значения, получено %d", len(got))
	}

	found100, found200 := false, false
	for _, v := range got {
		if v == 100 {
			found100 = true
		}
		if v == 200 {
			found200 = true
		}
	}

	if !found100 || !found200 {
		t.Error("Не все значения получены")
	}
}
