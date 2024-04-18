package main

import (
	"encoding/json" // Se usa para decodificar datos JSON
	"fmt"           // Se usa para formatear la salida
	"log"           // Se usa para registrar errores
	"net/http"      // Se usa para realizar peticiones HTTP
)

func main() {
	// Define la URL para obtener datos de usuarios de JSONPlaceholder
	const url = "https://jsonplaceholder.typicode.com/users"

	// Envía una petición GET HTTP a la URL y almacena la respuesta
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al obtener datos:", err) // Imprime mensaje de error
		return
	}
	defer resp.Body.Close() // Asegurar que el cuerpo de la respuesta se cierre incluso en caso de errores

	// Define una lista de mapas para almacenar los datos de usuario decodificados (se puede personalizar)
	var users []map[string]interface{}

	// Decodifica la respuesta JSON del cuerpo en la lista users
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		log.Fatalln("Error al decodificar JSON:", err) // Registra un error fatal para fallas de decodificación
	}

	// Itera a través de la lista de datos de usuario (mapas)
	fmt.Println("** Lista de Usuarios: **")
	for _, user := range users {
		// Accede a los datos del usuario usando las claves en el mapa
		fmt.Println("ID:", user["id"])
		fmt.Println("Nombre:", user["name"])
		// Imprime otros datos de usuario (correo electrónico, etc.) según sea necesario
	}
}
