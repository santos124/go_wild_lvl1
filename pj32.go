// Написать собственную функцию Sleep.
package main
import (
	"fmt"
	"time"
)

func ft_sleep(dur int){
	time1 := time.After(time.Second * time.Duration(dur))
	<-time1
}
func main() {
	fmt.Println("begin")
	ft_sleep(3)
	fmt.Println("end")
}

