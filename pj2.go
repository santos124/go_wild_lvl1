// Написать программу, которая конкурентно рассчитает значение квадратов значений взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main
import (
	"fmt"
	"sync"
)

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
	wg.Wait()
	fmt.Println(outArr) //теперь можно печатать
}
