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
	var a float64 // катет 1
	var b float64 // катет 2
	fmt.Println("Please enter the legs of a right triangle")
	fmt.Scanln(&a, &b)
	var c = triangleHypotenuse(a, b) // гипотенуза
	fmt.Println("The area of a right triangle:", triangleSqrt(a, b))
	fmt.Println("The perimeter of a right triangle:", trianglePerimeter(a, b, c))
	fmt.Println(c)
}

// Написать программу для конвертации рублей в доллары. Программа запрашивает сумму в
// рублях и выводит сумму в долларах. Курс доллара задайте константой.

func converter(amountRur float64) float64 {
	const currencyUsd float64 = 65.5
	return (amountRur / currencyUsd)
}

//Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу.
//Используйте тип данных float64 и функции из пакета math.

//Функция по вычислению площади
func triangleSqrt(a, b float64) float64 {
	return a * b * 0.5
}

//Функция по вычислению гипотенузы
func triangleHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

//Функция по вычислению периметра
func trianglePerimeter(a, b, c float64) float64 {
	return a + b + c
}
