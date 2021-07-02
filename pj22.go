// Написать быструю сортировку встроенными методами языка
package main
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)
func qSort(b int, e int, arr []int) {

	l, r := b, e //фиксируем границы
	m := (l + r) / 2//находим контрольную ячейку
	piv := arr[m]
	for l <= r {
		for arr[l] < piv {
			l++
		}
		for arr[r] > piv {
			r--
		}
		if l <= r { // если оба цикла нашли неправильную последовательность то происходит свап
			arr[l], arr[r] = arr[r], arr[l]
			r--
			l++
		}
	}
	if b < r {//если r не дошел до начала то рекурсивное продолжение
		qSort(b, r, arr) 
	}
	if e > l {//если l не дошел до конца то рекурсивное продолжение
		qSort(l, e, arr)
	}
}
func main() {
	
	lenS := 1000000
	slice := make([]int, lenS, lenS)
	slice2 := make([]int, lenS, lenS)
	slice3 := make([]int, lenS, lenS)
	for i := 0; i < lenS; i++ {
		slice[i] = rand.Intn(23)
		slice2[i] = slice[i]
		slice3[i] = slice[i]
	}
	start1 := time.Now()
	for i := lenS - 1; i > 0; i-- { //это сортировка пузырькоом по памяти написал для сравнения скорости, думал зачем эти быстрые сортировки если есть чудесная пузырьковая))
		for j := 0; j < i; j++ {
			if slice[j] > slice[j + 1] {
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}
	}
	duration1 := time.Since(start1)
	fmt.Println("1 Время выполнения: ", duration1)
	// fmt.Println(slice)

	start2 := time.Now()
	sort.Ints(slice2)//сортировка из пакета sort
	duration2 := time.Since(start2)
	fmt.Println("2 Время выполнения: ", duration2)
	// fmt.Println(slice2)

	start3 := time.Now()
	qSort(0, lenS - 1, slice3) //быстрая сортировка и вправду быстрая
	duration3 := time.Since(start3)
	fmt.Println("3 Время выполнения: ", duration3)
}