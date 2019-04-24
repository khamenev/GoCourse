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
