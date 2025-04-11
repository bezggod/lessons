package main /* Напишите функцию, которая принимает массив
целых чисел и возвращает сумму всех его элементов.*/

import (
	"fmt"
)

func main() {
	fmt.Println(massive1())

}

func massive1() int {
	numbers := [...]int{1, 3, 4, 7, 8}
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
