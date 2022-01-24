package main

import "fmt"

func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

func main() {
	a, b := troca(2, 3)
	fmt.Println(a, b)

	segundo, primeiro := troca(7, 1)
	fmt.Println(segundo, primeiro)
}
