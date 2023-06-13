package main

import "fmt"

// estructura
type Persona struct {
	nombre string
	edad int
	correo string
}

func main(){
	// instancia de persona
	//var p Persona
	// si no asignamos valores, saldrian string vacio o cero para los int
	
	//p.nombre = "Dariel"
	//p.edad = 25
	//p.correo = "dari@correo.com"
	//lo anterior, pero mas rapido
	p := Persona { "Dariel", 25, "dari@correo.com"}
	p.edad = 30
	fmt.Println(p)
	fmt.Println(p.nombre)
	p2 := Persona { "Ju", 40, "ju@correo.com"}

	fmt.Println(p2)
}