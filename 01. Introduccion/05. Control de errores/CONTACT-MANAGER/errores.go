package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int)(int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	/*
	str := "123" // error si ponemos 123f
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Numero:", num)
	*/
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Resultado:", result)
}