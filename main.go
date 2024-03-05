// main.go
package main

import (
	"log"
	"net/http"
	"test_go/handlers"
	"test_go/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Подключение к базе данных PostgreSQL
	_, err := utils.ConnectToPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer utils.ClosePostgres()

	// Создание роутера с использованием Gorilla Mux
	r := mux.NewRouter()

	// Регистрация обработчиков
	handlers.RegisterRoutes(r) // Используйте эту строку вместо handlers.RegisterGoodsHandlers(r)

	// Запуск HTTP-сервера
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
