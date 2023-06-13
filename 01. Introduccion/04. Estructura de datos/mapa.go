package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors["rojo"])
	colors["negro"] = "#000000"

	fmt.Println(colors)
	// si pongo blanco, me devolvera un string vacio y me dara un falso
	// las variables declaradas en un if, solo existen en el if o en la anidacion
	if valor, ok := colors["rojo"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe este clave")
	}

	delete(colors, "azul")

	fmt.Println(colors)
	// para recorrerlo debemos usar range
	for clave, valor := range colors{
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}
}