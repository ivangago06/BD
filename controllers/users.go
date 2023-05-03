package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func ShouUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	client := resty.New()

	resp, err := client.R().SetHeader("Content-Type", "application/json").Get("https://random-data-api.com/api/v2/users?size=2&is_xml=true")

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	var jsonResponse interface{}

	err = json.Unmarshal(resp.Body(), &jsonResponse)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	json.NewEncoder(w).Encode(&jsonResponse)
}
