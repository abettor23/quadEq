package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Эта программа выводит корни квадратного уравнения.")

	fmt.Println("Введите значение кэфицента a")
	var a float64
	fmt.Scan(&a)
	if a == 0 {
		fmt.Println("Уравнение не является квадратным.")
		return
	}

	fmt.Println("Введите значение кэфицента b")
	var b float64
	fmt.Scan(&b)

	fmt.Println("Введите значение кэфицента c")
	var c float64
	fmt.Scan(&c)

	D := math.Pow(b, 2) - 4*a*c

	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Print("Уравнение имеет два различных вещественных корня: x1 = ", x1, ", x2 = ", x2, "\n")
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Println("Уравнение имеет один кратный вещественный корень: x =", x)
	} else {
		fmt.Println("Уравнение не имеет корней.")
	}
}
