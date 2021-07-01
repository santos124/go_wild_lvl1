//Что выводит данная программа и почему?
package main
import "fmt"
func update(p *int) {
	b := 2
	fmt.Printf("b=%d|%p, p=%d|%p, \n", b, &b, *p, p)
	p = &b //тут p получает другой адрес и значение по этому адресу чтоб поменялось значение а в мейне следует прописать значение по адресу *p = b
	fmt.Printf("b=%d|%p, p=%d|%p, \n", b, &b, *p, p)
}
  
func main() {
var (
	a = 1
	p = &a
)
	fmt.Println(*p)
	fmt.Printf("a=%d|%p, p=%d|%p, \n", a, &a, *p, p)
	update(p) //но адрес не вернулся и значение не поменялось, тк тут передается значение адреса а не обьект
	fmt.Printf("a=%d|%p, p=%d|%p, \n", a, &a, *p, p)
	fmt.Println(*p)
}  