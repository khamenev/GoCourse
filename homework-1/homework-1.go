package main

import (
	"fmt"
	"math"
)

func main() {
	println("Hello")

	// Конвертация рублей в доллары
	var amountRur float64
	fmt.Println("Please enter your amount in RUR")
	fmt.Scanln(&amountRur)
	fmt.Println("Your amount in USD is", converter(amountRur), "$")

	// Находим площадь, периметр и гипотенузу треугольника
	fmt.Println("Please enter your amount in RUR")
	fmt.Scanln(&amountRur)
}

// Написать программу для конвертации рублей в доллары. Программа запрашивает сумму в
// рублях и выводит сумму в долларах. Курс доллара задайте константой.

func converter(amountRur float64) float64 {
	const currencyUsd float64 = 65.5
	return math.Round(amountRur / currencyUsd) //округляем до целого числа
}

//Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу.
//Используйте тип данных float64 и функции из пакета math.
