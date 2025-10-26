// stringintmap.go
package main

// StringIntMap — обёртка над map[string]int с безопасными методами
type StringIntMap struct {
	data map[string]int
}

// NewStringIntMap создаёт новый экземпляр StringIntMap
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		make(map[string]int),
	}
}

// Add добавляет пару ключ-значение в карту
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove удаляет элемент по ключу (если ключ отсутствует — ничего не делает)
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy возвращает **глубокую копию** внутренней карты
func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int, len(m.data))
	for k, v := range m.data {
		copyMap[k] = v
	}
	return copyMap
}

// Exists проверяет, существует ли ключ в карте
func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

// Get возвращает значение по ключу и флаг успешности
func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]
	return val, ok
}
