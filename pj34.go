// Написать программу, которая проверяет, 
//что все символы в строке уникальные.
package main
import "fmt"
func Unique(str string) bool {
	for i := range str {
		for j := range str {
			if i != j && str[i] == str[j] { //если находит такой же символ в других ячейках значит не уникальные символы
				return false
			}
		}
	}
	return true
}
func main() {
	str1 := "aa1234"
	str2 := "a12345"
	fmt.Println(Unique(str1))
	
	fmt.Println(Unique(str2))
}