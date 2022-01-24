package main

import "fmt"

func init() { //Funções init sempre atuam antes da main
	fmt.Println("Inicializando")
}

func main() {
	fmt.Println("Main...")
}
