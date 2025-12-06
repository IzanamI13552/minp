package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("users.json")
		if err != nil && !os.IsNotExist(err) {
			http.Error(w, "Ошибка чтения файла", 500)
			return
		}
		var users []User
		if len(data) > 0 {
			if err := json.Unmarshal(data, &users); err != nil {
				http.Error(w, "Ошибка парсинга JSON", 500)
				return
			}
		} else {
			users = []User{}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := map[string]string{
			"time": time.Now().Format("2006-01-02 15:04:05"),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(currentTime)
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
