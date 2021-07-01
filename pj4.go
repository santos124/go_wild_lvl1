package main

import (
    "fmt"
    "time"
	"sync"
	"math/rand"
	"os"
)

func worker(id int, task <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
	for j := range task {
		time.Sleep(time.Second)

        fmt.Println("Для завершение введите любой/любые символы и Enter ", "worker", id, "done task", j)
	}
}

func main() {
	
	fmt.Printf("Введите число воркеров и нажмите Enter\n")
	var w int
	fmt.Scan(&w) // ввод числа воркеров
	t := 100500
    task := make(chan int) //создание канала интов
	wg := new(sync.WaitGroup)//создание элемента для синхронизации
    for i := 1; i <= w; i++ {
		wg.Add(1) // плюсуем счетчик синхронизации
        go worker(i, task, wg) //запуск отдельной горутины
    }
	go func() { //еще один запуск горутины которая ждет ввод с клавиатуры для завершения
		fmt.Printf("Для завершение введите любой/любые символы и Enter\n")
		s := "123"
		fmt.Scan(&s)
		os.Exit(0)
	}()
    for j := 1; j <= t; j++ {
		if j == t {
			j = 1
		}
		random := rand.Intn(100)
     	task <- random
    }
	
	wg.Wait() //ждем обнуления счетчика, чтобы горутина успела отработать до выхода из мейна
}
