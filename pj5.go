package main

import (
    "fmt"
	"math/rand"
    "time"
	"os"
)
func main() {
	fmt.Println("Введите секунды времени работы (число)")
	var N int
	_, err := fmt.Scan(&N)
	for err != nil {
		fmt.Println("Введите секунды времени работы (число)")
		_, err = fmt.Scan(&N)
	}
	chInt := make(chan int, 1)
	go func() {										//отложенное завершение
		time.Sleep(time.Second * time.Duration(N))
		os.Exit(0)
	}()
	
	for {
		time.Sleep(time.Millisecond * time.Duration(112))
		x := rand.Intn(100500)
		chInt <- x 									//запись в канал
		fmt.Printf("Записано число %d\n", x)
		fmt.Printf("Считано число%d\n", <- chInt)	//чтение
	}
}

