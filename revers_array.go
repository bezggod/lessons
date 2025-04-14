package main //Реализуйте функцию, которая принимает массив и возвращает его в обратном порядке.

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	reversed := reverse(arr)
	fmt.Println(reversed)
}

func reverse(arr [5]int) [5]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
