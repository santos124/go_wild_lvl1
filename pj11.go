// Написать пересечение двух неупорядоченных массивов.
package main
import (
	"fmt"
)
func main() {
	Am := []int{1, 2, 3, 4, 5, 5}
	Bm := []int{4, 5, 5, 6}
	Cm := []int{}
	for i := range Am { //перебор 1го массива
		for j := range Bm { //перебор 2го массива
			if Am[i] == Bm[j] {//если совпадение
				
				temp := true
				
				for k := range Cm {//перебор результативного массива чтоб не повторятся
					
					if Am[i] == Cm[k] {
						temp = false
						break
					}
				}
				if temp || len(Cm) == 0 {//если нет повторов в результативном массиве то пропишем
					Cm = append(Cm, Am[i])
				}
			}
		}
	}
	fmt.Printf("%v %v %v", Am, Bm, Cm)
}