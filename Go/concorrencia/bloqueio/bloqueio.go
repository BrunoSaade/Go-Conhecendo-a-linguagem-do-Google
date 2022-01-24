package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operação bloqueante
	fmt.Println("Só depois que for lido!")
}

func main() {
	c := make(chan int)

	go routine(c)

	fmt.Println(<-c)
	fmt.Println("Foi lido")
	fmt.Println(<-c) //Deadlock
	fmt.Println("Fim")
}
