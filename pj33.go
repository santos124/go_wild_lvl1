// Даны 2 канала - в первый пишутся рандомные числа 
// после чего они проверяются на четность
//  и отправляются во второй канал. 
//  Результаты работы из второго канала пишутся в stdout.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	done1 := make(chan struct{})// доны это каналы для ожидания выполнения горутин
	done2 := make(chan struct{})
	done3 := make(chan struct{})
	ch1 := make(chan int) 
	ch2 := make(chan string)
	go Writer(done1, ch1)//пишет псевдораедомные числа в ch1
	go Cheker(done2, ch1, ch2) //встречает числа из ch1 проверяет на четность и результат для печати шлет в ch2
	go Printer(done3, ch2)//печать в консоль
	<-done1//при закрытии канала мейн пойдет дальше
	<-done2
	<-done3
}

func Writer(done1 chan struct{}, ch1 chan int) {
	for i := 0; i < 10; i++ {//программа настроена на 10 чисел
		ch1 <- rand.Intn(100)
	}
	close(done1)
}
func Cheker(done2 chan struct{}, ch1 chan int, ch2 chan string) {
	for i := 0; i < 10; i++ {
		a := <-ch1
		if a % 2 == 0 {
			s := fmt.Sprint(a, " четное")
			ch2 <- s
			
		} else {
			s := fmt.Sprint(a, " нечетное")
			ch2 <- s
		}
		
	}
	close(done2)
}
func Printer(done3 chan struct{}, ch2 chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch2)		
	}
	close(done3)
}

