package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type PageData struct {
	Title string
}

func main() {
	// Стартовая страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{Title: "Добро пожаловать"}
		renderTemplate(w, "start.html", data)
	})

	// Страница генератора
	http.HandleFunc("/generator", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{Title: "Генератор UUID"}
		renderTemplate(w, "generator.html", data)
	})

	// API для генерации UUID
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var params map[string]int
			err := json.NewDecoder(r.Body).Decode(&params)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			length := params["length"]
			if length <= 0 {
				length = 36 // Default UUID length
			}

			result := generateUUID(length)
			fmt.Fprintf(w, result)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func generateUUID(length int) string {
	rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
