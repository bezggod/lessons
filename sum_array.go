package main /* Напишите функцию, которая принимает массив
целых чисел и возвращает сумму всех его элементов.*/

import "fmt"

func main() {
	fmt.Println(arrays())
}

func arrays() int {
	numbers := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
