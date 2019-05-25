package main

/*
** S.Khamenev
** 22.05.2019
** Homework-7: Task-2
 */

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 10; x++ {

			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for x := range squares {
		fmt.Println(x)
	}

}
