//Написать программу, которая 
// перемножает, делит, складывает, вычитает 
// 2 числовых переменных a,b, значение которые > 2^20.
package main
import (
	"fmt"
	"math/big"
)
func main() {
	a := big.NewInt(1200000000000000000)
	b := big.NewInt(1200000000000000000)
	
	res := big.NewInt(1)
	res = res.Mul(a, b)
	fmt.Printf("a=%v, b=%v, a*b=%v\n", a, b, res)
	res = res.Div(a, b)
	fmt.Printf("a=%v, b=%v, a/b=%v\n", a, b, res)
	res = res.Add(a, b)
	fmt.Printf("a=%v, b=%v, a+b=%v\n", a, b, res)
	res = res.Sub(a, b)
	fmt.Printf("a=%v, b=%v, a-b=%v\n\n", a, b, res)
	
}