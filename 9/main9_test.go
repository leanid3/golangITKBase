// pipeline_test.go
package main

import (
	"math"
	"reflect"
	"testing"
)

func TestCubePipeline(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)

	go func() {
		defer close(input)
		testData := []uint8{0, 1, 2, 10, 255}
		for _, n := range testData {
			input <- n
		}
	}()

	var results []float64
	for val := range output {
		results = append(results, val)
	}

	expected := []float64{
		0,
		1,
		8,
		1000,
		16581375,
	}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Получено: %v, ожидалось: %v", results, expected)
	}
}

func TestCubePipeline_Empty(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)

	close(input)

	results := []float64{}
	for val := range output {
		results = append(results, val)
	}

	if len(results) != 0 {
		t.Errorf("Ожидался пустой результат, получено: %v", results)
	}
}

func TestCubePipeline_SingleValue(t *testing.T) {
	input := make(chan uint8, 1)
	output := make(chan float64)

	input <- 5
	close(input)

	go CubePipeline(input, output)

	result := <-output
	_, open := <-output
	if open {
		t.Error("Выходной канал не закрыт после обработки")
	}

	expected := 125.0 // 5^3
	if result != expected {
		t.Errorf("Результат = %v, ожидалось %v", result, expected)
	}
}

func TestCubePipeline_MaxUint8(t *testing.T) {
	input := make(chan uint8, 1)
	output := make(chan float64)

	input <- 255
	close(input)

	go CubePipeline(input, output)

	result := <-output
	_, open := <-output
	if open {
		t.Error("Канал не закрыт")
	}

	expected := math.Pow(255, 3)
	if result != expected {
		t.Errorf("Результат = %v, ожидалось %v", result, expected)
	}
}
