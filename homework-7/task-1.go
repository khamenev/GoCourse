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
	go spinner()
	time := countdown(10)
	fmt.Println(time)
}

func spinner() {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
		}
	}
}

func countdown(t int) string {
	tick := make(<-chan time.Time)    // создаем однонаправленный канал
	tick = time.Tick(1 * time.Second) // создаём поток секундных "тиков"
	for countdown := t; countdown > 0; countdown-- {
		moment := <-tick
		fmt.Println(moment.Format("15:04:05"), countdown)
	}
	return "It is Time! (c) Bruce Buffer"
}
