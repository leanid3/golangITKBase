package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestGetType(t *testing.T) {
	if got := GetType(42); got != "int" {
		t.Errorf("GetType(42) = %s; want int", got)
	}
	if got := GetType(3.14); got != "float64" {
		t.Errorf("GetType(3.14) = %s; want float64", got)
	}
	if got := GetType("test"); got != "string" {
		t.Errorf("GetType(\"test\") = %s; want string", got)
	}
	if got := GetType(true); got != "bool" {
		t.Errorf("GetType(true) = %s; want bool", got)
	}
	if got := GetType(complex64(1 + 2i)); got != "complex64" {
		t.Errorf("GetType(complex64) = %s; want complex64", got)
	}
}

func TestConvertToString(t *testing.T) {
	result := ConvertToString(42, 42, 42, 3.14, "Golang", true, complex64(1+2i))
	expected := "4242423.14Golangtrue(1+2i)"
	if result != expected {
		t.Errorf("ConvertToString = %s; want %s", result, expected)
	}
}

func TestHashSalt(t *testing.T) {
	input := "4242423.14Golangtrue(1+2i)"
	runes := []rune(input)

	// Вручную воссоздаём ожидаемый хэш для проверки
	mid := len(input) / 2
	salted := input[:mid] + "go-2024" + input[mid:]
	expectedHash := sha256.Sum256([]byte(salted))
	expected := fmt.Sprintf("%x", expectedHash)

	got := HashSalt(runes, "go-2024")
	if got != expected {
		t.Errorf("HashSalt = %s; want %s", got, expected)
	}
}
