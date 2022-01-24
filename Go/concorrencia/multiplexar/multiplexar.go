package main

import (
	"fmt"

	"github.com/brunosaade/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Title("https://www.abordagemnoticias.com/", "https://www.queroquitar.com.br/"),
		html.Title("https://letscode.com.br/", "https://www.udemy.com/"),
	)
	fmt.Println(<-c, "|||", <-c, "|||", <-c, "|||", <-c)
}
