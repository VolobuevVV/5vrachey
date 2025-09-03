package main

import (
	"fmt"
	"log"
	"net/http"
)

// Обработчик для главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Разрешаем запросы с любого origin (CORS)
	// В продакшене нужно указать конкретный origin вашего фронтенда!
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Отвечаем JSON-данными
	w.Header().Set("Content-Type", "application/json")

	// Простое сообщение, которое мы будем отправлять на фронтенд
	response := `{"message": "Hello from Go backend!"}`
	fmt.Fprint(w, response)
}

func main() {
	// Регистрируем наш обработчик для пути "/api"
	http.HandleFunc("/api", homeHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Backend server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
