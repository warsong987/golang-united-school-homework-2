package square

import (
	"fmt"
)

func main() {
	fmt.Println("Площать круга: ", CalcSquare(10.0, SidesCircle))
	fmt.Println("Площать треугольника: ", CalcSquare(10.0, SidesTriangle))
	fmt.Println("Площать квадрата: ", CalcSquare(10.0, SidesSquare))
}
