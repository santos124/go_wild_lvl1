// Написать программу, которая в рантайме 
// способна определить тип переменной 
// — int, string, bool, channel 
// из переменной типа interface{}.
package main

import "fmt"

func checkType(I interface{}) {
	switch i := I.(type) {
	case int:
		fmt.Printf("%T, %v\n", i, i)
	case string:
		fmt.Printf("%T, %v\n", i, i)
	case bool:
		fmt.Printf("%T, %v\n", i, i)
	case chan int:
		fmt.Printf("%T, %v\n", i, i)
	default:
		fmt.Printf("хз %T\n", i)
	}
}

func main() {
	
	a := 12
	fmt.Scan(&a)
	checkType(a)
	checkType(21)
	checkType("hello")
	checkType(true)
	ch := make(chan int, 13)
	checkType(ch)
}