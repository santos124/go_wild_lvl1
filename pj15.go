// Поменять местами два числа без создания временной переменной.
package main
import "fmt"
func main() {
	a, b := 100500, 1
	fmt.Printf("a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("a=%d, b=%d\n", a, b)
}