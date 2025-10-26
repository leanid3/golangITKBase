package main

import "fmt"

// #### Задание 3

// Реализуйте структуру данных StringIntMap, которая будет использоваться для хранения пар "строка - целое число". Ваша
// структура должна поддерживать следующие операции:

// 1. Добавление элемента: Метод Add(key string, value int), который добавляет новую пару "ключ-значение" в карту.

// 2. Удаление элемента: Метод Remove(key string), который удаляет элемент по ключу из карты.

// 3. Копирование карты: Метод Copy() map[string]int, который возвращает новую карту, содержащую все элементы текущей карты.

// 4. Проверка наличия ключа: Метод Exists(key string) bool, который проверяет, существует ли ключ в карте.

// 5. Получение значения: Метод Get(key string) (int, bool), который возвращает значение по ключу и булевый флаг, указывающий на успешность операции.

// * Напишите unit тесты к созданным функциям

func main() {
	m := NewStringIntMap()
	m.Add("count", 42)
	m.Add("version", 1)

	fmt.Println("Существует 'count':", m.Exists("count"))
	if value, exists := m.Get("count"); exists {
		fmt.Println("Значение 'count':", value)
	} else {
		fmt.Println("Ключ 'count' не найден")
	}

	m.Remove("version")
	fmt.Println("После удаления 'version', существует:", m.Exists("version"))

	copyMap := m.Copy()
	fmt.Println("Копия карты:", copyMap)
}
