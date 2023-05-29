//sintaxis de go

package main //Todos los archivos de Go deben pertenecer a un paquete

import "fmt" //importar los paquetes necesarios

var noPublica string
var Publica string //Modificadores de acceso Minuscula: privado, Mayuscula Publica

func main() { //Funcion main es el punto de entrada de una aplicacion Go
	saludo := "Hola Mundo"  //Variable tipo Python
	var saludo2 = "Hola Go" //Variable tipo JavaScript

	fmt.Println(saludo, saludo2) //No se usa punto y coma
}
