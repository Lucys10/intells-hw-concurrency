package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	var wg sync.WaitGroup

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	for i := 0; i < len(n); i++ {

		wg.Add(1)
		go func(sl []int, numberSl int) {
			defer wg.Done()
			sum(sl, numberSl)
		}(n[i], i)
	}

	wg.Wait()

}

func sum(sl []int, numberSl int) {
	sumSl := 0
	for _, v := range sl {
		sumSl += v
	}
	fmt.Printf("slice %v: %v\n", numberSl, sumSl)

}
