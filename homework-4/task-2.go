package main

import (
	"fmt"
	"sort"
)

/*
** Sergei Khamenev
** 09.05.2019
** Task-2: Создать тип, описывающий контакт в телефонной книге. Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него интерфейс Sort{}
 */

type Elements struct {
	Name        string
	PhoneNumber string
}

type PhoneBook []Elements

// Lens method
func (pb PhoneBook) Len() int {
	return len(pb)
}

// Less method
func (pb PhoneBook) Less(i, j int) bool {
	return pb[i].Name < pb[j].Name
}

// Swap method
func (pb PhoneBook) Swap(i, j int) {
	pb[i], pb[j] = pb[j], pb[i]
}

func main() {
	contacts := []Elements{
		{"Sergei", "79162222222"},
		{"Maks", "79162222223"},
		{"Niko", "79162222224"},
	}
	fmt.Println(contacts)
	sort.Sort(PhoneBook(contacts))
	fmt.Println(contacts)
}
