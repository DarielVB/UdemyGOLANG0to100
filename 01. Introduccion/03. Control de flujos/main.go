package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//t := time.Now()
	//hora := t.Hour()

	//fmt.Println(t.Hour())

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println(t.Hour(), "¡Mañana!")
	case t.Hour() < 17:
		fmt.Println(t.Hour(), "¡Tarde!")
	default:
		fmt.Println(t.Hour(), "¡Noche!")
	}
	//os := runtime.GOOS
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
		
	case "linux":
		fmt.Println("Go run -> Linux")
	
	case "darwin":
		fmt.Println("Go run -> macOS")
	
	default:
		fmt.Println("Go run -> another Os")
	}
	//en go no es necesario break

	// bucle for
	//var i int
	for i := 1; i <= 10; i++ {
		
		if i == 5 {
			//break
			continue //salta a la siguiente iteracion, no se ejecutara el print
		}
		fmt.Println(i)
	}
	saludo := hello("Dariel")
	fmt.Println(saludo)
	sum, mul := calc(4, 5)
	fmt.Println("La suma es: ", sum)
	fmt.Println("La multi es: ", mul)

}
// nuestra funcion
func hello(name string) string {
	fmt.Println("Holiwis, ", name)
	return "holiwis, " + name
}

func calc(a, b int)(sum, mul int) {
	sum = a+b
	mul = a*b
	//return sum, mul
	return
}