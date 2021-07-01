//Реализовать композицию структуры Action от родительской структуры Human.
package main

import "fmt"

//описание структуры
type Human struct {
	Name	string
	Age		int
}

//описание структуры дочерней от Human
type Action struct {
	Human
	Direct string
}

func main() {
	Intership := Action{} //элемент структуры Action
	Intership.Direct = "Golang"
	SanHum := Human{"Santos", 28} // элемент структуры Human
	Intership.Human = SanHum // поле Human структуры Intership типа Action получает значение элемента SanHum структуры Human
	fmt.Printf("%v", Intership)
}