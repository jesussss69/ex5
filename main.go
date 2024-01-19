package main

import (
	"fmt"
	"log"

	"github.com/jesussss69/carStruct/internal/carStruct"
)

func main() {
	carSlice := make([]carStruct.Icar, 0)

	carSlice = append(carSlice, createNewCar("Ford", 1900))
	carSlice = append(carSlice, createNewCar("RR", 2000))
	carSlice = append(carSlice, createNewCar("Merc", 1950))
	carSlice = append(carSlice, createNewCar("RR", 1900))

	fmt.Println(carStruct.TryAdd(&carSlice, createNewCar("RR", 1950)))
}

func createNewCar(brand string, year int) carStruct.Icar {
	car, err := carStruct.NewCar(year, brand)
	if err != nil {
		log.Fatal(err)
	}
	return car
}
