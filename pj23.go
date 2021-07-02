//написмать бинарный поиск встроенными методами языка
package main
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)
func binSearch(arr []int, key int) int {   // Запускаем бинарный поиск
    l := -1                      // l, r — левая и правая границы
    r := len(arr)    
    for l < r - 1 {                // Запускаем цикл
        m := (l + r) / 2            // m — середина области поиска
        if arr[m] < key {
            l = m
		} else { 						// Сужение границ
            r = m
		}  
	}                
    return r
}
func main() {
	
	lenS := 25
	lastkey := 1
	slice := make([]int, lenS, lenS)
	for i := 0; i < lenS; i++ {
		slice[i] = rand.Intn(23)
		lastkey = slice[i]
	}
	start2 := time.Now()
	sort.Ints(slice)
	duration2 := time.Since(start2)
	fmt.Println("2 Время выполнения: ", duration2)
	res := binSearch(slice, lastkey)
	fmt.Println(slice, lastkey, res)

}