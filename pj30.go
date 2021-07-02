//Удалить i-ый элемент из слайса.
package main
import "fmt"
func deleter(arr []int, i int) []int {
	return append(arr[0 : i], arr[(i + 1) :]...)
}
func main() {
	i := 3
	arr := []int{0, 1, 2, 3, 4}
	fmt.Println(arr)
	arr = deleter(arr, i)
	fmt.Println(arr)
}