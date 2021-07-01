//Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов(22+32+42….) с использованием конкурентных вычислений.
package main
import (
	"fmt"
)
func compute(inArr [5]int, outArr *[5]int, i int) {
	outArr[i] = inArr[i] * inArr[i]
}
func main() {
	var inArr [5]int = [5]int{2, 4, 6, 8, 10}
	var outArr [5]int
	ch := make(chan struct{})
	func() {
		defer close(ch)
		for i := range inArr {
			compute(inArr, &outArr, i)
		}

	}()
	<-ch // ждем выполнения анонимной функции
	var summ int
	for i := range outArr {
		summ = summ + outArr[i]

	}
	fmt.Println(summ) //теперь можно печатать
}