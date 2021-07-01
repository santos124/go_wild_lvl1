// // Чем завершится данная программа?
// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 5; i++ {
// 	   wg.Add(1)
// 	   go func(wg sync.WaitGroup, i int) {
// 		  fmt.Println(i)
// 		  wg.Done()
// 	   }(wg, i)
// 	}
// 	wg.Wait()
// 	fmt.Println("exit")
//   }
//ошибкой, тк wg необходимо создавать с помощью new(), и горутина принимает указатель на элемент sync.WaitGroup
//думаю методы add/done/wait настроены на то что wg будет указателем
package main
import (
	"fmt"
	"sync"
)
func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
