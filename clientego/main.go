package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type Credenciales struct {
	Username string
	Password string
}

func main() {
	// servidorURL := os.Getenv("SERVIDORGO")
	cliente := &http.Client{}
	URL := "http://localhost:80/login"

	// if servidorURL == "" {
	// 	fmt.Println("Variable de entorno SERVIDORGO no configurada")
	// 	return
	// }

	// Se crea la info que se va mandar al servidor
	user := "usuario." + strconv.Itoa(rand.Intn(100))
	pass := "contrasenia." + strconv.Itoa(rand.Intn(100))
	payload := Credenciales{Username: user, Password: pass}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonPayload))
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("---Imprimiendo respuesta del servidor---")
	fmt.Println(string(body))
}
