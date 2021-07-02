/*Какой результат выполнения данного кода и почему?


func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}*/
package main
import "fmt"

func main() {
	slice := []string{"a", "a"}
  
	func(slice []string) { //в этот момент создается срез с именем существуещего среза вне анонимной функции
	   slice = append(slice, "a")//но голанг настроен так что внутри анонимной функции приоритет одноименной переменной созданной внутри неё выше
	   slice[0] = "b"
	   slice[1] = "b"
	   fmt.Print(slice)//это печатается slice который был создан в этой анонимной функции
	}(slice)
	fmt.Print(slice)//а это печатается срез который создан вне анонимной функции
}

  
