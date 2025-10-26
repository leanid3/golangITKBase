package main

import "fmt"

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)

	go func() {
		defer close(input)
		numbers := []uint8{1, 2, 3, 4, 5}
		for _, n := range numbers {
			input <- n
		}
	}()

	fmt.Println("Кубы чисел:")
	for cubed := range output {
		fmt.Println(cubed)
	}
}

func CubePipeline(in <-chan uint8, out chan<- float64) {
	defer close(out)
	for val := range in {
		cubed := float64(val) * float64(val) * float64(val)
		out <- cubed
	}
}
