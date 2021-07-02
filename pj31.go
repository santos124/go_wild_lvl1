//Написать программу нахождения расстояния
//  между 2 точками, которые представление в 
//  виде структуры Point с инкапсулированными
//   параметрами x,y и конструктором.
package main
import (
	"fmt"
	"math"

)
type Point struct { //описание структуры
	x float64//поля приватные(с маленькой буквы), т.е. не видны из других гошных файлов
	y float64
}

func NewPoint(x, y float64) *Point { //конструктор публичный, т.к. с большой буквы
	dot := new(Point)
	dot.x = x
	dot.y = y
	return dot
}
func main() {
	a := NewPoint(0, 0)//создание элемента типа Point
	b := NewPoint(3.2, 4.1)
	c := math.Sqrt((a.x - b.x) * (a.x - b.x) + (a.y - b.y) * (a.y - b.y))// квадрат гипотенузы равен сумме квадратов катетов
	fmt.Println(a, b, c)
}