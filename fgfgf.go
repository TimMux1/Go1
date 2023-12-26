package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var request JsonRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Некорректное JSON-сообщение", http.StatusBadRequest)
		return
	}

	if request.Message == "" {
		http.Error(w, "Некорректное JSON-сообщение", http.StatusBadRequest)
		return
	}

	fmt.Println("Получено сообщение:", request.Message)

	response := JsonResponse{
		Status:  "success",
		Message: "Данные успешно приняты",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handleRequest)

	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", nil)
}
