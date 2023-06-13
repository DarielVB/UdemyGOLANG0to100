package main

import "fmt"

func main() {
	var matriz = [...]int {10, 20, 30, 40, 50}
	matriz[0] = 100
	matriz[1] = 200
	
	for i:=0; i< len(matriz); i++{
		fmt.Println(matriz[i])
	}

	for index, value := range matriz {
		fmt.Printf("indice: %d, valor: %d\n", index, value)
	}

	var matrizBi = [3][3] int {{1,2,3}, {4,5,6}, {7,8,9}}
	fmt.Println(matrizBi)
	
}