// Написать программу, которая в конкурентном
//  виде читает элементы из массива в stdout.
package main
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
func readerArr(ch <-chan int) {		//декремент счетчика
	for  {
		fmt.Printf("%v\n", <-ch) //вывод
	}
}
func main() {
	start := time.Now() //отсчет времени
	var arr [1005000]int
	for i := range arr {
		arr[i] = rand.Intn(10) //прописка значений
	}
	ch := make(chan int)
	for i := 0; i < 1; i++ { //запуск горутин для чтения
		go readerArr(ch)
	}
	wg := new(sync.WaitGroup)
	for i := range arr {
		wg.Add(1)
		go func(j int, WG *sync.WaitGroup) {
			defer WG.Done()
			ch <- arr[j] //отправка в канал для чтения
		}(i, wg)
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Время выполнения: ", duration)
}