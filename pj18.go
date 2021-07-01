// Что выведет данная программа и почему?


// func someAction(v []int8, b int8) {
//   v[0] = 100
//   v = append(v, b)
// }

// func main() {
//   var a = []int8{1, 2, 3, 4, 5}
//   someAction(a, 6)
//   fmt.Println(a)
// }

package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100 //в 0й индекс среза присваивается сотка
	fmt.Printf("%p\n", v)
	v = append(v, b) //склеивается срез v и элемент b и в переменную v присваивается новый адрес новой переменой среза
	fmt.Printf("%p\n", v)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", a)
	someAction(a, 6)//поэтому в старом адресе поменялся только элемент среза по индексу 0
	fmt.Printf("%p\n", a)
	fmt.Println(a)
}