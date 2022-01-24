package main

import "fmt"

func main() {

	funcsESalario := map[string]float64{
		"Bruno":  1500.00,
		"Rafael": 3500.55,
		"Jose":   15000.00,
		"Dag":    20000.00,
	}

	funcsESalario["Luiz"] = 250.00
	delete(funcsESalario, "Inexistente")

	for nome, salario := range funcsESalario {
		fmt.Println(nome, salario)
	}

}
