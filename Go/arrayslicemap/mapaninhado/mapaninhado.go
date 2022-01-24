package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{

		"A": {
			"Arnaldo": 12345,
			"Antonio": 78965,
		},
		"B": {
			"Bruno":   58472,
			"Beatriz": 14253,
		},
		"C": {
			"Carol":   58632,
			"Carlos0": 65732,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
