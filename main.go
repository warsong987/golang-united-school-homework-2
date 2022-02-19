package main

import (
	"fmt"
	"golang-united-school-homework-2/square"
)

func main() {
	fmt.Println("Площать круга: ", square.CalcSquare(10.0, square.SidesCircle))
	fmt.Println("Площать треугольника: ", square.CalcSquare(10.0, square.SidesTriangle))
	fmt.Println("Площать квадрата: ", square.CalcSquare(10.0, square.SidesSquare))
}
