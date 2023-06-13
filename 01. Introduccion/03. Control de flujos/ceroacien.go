package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// esto ya no es necesario, solo en versiones antiguas
	//rand.Seed(time.Now().UnixNano()

	jugar()

}

func jugar() {
	numAletorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10 

	for intentos < maxIntentos {
		intentos ++
		fmt.Printf("Ingresa un numero (intentos restantes: %d): ", maxIntentos - intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAletorio {
			fmt.Println("¡Felicitaciones, adivinaste el numero!")
			jugarNuevamente()
			return
		} else if numIngresado < numAletorio {
			fmt.Println("El numero a adivinar es mayor.")
		} else if numIngresado > numAletorio {
			fmt.Println("El numero a adivinar es menor.")
		}
	}

	fmt.Println("Se acabaron los intentos. El numero era:", numAletorio)
	jugarNuevamente()
}

func jugarNuevamente(){
	var eleccion string
	fmt.Print("¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar!")
	default:
		fmt.Println("Eleccion invalida. Intentalo nuevamente.")
		jugarNuevamente()
	}
}