package main

import "fmt"

// #### Задание 2

// 1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным
// образом (при каждом запуске должны получаться новые значения)

// 2. Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.

// 3. Напишите функцию addElements, которая принимает слайс и число. Функция должна добавлять это число в конец слайса и возвращать новый слайс.

// 4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.

// 5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.

// 6. Напишите main функцию, в которой протестируете все вышеописанные функции. Выведите результаты на экран.

// * Напишите unit тесты к созданным функциям

// Примечание.
// В качестве originalSlice можно использовать ```originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}```

func main() {
	//1
	origSlice := GenerateRandomSlice()
	fmt.Println("Созданный слайс: ", origSlice)

	//2
	evenSlice := sliceExample(origSlice)
	fmt.Println("Четные числа", evenSlice)

	//3
	sliceAddElem := addElements(origSlice, 999)
	fmt.Println("Добавить число: ", sliceAddElem)

	//4
	copid := copySlice(origSlice)
	origSlice[0] = -1
	fmt.Println("OrigSlice", origSlice)
	fmt.Println("copySlice", copid)

	if len(origSlice) > 0 {
		removed, err := removeElement(origSlice, 0)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("После удаления: ", removed)
		}
	}
}
