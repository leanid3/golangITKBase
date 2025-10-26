package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

// ### Задание 1
// Напишите программу на Go, которая:

// 1. Создает несколько переменных различных типов данных:
// ```
// int (три числа в десятичной, восьмеричной и шеснадцатиричной системах)
// float64
// string
// bool
// complex64
// ```
// 2. Определяет тип каждой переменной и выводит его на экран.
// 3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
// 4. Преобразовать эту строку в срез рун.
// 5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.

// * Напишите unit тесты к созданным функциям

// Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.

// Входные числа из пункта 1 могут быть:
// ```
// var numDecimal int = 42           // Десятичная система
// var numOctal int = 052            // Восьмеричная система
// var numHexadecimal int = 0x2A     // Шестнадцатиричная система
// var pi float64 = 3.14             // Тип float64
// var name string = "Golang"         // Тип string
// var isActive bool = true           // Тип bool
// var complexNum complex64 = 1 + 2i  // Тип complex64
// ```

// ---
func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	//2
	fmt.Println("numDecimal", GetType(numDecimal))
	fmt.Println("numOctal", GetType(numOctal))
	fmt.Println("numHexadecimal", GetType(numHexadecimal))
	fmt.Println("pi", GetType(pi))
	fmt.Println("name", GetType(name))
	fmt.Println("isActive", GetType(isActive))
	fmt.Println("complexNum", GetType(complexNum))

	//3
	combined := ConvertToString(numDecimal, numDecimal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("\n Строка: ", combined)

	//4
	runes := []rune(combined)
	fmt.Println("Срез рун(длина): ", len(runes))

	//5
	hash := HashSalt(runes, "go-2024")
	fmt.Println("hash: ", hash)
}

func GetType(t interface{}) string {
	return fmt.Sprintf("%T", t)
}

func ConvertToString(numDecimal, numOctal, numHexadecimal int, pi float64, name string, isActive bool, complexNum complex64) string {
	arr := []string{
		strconv.Itoa(numDecimal),
		strconv.Itoa(numOctal),
		strconv.Itoa(numHexadecimal),
		strconv.FormatFloat(pi, 'f', -1, 64),
		name,
		strconv.FormatBool(isActive),
		strconv.FormatComplex(complex128(complexNum), 'f', -1, 64),
	}

	return strings.Join(arr, "")
}

func HashSalt(runes []rune, salt string) string {
	runeString := string(runes)
	mid := len(runeString) / 2
	withSalt := runeString[:mid] + salt + runeString[mid:]

	hash := sha256.Sum256([]byte(withSalt))
	return fmt.Sprintf("%x", hash)
}
