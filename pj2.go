// Написать программу, которая конкурентно рассчитает значение квадратов значений взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
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
	fmt.Println(outArr) //теперь можно печатать
}
