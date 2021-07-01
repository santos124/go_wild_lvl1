package main
import (
	"fmt"
	"sync"
)
func MapperEven(m map[int]int, n int, mu *sync.Mutex) {
	for i := 0; i < n * 2; i = i + 2 {
		mu.Lock() // теперь горутины в которых произошел вызов метода Lock()
		m[i] = i
		mu.Unlock()//будут ждать пока не произойдет вызов Unlock()
	}
}
func MapperOdd(m map[int]int, n int, mu *sync.Mutex) {
	for i := 1; i < n * 2; i = i + 2 {
		mu.Lock()
		m[i] = i
		mu.Unlock()
	}
}
func main() {
	m := make(map[int]int)
	mu := new(sync.Mutex) //создаем элемент sync.Mutex
	go MapperEven(m, 3, mu) // при вызове горутин отправляем mu
	go MapperOdd(m, 3, mu)
	var a int
	fmt.Println("введите любой символ")
	fmt.Scan(&a)//притормаживаем мейн чтобы горутины успели
	fmt.Printf("%v, %T\n", m, m)
}
//Реализовать конкурентную запись в map.