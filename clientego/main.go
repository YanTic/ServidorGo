package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Credenciales struct {
	Username string
	Password string
}

func main() {
	servidorURL := os.Getenv("SERVIDORGO")
	cliente := &http.Client{}
	// URL := "http://localhost:80"

	if servidorURL == "" {
		fmt.Println("Variable de entorno SERVIDORGO no configurada")
		return
	}

	// Se crea la info que se va mandar al servidor
	user := "usuario." + strconv.Itoa(rand.Intn(100))
	pass := "contrasenia." + strconv.Itoa(rand.Intn(100))
	payload := Credenciales{Username: user, Password: pass}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Request a la ruta /login
	request, err := http.NewRequest("POST", servidorURL+"/login", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := cliente.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("---Imprimiendo respuesta del servidor (Ruta /login)---")
	fmt.Println(string(body))

	// Request a la ruta /saludo , ahora usando datos por URL y no por JSON
	datos := url.Values{}
	datos.Add("nombre", user)
	request, err = http.NewRequest("GET", servidorURL+"/saludo?"+datos.Encode(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Authorization", "Bearer "+string(body))

	response, err = cliente.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("---Imprimiendo respuesta del servidor (Ruta /saludo)---")
	fmt.Println(string(body))
}
