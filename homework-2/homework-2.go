package main

/*
Homework-2
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
	// Задание 3
	fmt.Println(fib(99))
}

// Написать функцию, которая определяет, четное ли число.
func isEven(n int) {
	if n%2 == 0 {
		fmt.Print(n, " is Even")
	} else {
		fmt.Println(n, "is not even")
	}
}

// Написать функцию, которая определяет, делится ли число без остатка на 3.

func third(n int) {
	if n%3 == 0 {
		fmt.Printf("\n%v can be divided by 3", n)
	} else {
		fmt.Printf("\n%v cannot be divided by 3", n)
	}
}

// Написать функцию, которая последовательно выводит на экран N первых чисел Фибоначчи,
// начиная от 0. (числа Фибоначчи определяются соотношениями ​ Fn=Fn-1 + Fn-2)

/* func fibonacciNumners(x int) {
	var fn0 = 1
	var fn1 = 2
	fmt.Println(fn0)
	fmt.Println(fn1)
	for fn := 2; fn <= x; fn += 1 {
		fn = (fn - 1) + (fn - 2)
		fmt.Println(fn)
	}

} */

func fib(N uint) uint {
	var table []uint
	table = make([]uint, N+1)
	table[0] = 0
	table[1] = 1
	for i := uint(2); i <= N; i += 1 {
		table[i] = table[i-1] + table[i-2]
		fmt.Println(table[i])
	}
	return table[N]
}
