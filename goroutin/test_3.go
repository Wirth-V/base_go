package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// Записываем сумму в c
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// Инициализируем канал
	c := make(chan int)

	// Делим слайс пополам:
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// Читаем из c
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
