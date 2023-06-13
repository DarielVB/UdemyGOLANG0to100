package main

import "fmt"

func main() {
	
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles",
		"Jueves", "Viernes", "Sabado"}

	diaRebanada := diasSemana[0:5]
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Otro Dia")
	fmt.Println(diaRebanada)
	// aqui se crea la rebanada sin contar el dos, el primero es de 0 a 1, y el segundo de 3 al final
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	nombres := make([]string, 5, 10)
	nombres[0] = "Dariel"
	fmt.Println(nombres)

	rebanada1 := []int{1,2,3,4,5}
	rebanada2 := make([]int, 5)
	// rebanada 2 a rebanada 1
	copy(rebanada1, rebanada2)

	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
	

}