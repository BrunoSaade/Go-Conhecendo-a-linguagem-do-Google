package main

import (
	"fmt"
	"time"

	"github.com/brunosaade/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Title(url1)
	c2 := html.Title(url2)
	c3 := html.Title(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
	}
}

func main() {
	resposta := oMaisRapido(
		"https://www.abordagemnoticias.com/",
		"https://www.queroquitar.com.br/",
		"https://www.udemy.com/",
	)

	fmt.Println(resposta)
}
