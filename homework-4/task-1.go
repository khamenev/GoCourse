package main

/*
** Sergei Khamenev
** 09.05.2019
** Task-1: Написать свой интерфейс и создать несколько структур, удовлетворяющих ему
 */

import "fmt"

type Printer interface {
	whatIs()
}

func printWallet(p Printer) {
	p.whatIs()
}

// cash structure

type Cash struct{}

// card structure

type Card struct{}

func (cash Cash) whatIs() {
	fmt.Println("Money: cash")
}

func (card Card) whatIs() {
	fmt.Println("Money: card")
}

func main() {
	rublesCash := Cash{}
	rublesCard := Card{}
	printWallet(rublesCash)
	printWallet(rublesCard)
}
