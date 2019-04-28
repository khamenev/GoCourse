package main

import "fmt"

/*
Sergei Khamenev
28.04.2019
Homework-3
*/

// Задание 1
// Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны
// содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты
// ли окна, насколько заполнен объем багажника

type car struct {
	brand         string
	model         string
	vehicleYear   string
	trunkVolume   float64
	isStarting    bool
	isWindowsOpen bool
	filledPlace   float64
}

type truck struct {
	brand         string
	model         string
	vehicleYear   string
	bodyVolume    float64
	isStarting    bool
	isWindowsOpen bool
	filledPlace   float64
}

// Задание 2
// Инициализировать несколько экземпляров структур. Применить к ним различные действия.
// Вывести значения свойств экземпляров в консоль.

var car1 = car{
	brand:         "Skoda",
	model:         "Rapid",
	vehicleYear:   "2014",
	trunkVolume:   0.53,
	isStarting:    false,
	isWindowsOpen: false,
	filledPlace:   0.03,
}

var truck1 = truck{
	brand:         "KAMAZ",
	model:         "5490",
	vehicleYear:   "2019",
	bodyVolume:    12.0,
	isStarting:    false,
	isWindowsOpen: false,
	filledPlace:   0,
}

func main() {
	fmt.Println("Before changing:")
	fmt.Println(car1)
	fmt.Println(truck1)
	// Start the engine and put the bag in the trunk of the car
	car1.isStarting = true
	car1.filledPlace += 0.1
	//Start the engine and lower the truck window
	truck1.isStarting = true
	truck1.isWindowsOpen = true
	fmt.Println("After changing:")
	fmt.Println(car1)
	fmt.Println(truck1)
}
