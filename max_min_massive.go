package main // Создайте функцию, которая находит минимальное и максимальное значение в массиве.

import (
	"fmt"
)

func main() {
	min, max := maxMin()
	fmt.Println("Минимум:", min)
	fmt.Println("Максимум:", max)
}

func maxMin() (int, int) {
	numbers := [5]int{1, 2, 3, 4, 5}
	min := numbers[0]
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}
	return min, max
}
