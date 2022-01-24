package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[12345] = "Bruno"
	aprovados[54321] = "Jose"
	aprovados[89654] = "Rafa"
	aprovados[56921] = "Dag"

	for id, nome := range aprovados {

		fmt.Printf("%s [Id: %d]\n", nome, id)

	}

	fmt.Println(aprovados[12345])
	delete(aprovados, 12345)

	if aprovados[12345] == "" {
		fmt.Println("Map inexistente")
	}

}
