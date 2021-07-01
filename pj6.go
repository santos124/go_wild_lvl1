package main
import (
	"fmt"
	"math/rand"
	"time"
	"context"
)
func contextFn() {
	forever := make(chan struct{})
    ctx, cancel := context.WithCancel(context.Background())

    go func(ctx context.Context) {
        for {
            select {
            case <-ctx.Done():  // if cancel() execute
                forever <- struct{}{}
                return
            default:
                fmt.Println("Работает функция по тестированию context")
            }

            time.Sleep(500 * time.Millisecond)
        }
    }(ctx)

    go func() {
        time.Sleep(3 * time.Second)
        cancel()
    }()

    <-forever
    fmt.Println("функция по тестированию context завершает свою работу")
}

func main() {
	fmt.Println("1 способ остановки использование context() из пакета 'context'")
	contextFn()
	fmt.Println("2 способ использование каналов без сторонних библиотек")
	stop		:= make(chan bool)
	stopTimer 	:= make(chan bool)
	start		:= make(chan bool)
	chInt		:= make(chan int)
	input		:= ""
	

	go func() {
		for {
			select {
			case x := <-chInt:
				fmt.Println(x)
			case <- stop:
				fmt.Println("принтер остановлен")
				return
			default:
			}
		}
	}()
	go func() {

		for {
			select {
			case <-start:
				for {
					select {
					case <-stopTimer:
						fmt.Println("таймер остановлен")
						stop <- true
						return
					default :
						chInt <- rand.Intn(100500)
						time.Sleep(time.Duration(rand.Intn(666)) * time.Millisecond)
					}
				}
			default:
			}
		}
	}()
	fmt.Println("нажмите одну/несколько из клавиш с буквами/цифрами и Ентер для старта, для последующей остановки повторить операцию")
	fmt.Scan(&input)
	if input != "" {
		start <- true
		input = ""
	}
	fmt.Scan(&input)
	if input != "" {
		stopTimer <- true
		input = ""
	}
	time.Sleep(5 * time.Millisecond)
}