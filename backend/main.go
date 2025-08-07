package main

import (
	"encoding/json"
	"net/http"
)

type api struct {
	addr string
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	user := User{
		Name:  "Roy",
		Email: "roydebrajworks@gmail.com",
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	api := &api{
		addr: ":8080",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", api.getUsersHandler)
	srv := http.Server{
		Addr:    api.addr,
		Handler: mux,
	}
	srv.ListenAndServe()
}
