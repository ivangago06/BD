package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func ShowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Crea un cliente Resty
	client := resty.New()

	// Realiza una solicitud GET a la API
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get("https://random-data-api.com/api/v2/users?size=2&is_xml=true")
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Convertir la respuesta a JSON
	var jsonResponse interface{}
	err = json.Unmarshal(resp.Body(), &jsonResponse)
	if err != nil {
		fmt.Println("Error:", err)
	}
	json.NewEncoder(w).Encode(&jsonResponse)
	// Imprimir la respuesta JSON
}
