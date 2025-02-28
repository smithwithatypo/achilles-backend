package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "John Doe"}
	json.NewEncoder(w).Encode(user)
}
