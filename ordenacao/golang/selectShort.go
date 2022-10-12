package main

import "fmt"

func main() {
	var lista = []int64{5, 3, 2, 1, 0}

	for i := 0; i < len(lista); i++ {
		for j := i + 1; j < len(lista); j++ {
			if lista[j] < lista[i] {
				temp := lista[i]
				lista[i] = lista[j]
				lista[j] = temp
			}
		}
	}
	fmt.Println("resultado:", lista)
}
