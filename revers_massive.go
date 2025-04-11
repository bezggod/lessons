package main //Реализуйте функцию, которая принимает массив и возвращает его в обратном порядке.

import "fmt"

func main() {
	reversed := reverse()
	fmt.Println(reversed)
}

func reverse() []int {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
