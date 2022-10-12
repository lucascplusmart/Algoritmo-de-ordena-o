package main

import "fmt"

func main() {
	var lista = []int{4, 2, 1, 3}

	for j := 0; j < len(lista); j++ {
		for i := 0; i < len(lista); i++ {
			if i+1 >= len(lista) {
				break
			}

			if lista[i] > lista[i+1] {
				aux := lista[i]
				lista[i] = lista[i+1]
				lista[i+1] = aux
			}
		}
	}

	fmt.Println("resultado:", lista)
}
