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
	arr := [7]int{1, 20, 8, 3, 19, 99, 6}
	min := arr[0]
	max := arr[0]
	for _, value := range arr {
		if value > min {
			min = value
		}
		if value < max {
			max = value
		}
	}
	return max, min

}
