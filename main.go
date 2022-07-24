package main

import (
	"fmt"
)

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	var numStream = make(chan int)
	var sum int

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	go func() {
		defer close(numStream)
		for _, slice := range n {
			for _, num := range slice {
				numStream <- num
			}
		}
	}()
	for i := range numStream {
		sum += i
	}
	fmt.Printf("result: %v\n", sum)
}
