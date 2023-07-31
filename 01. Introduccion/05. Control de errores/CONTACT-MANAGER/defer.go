package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	defer fmt.Println(3) //se agregan a una pila, 3 2 1
	defer fmt.Println(2) //primero en llegar, ulimo en ejecutarse
	defer fmt.Println(1)
	*/

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Alex Roel"))
	if err != nil {
		fmt.Println(err)
		//file.Close() se repite codigo, para eso se usara defer
		return
	}

	// file.Close()
	
}