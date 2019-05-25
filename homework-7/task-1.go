package main

/*
** S.Khamenev
** 22.05.2019
** Homework-7: Task-1
 */

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	time.Sleep(10 * time.Second) // waiting
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
