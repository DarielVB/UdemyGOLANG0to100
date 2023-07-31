package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("main():") //para poner prefijo, todo se va a ver acompañado de esto main(): hora - mensaje
	file, err := os.OpenFile("info.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644) // 0644 es un valor octal para los permisos de los archvis, lectura y escritura

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Print("!Oye, soy un logogogogogogogogo!")

	defer file.Close()
	log.Print("Este es un mensaje de registro")
	log.Print("Este es otro mensaje de registro")
	
	log.Fatal("Holiwis") //imprimir error con tiempo y fecha, y detendra el programa
	log.Panic("Hago lo mismo que fatal")

	//registrar todos los registros en un archivo
}