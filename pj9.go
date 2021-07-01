// Написать конвейер чисел. Даны 2 канала - 
// в первый пишутся числа из массива, во второй 
// пишется результат операции 2*x, после 
// чего данные выводятся в stdout.
package main
import (
	"fmt"
)
func Duble(x int) int {
	return 2 * x
} 
func main() {
	arr := []int{1, 2, 3, 4, 5}
	in := make(chan int)
	out := make(chan int)
	//запуск горутины для чтения из массива в канал
	go func() {
		for i := range arr {
			in <- arr[i]
		}
	}()
	//запуск горутины для вызова Duble() и удвоения числа
	go func() {
		for i := 0; i < len(arr); i++ {
			out <- Duble(<-in)
		}
	}()	
	//вывод в stdout
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-out)
	}
}