//Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов(22+32+42….) с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)
func compute(inArr [5]int, outArr *[5]int, i int) {
	outArr[i] = inArr[i] * inArr[i]
}
func main() {
	var inArr [5]int = [5]int{2, 4, 6, 8, 10}
	var outArr [5]int
	//счетчик горутин
	wg := new(sync.WaitGroup)
	//мьютекс для избежания одновременной записи/чтения
	mu := new(sync.Mutex)
	for i := range inArr {
		wg.Add(1)
		go func(WG *sync.WaitGroup, MU *sync.Mutex, i int) {
			defer WG.Done()
			MU.Lock()
			outArr[i] = inArr[i] * inArr[i]
			MU.Unlock()
		}(wg, mu, i)
	}
	wg.Wait() // ждем пока счетчик обнулится
	var summ int
	for i := range outArr {
		summ = summ + outArr[i]
	}
	fmt.Println(summ) //теперь можно печатать
}