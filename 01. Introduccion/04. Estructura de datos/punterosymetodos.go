package main

import "fmt"

type Persona struct {
	nombre string
	edad int
	correo string
}
//metodo con la estructura de persona
func (p *Persona) diHola(){
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {
	var x int = 10
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)
	//var p *int = &x

	// tienen la misma referencia en memoria
	//fmt.Println(&x)
	//fmt.Println(p)
	p := Persona{"Dariel", 25, "dari@correo.com"}
	p.diHola() //ES UN metodo de Persona, pues le pertenece a esa estructura
}

func editar(x *int){
	*x = 20
}