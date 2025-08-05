package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data Message
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	fmt.Printf("The response is: %s, %s\n", data.Name, data.Message)
}
