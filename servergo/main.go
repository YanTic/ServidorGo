package main

import (
	"fmt"
	"net/http"
	"os"
	jwt "servidor/jwt"
)

type User struct {
	Usuario string
	Clave   string
}

func main() {
	usuario, _ := os.Hostname()
	router := http.NewServeMux()

	// Handler para cada ruta
	router.HandleFunc("/saludo", jwt.SaludoHandler)
	router.HandleFunc("/login", jwt.LoginHandler)
	// router.HandleFunc("/protected", jwt.ProtectedHandler)

	// Handler para rutas no conocidas
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Recurso no encontrado", http.StatusNotFound)
	})

	fmt.Println("Iniciando Servidor en http://localhost:80/ user:" + usuario)
	err := http.ListenAndServe(":80", router)
	if err != nil {
		fmt.Println("Could not start the server", err)
	}

}
