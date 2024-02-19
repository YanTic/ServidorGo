package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Con esto se le dice al cliente que espera recibir datos en formato JSON
	w.Header().Set("Content-Type", "application/json")

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("Valores enviados por JSON %v", u)

	if u.Username == "" || u.Password == "" {
		http.Error(w, "Usuario y contraseña son obligatorios", http.StatusBadRequest)
	} else if u.Username != "julian" {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
	} else {
		tokenString, err := CreateToken(u.Username)
		if err != nil {
			http.Error(w, "Token couldn't be created", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
	}
}

func SaludoHandler(w http.ResponseWriter, r *http.Request) {
	// Con esto se le dice al cliente que espera recibir datos en formato JSON
	w.Header().Set("Content-Type", "application/json")
	usuario := r.URL.Query().Get("nombre")

	// Se Verifique que la solicitud contenga una cabecera Authorization con un JWT valido
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Falta la cabecera Authorization")
		return
	}
	tokenString = tokenString[len("Bearer "):] // Se elimina el Bearer del token (Bearer eyJhbGciOiJ...)

	// Se verifica el token enviado
	err := VerifyToken(tokenString, usuario)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token invalido | ", err)
		return
	}

	// Se verifica si el cliente envió un usuario
	if usuario == "" {
		http.Error(w, "Solicitud no valida: El nombre es obligatorio", http.StatusNotFound)
		return
	} else {
		response := fmt.Sprintf("Hola %s", usuario)
		fmt.Fprintln(w, response)
		w.WriteHeader(http.StatusOK)
	}
}

// func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	tokenString := r.Header.Get("Authorization")

// 	if tokenString == "" {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Missing authorization header")
// 		return
// 	}
// 	tokenString = tokenString[len("Bearer "):]

// 	err := VerifyToken(tokenString)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Invalid token")
// 		return
// 	}

// 	fmt.Fprint(w, "Welcome to the protected area")
// }
