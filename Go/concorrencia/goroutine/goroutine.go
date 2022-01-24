package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Bruno", "Pq vc n fala cmg?", 3)
	// fale("Rafa", "Só posso falar depois de vc!", 1)

	// go fale("Bruno", "Ei.....", 500)
	// go fale("Rafael", "Opa.....", 500)

	go fale("Maria", "Entendi!!", 10)
	fale("João", "Parabéns", 5)

}
