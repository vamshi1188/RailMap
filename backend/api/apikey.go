package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var apiKey string

func HandleGetKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"api_key":"` + apiKey + `"}`))
}

func LoadAPI() {
	keyBytes, err := os.ReadFile("api.txt")
	if err != nil {
		log.Fatal("Error reading api.text:", err)
	}
	apiKey = strings.TrimSpace(string(keyBytes))
	fmt.Println(apiKey)
}
