// К каким негативным последствиям может привести данный кусок кода и как это исправить?


// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }
package main
import "fmt"

var justString string //думаю глобальные переменные по необходимости стоит избегать чтоб не было задвоения имен

func createHugeString(l int) string{
	x := make([]byte, l, l)
	for i := range x {
		x[i] = 'w'
	}
	return string(x)

}
func someFunc() {
  v := createHugeString(1 << 10) //думаю что битовые операции выглядят неоч и могут запутать кодеров
  justString = v[:100] //был ли смысл создавать  строку на 1024 символа если нужно только 100, если это дело буде запускаться у каждого клиента и в горутинках и таких горутин/клиентов пара тысяч в секунду, нужно будет ехать за новыми серваками
}

func main() {
	someFunc()
	fmt.Println(justString)
}

