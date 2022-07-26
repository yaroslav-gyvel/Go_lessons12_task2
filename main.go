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
	var sum int

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	m := make(map[int]chan int)

	for k, slice := range n {
		m[k] = make(chan int)
		go sumSlice(slice, m[k])
	}

	for k := 0; k < len(n); k++ {
		sum += <-m[k]
	}

	fmt.Printf("result: %v\n", sum)
}

func sumSlice(s []int, Ch chan<- int) {
	total := 0
	for _, num := range s {
		total += num
	}
	Ch <- total
	close(Ch)
}
