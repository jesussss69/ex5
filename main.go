package main

import {
	"fmt"
	"log"
}

func main() {

}

func createNewCar(brand string, year int) Icar {
	car, err := carStruct.NewCar(year, brand)
	if err != nil {
		log.Fatal(err)
	}
	return car
}