package main

import (
	"fmt"
	"math"
	"strconv"
)

//var primerNombre, apellido, edad = "Dariel", "Vega", 25
// Declaracionm de constantes
// Los constantes deben iniciar con mayuscula
const Pi float32 = 3.14
const(
	x = 100
	y = 0b1010 //binario
	z = 0o12 //Octal
	w = 0xFF //Hexadecimal
)
//valor iota
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)
func main() {
	var integer int8 = 126 
	//-128 a 127
	var integer2 uint = 29
	//solo positivos
	var float float32 = 2.2
	//float 32 es mas pequeño que float64
	fmt.Println(integer)
	fmt.Println(integer2)
	fmt.Println(float)
	fmt.Println(math.MaxFloat64)
	fullName := "Ale Roel \t(alias \"roelcode\")\n"
	fmt.Println(fullName)
	var firstName string
	var ascii byte = 'a'
	//almacenar valos ASCII
	fmt.Println(ascii)
	s := "hola"
	fmt.Println(s[0])
	//imprime el primer valor de la cadena pero en ascii
	var r rune = '😄'
	//valor unicode
	fmt.Println((r))
	firstName = "Alex"
	var lastName string = "Dariel"
	lastName = "Alex"
	age := 26
	fmt.Println("Programa")
	fmt.Println(firstName, lastName, age)

	// Nos manda error pi = 3.1415
	/*
	var (
		primerNombre, SegundoNombre string
		edad 						int
	)
	*/
	/*
	var(
		primerNombre = "Dariel"
		apellido = "Vega"
		edad = 25
	)
	*/

	// var primerNombre, apellido, edad = "Dariel", "Vega", 25
	primerNombre, apellido, edad := "Dariel", "Vega", 25
	fmt.Println(primerNombre, apellido, edad)

	var a int 
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c,d)
	const pi = 3.141592
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)
	fmt.Println(pi)
	//Valores predeterminados
	// int o numeros es cero
	// bool es false
	var (
		defaulInt int
		defaultUInt uint
		defaulFloat float32
		defaultBool bool
		defaultString string
	)
	fmt.Println(defaulInt, defaultUInt, defaulFloat, defaultBool, defaultString)
	//Conversiones de tipos
	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)
	//no se pueden sumar, toca convertirlos

	ss := "100"
	i, _ :=strconv.Atoi(ss)

	fmt.Println(i + 1)

	n := 42

	s = strconv.Itoa(n)
	// se vuelve en str el int, pero solo ints, no se puede int8
	fmt.Println(s)

	// El paquete fmt

	fmt.Print("Otro mensaje") //no hace salto de linea

	name2 := "Dariel"
	age2 := 28

	fmt.Printf("Hola me llamo %s y tengo %d años.\n ", name2, age2)

	greeting := fmt.Sprintf("Hola me llamo %s y tengo %d años. ", name2, age2)

	fmt.Println(greeting)

	var name3 string
	var age3 int
	fmt.Print("Ingrese su nombre: ")
	//referencia de la memoria, con ampersan
	fmt.Scanln(&name3)
	fmt.Print("Ingrese su edad: ")
	//referencia de la memoria, con ampersan
	fmt.Scanln(&age3)
	//cuando imprimimos con V, es porque no savemos que valor es, y con T mayuscula, es para saber que tipo de valor es
	greeting = fmt.Sprintf("Hola me llamo %v y tengo %v años. ", name3, age3)
	fmt.Println(greeting)

	//Operadores aritmeticos y paquete Math
	a2 := 10
	b2 := 3

	a2 += 5
	b2++
	b2--
	fmt.Println(a2 + b2)
	fmt.Println(math.Pi)

	c2 := math.Pow(2, 4)
	fmt.Printf("%.1f", c2)

}