package main

import (
	"fmt"
	"time"
)

// cd
func main() {

	intCh := make(chan int)

	go func() {
		intCh <- 5 // блокировка, пока данные не будут получены функцией main
		fmt.Println("Go routine starts")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(<-intCh) // получение данных из канала

	time.Sleep(2 * time.Second)

	fmt.Println("The End")
}
