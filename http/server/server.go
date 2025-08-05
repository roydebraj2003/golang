package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	var age int
	fmt.Println("ENTER YOUR AGE :")
	fmt.Scan(&age)
	msg := fmt.Sprintf("You are %d years old", age)

	w.Header().Set("Content-Type", "application/json")

	response := Message{
		Name:    "Roy",
		Message: msg,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	PORT := 8080
	http.HandleFunc("/", writeHandler)
	fmt.Printf("Server is listening on %v\n", PORT)
	http.ListenAndServe(":8080", nil)

}
