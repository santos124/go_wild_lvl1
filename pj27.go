//Написать программу, которая переворачивает слова 
// в строке (snow dog sun - sun dog snow).
package main
import (
	"fmt"
	"strings"
)
func main() {
	str := "snow dog sun"
	strs := strings.Split(str, " ") //делим строку на подстроки и сохраняем в массив
	strs2 := make([]string, 0)
	for i := len(strs) - 1; i >= 0; i-- {
		strs2 = append(strs2, strs[i])//сохранияем в обратном порядке в новый массив
	}
	fmt.Println(strs2)
}