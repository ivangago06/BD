package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Content string
}

func GetDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	file, err := os.Open("archivos/archivo.txt")

	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return
	}
	defer file.Close()

	contenido, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Error al leer el archivo: ", err)
		return
	}

	encode := base64.StdEncoding.EncodeToString(contenido)

	response := Response{Content: encode}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		fmt.Println("Error al codificar la respuesta JSON: ", err)
		return
	}

	decode, err := base64.StdEncoding.DecodeString(encode)

	if err != nil {
		fmt.Println("Error al decodificar la respuesta JSON: ", err)
		return
	}

	str := string(decode)
	fmt.Println("El valor es: ", str)

	json.NewEncoder(w).Encode(jsonResponse)

}
