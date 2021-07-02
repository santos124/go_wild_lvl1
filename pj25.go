//Написать свою структуру счетчик, которая 
// будет инкрементировать и выводить значения в конкурентной среде.
package main
import (
	"fmt"
	"sync"
	// "math/rand"
	"time"
)
func outer( clos chan bool, slice[][]int, mu *sync.Mutex) {
	for {
		select {
		case <- clos: //ожидания сигнала закрытия
			return
		default: //а пока печатаем слайс
			for i := range slice[1] {
				mu.Lock()//блокировка ячейки на случай одновременного чтения/записи
				if slice[1][i] == 0 {
					
					fmt.Println(slice[0][i])
					
					slice[1][i] = 1
					
				}
				mu.Unlock()
				
			}
		}
	}
	
	
}
func main() {
	lenS := 10000
	slice := make([][]int, 2) // создаем двумерный срез, 1 строка значения на вывод, вторая флаги(0 по умолчанию)
	for i := range slice {
		slice[i] = make([]int, lenS, lenS)
	}
	for i := range slice[0] {
		slice[0][i] = i
	}
	clos := make(chan bool)
	mu := new(sync.Mutex)//локер переменных и не только на случай параллельной записи/чтения при итерировании
	start1 := time.Now() //секундомер
	for i := 0; i < 3; i++ {
		go outer(clos, slice, mu)
	}
	
	for { //проверка флага последнего элемента 
		time.Sleep(time.Millisecond) //снимем нагрузку с этого цикла
		if slice[1][lenS - 1] == 1 {
			close(clos)//сигнал горутинам прекращать
			break
		}
	}
	duration1 := time.Since(start1)
	fmt.Println("1 Время выполнения: ", duration1)
}