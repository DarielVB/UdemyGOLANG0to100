// guardar en un archivo los contactos

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

)

// estructura de contactos
type Contact struct {
	Name  string `jeson:"name"`
	Email string `jeson:"email"`
	Phone string `jeson:"name"`
}

// Guarda contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file) // de tipo file lo vuelve a json
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

// Cargar contactos desde un archivo json

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	// codificar de json a slas (no se si se llama asi)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	// Slice de contactos
	var contacts []Contact

	// Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)

	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		// Mostrar opciones al usuario
		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Eliger una opcion: ")

		// Leer la opcion del usuario
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion: ", err)
			return
		}

		// Manejar la opcion del usuario
		switch option {
		case 1:
			// Ingresar y crear contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			// agregar un contacto a Slice
			contacts = append(contacts, c)

			//Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar contacto:", err)
			}
		case 2:
			// Mostrar todos los contactos
			fmt.Println("=================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=================================================")

		case 3:
			// Salir del programa
			return
		default: 
			fmt.Println("Opcion invalida")
		}

	}

}
