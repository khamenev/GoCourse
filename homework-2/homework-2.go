package main

/*
Homework-2: Tasks 1-4
Sergei Khamenev
24.04.2019
*/

import "fmt"

func main() {
	// Задание 	1
	fmt.Println("Enter a number:")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	isEven(n)
	// Задание 2
	third(n)
	// Задание 3 и 4
	fmt.Printf("\n %v", fib(30))
}

// 1. Написать функцию, которая определяет, четное ли число.
func isEven(n int) {
	if n%2 == 0 {
		fmt.Println(n, " is Even")
	} else {
		fmt.Println(n, "is not even")
	}
}

// 2. Написать функцию, которая определяет, делится ли число без остатка на 3.

func third(n int) {
	if n%3 == 0 {
		fmt.Printf("%v can be divided by 3", n)
	} else {
		fmt.Printf("%v cannot be divided by 3", n)
	}
}

// 3. Написать функцию, которая последовательно выводит на экран N первых чисел Фибоначчи,
// 	  начиная от 0. (числа Фибоначчи определяются соотношениями ​ Fn=Fn-1 + Fn-2)
// 4. Написать функцию, которая заполняет срез числами Фибоначчи.

func fib(N uint) []uint {
	var table []uint // new slise
	table = make([]uint, N+1)
	table[0] = 0 // zero element
	table[1] = 1 // first element
	for i := uint(2); i <= N; i++ {
		table[i] = table[i-1] + table[i-2]
		//fmt.Println(table[i])
	}
	return table
}
