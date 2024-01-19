package carStruct

import (
	"fmt"
)

type Icar interface {
	SetBrand(c string)
	GetBrand() string

	SetYear(c int) error
	GetYear() int

	GetSpeed() float64
}

var m = map[string]map[int]float64{
	"Ford": map[int]float64{
		2024: 130,
		2000: 100,
		1950: 55,
		1900: 50,
	},
	"Merc": map[int]float64{
		2024: 300,
		2000: 100,
		1950: 45,
		1900: 40,
	},
	"RR": map[int]float64{
		2024: 150,
		2000: 120,
		1950: 75,
		1900: 45,
	},
}

type car struct {
	brand string
	year  int
	speed float64
}

func NewCar(y int, b string) (Icar, error) {
	result := car{
		brand: b,
	}
	e := result.SetYear(y)
	result.determineSpeed(result.year)
	return &result, e
}

func (c *car) SetYear(num int) error {
	if num < 1900 || num > 2024 {
		return fmt.Errorf("отсоси мои большие черные яйца")
	}
	c.year = num
	return nil
}

func (c *car) SetBrand(s string) {
	c.brand = s
}

func (c car) GetYear() int {
	return c.year
}

func (c car) GetBrand() string {
	return c.brand
}

func (c car) GetSpeed() float64 {
	return c.speed
}

func (c *car) determineSpeed(year int) {
	for y, s := range m[c.brand] {
		if c.year <= y {
			c.speed = s
			return
		}
	}
}

func TryAdd(targetSlice *[]Icar, g Icar) bool {
	for _, v := range *targetSlice {
		if g.GetBrand() == v.GetBrand() && g.GetYear() == v.GetYear() {
			return false
		}
	}
	*targetSlice = append(*targetSlice, g)
	return true
}
