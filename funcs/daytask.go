package function

import (
	"day4/taskd2"
	"day4/taskd4"
	"day4/taskd7"
	"fmt"
	//"log"
	//"math"
)

func Daytask() {

	day1 := "День 1. Знакомство."
	day2 := "День 2. Первая прогграмма."
	day3 := "День 3. Типы, Функции и Модули."
	day4 := "День 4. Условия и Циклы."
	day5 := "День 5. Структуры."
	day6 := "День 6. Указатели."
	day7 := "День 7. Массивы и Срезы."

	fmt.Println("Здравствуйте! Пожалуйства, выберите день обучения:")
	fmt.Println(day1)
	fmt.Println(day2)
	fmt.Println(day3)
	fmt.Println(day4)
	fmt.Println(day5)
	fmt.Println(day6)
	fmt.Println(day7)

	var day, task int
	fmt.Scan(&day)

	switch day {
	case 1: // дз 1 дня
		fmt.Println("В этот  день нет заданий.")
	case 2: // дз 2 дня
		fmt.Println("Вы выбрали 2 день. Выберите задание:\n" +
			"1. Периметр квадрата\n" +
			"2. Периметр прямоугольника\n" +
			"3. Площадь прямоугольника\n" +
			"4. До юбилея")

		fmt.Scan(&task)
		switch task {
		case 1:
			// perimetrk
			taskd2.Perimeterk()
		case 2:
			// perimetrp
			taskd2.Perimetrp()
		case 3:
			// areap
			taskd2.Areap()
		case 4:
			// anniversaryyear
			taskd2.Anniversaryyear()
		default:
			fmt.Println("Выбранное вами дз не существует.")
		}

	case 3: // дз 3 дня
		fmt.Println("Вы выбрали 3 день. Выберите задание:\n" +
			"1. Разбить блоки кода на функции")
		// я разделил код на 4 фунцкии и вызывал их в main
		fmt.Scan(&task)
		if task == 1 {
			taskd2.Perimeterk()
			taskd2.Perimetrp()
			taskd2.Areap()
			taskd2.Anniversaryyear()
		} else {
			fmt.Println("Выбранное вами дз не существует.")
		}
	case 4:
		{ // дз 4 дня
			fmt.Println("Вы выбрали 4 день. Выберите задание:\n" +
				"1. Обединить дз в единую программу.\n" +
				"2. Вывести таблицу умножения в консоли.")
			fmt.Scan(&task)
			if task == 1 {
				fmt.Println("Данная программа и есть это задание.")
			} else if task == 2 {
				taskd4.Table()
			} else {
				fmt.Println("Выбранное вами дз не существует.")
			}

		}
	case 5: // дз 5 дня
		fmt.Println("Вы выбрали 5 день. Выберите задание:\n" +
			"1. Собственный тематический мини-проект.")
		fmt.Scan(&task)
		if task == 1 {
            fmt.Println()
		} else {
			fmt.Println("Выбранное вами дз не существует.")
		}

	case 6: // дз 6 дня
		fmt.Println("Вы выбрали 6 день. Выберите задание:\n" +
			"1. Рефакторинг кода Дурной Больницы.")
		fmt.Scan(&task)
		if task == 1 {
        fmt.Println()
		} else {
			fmt.Println("Выбранное вами дз не существует.")
		}

	case 7: // дз 7 дня
		fmt.Println("Вы выбрали 4 день. Выберите задание:\n" +
			"1. Функция, которая отзеркаливает массив наоборот.\n" +
			"2. Функция, которая разбивает один массив с числами на два слайса.")
		fmt.Scan(&task)
		if task == 1 {
			m := []string{"Hello", "how", "are", "you", "?"}
			mirm := taskd7.Mirror(m)
			fmt.Println("Перевернутый массив строк:", mirm)
		} else if task == 2 {
			n := []int{1, 2, 3, 4, 5, 6, 19, 20, 27, 28}
			even, odd := taskd7.Split(n)
			fmt.Println("Четные числа:", even)
			fmt.Println("нечетные числа:", odd)
		} else {
			fmt.Println("Выбранное вами дз не существует.")
		}
	}
}
