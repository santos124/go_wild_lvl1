package main
import (
	"fmt"
	"math/rand"
	"time"
	"context"
)
func contextFn() {
	forever := make(chan struct{})
    ctx, cancel := context.WithCancel(context.Background()) //ctx канал для получения данных об остановке, при запуске cancel() отправляется сигнал в ctx

    go func(ctx context.Context) {
        for {
            select {
            case <-ctx.Done():  // остановка
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
        cancel()											//запуск cancel
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
			case x := <-chInt:										//стандартная работа горутины
				fmt.Println(x)
			case <- stop:											//работа горутины при появлении данных в stopTimer
				fmt.Println("принтер остановлен")
				return
			default:
			}
		}
	}()
	go func() {

		for {
			select {
			case <-start:											//запуск при появлении данных в канале start
				for {
					select {
					case <-stopTimer: 								//работа горутины при появлении данных в stopTimer
						fmt.Println("таймер остановлен")
						stop <- true 								// отправка в канал stop для остановки принтера
						return
					default : 										//стандартная работа горутины
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
		start <- true												//отправка данных в канал start для начала
		input = ""
	}
	fmt.Scan(&input)
	if input != "" {
		stopTimer <- true											//отправка данных в канал stop для конца
		input = ""
	}
	time.Sleep(5 * time.Millisecond)
}