package main // Напишите функцию, которая вычисляет среднее значение элементов массива и возвращает его.

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	avg := average(array)

	fmt.Println(avg)

	fmt.Println("Надеюсь все заработало: аминь  ", avg)
}

func average(array [5]int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
