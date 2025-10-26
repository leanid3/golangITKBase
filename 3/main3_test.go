package main

import "testing"

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	// Тест добавления
	m.Add("test", 42)
	if !m.Exists("test") {
		t.Error("Ключ 'test' должен существовать после добавления")
	}

	// Тест получения значения
	if value, exists := m.Get("test"); !exists || value != 42 {
		t.Errorf("Ожидалось значение 42, получено %d, exists: %v", value, exists)
	}

	// Тест удаления
	m.Remove("test")
	if m.Exists("test") {
		t.Error("Ключ 'test' не должен существовать после удаления")
	}

	// Тест копирования
	m.Add("key1", 1)
	m.Add("key2", 2)
	copyMap := m.Copy()

	if len(copyMap) != 2 {
		t.Errorf("Ожидалась длина копии 2, получена %d", len(copyMap))
	}

	// Проверяем независимость копии
	m.Add("key3", 3)
	if len(copyMap) != 2 {
		t.Error("Изменение оригинала повлияло на копию")
	}
}

func TestGetNonExistentKey(t *testing.T) {
	m := NewStringIntMap()
	value, exists := m.Get("nonexistent")
	if exists {
		t.Error("Ключ 'nonexistent' не должен существовать")
	}
	if value != 0 {
		t.Errorf("Ожидалось значение 0 для несуществующего ключа, получено %d", value)
	}
}
