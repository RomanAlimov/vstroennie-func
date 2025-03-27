package main

import "fmt"

func umn(function func(int) int) {
	fmt.Println(function(5))
}

func st(fn string) func() {
	return func() {
		fmt.Println(fn)
	}
}

func main() {
	a := func(x int, y int) { // переменная где будет встроенная функция
		fmt.Println(x + y) // складываю 2 переменные
	}
	a(1, 2) // инициализация переменных = 3

	b := func(b int, g int) int { // функция с возвращением
		return b + g
	}
	sum := b(4, 6)   // переменная для работы return
	fmt.Println(sum) // 10

	// далее сделал(выше) функцию для переменной в встроеной функции

	summa := func(x int) int {
		return x * x
	}
	umn(summa) // 5 na 5 = 25

	// еще добавлю со строками

	st("RomanAlimov")() // Выводит RomanAlimov без проблем. () для инициализации в main

}
