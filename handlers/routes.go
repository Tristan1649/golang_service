// handlers/routes.go
package handlers

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes регистрирует маршруты для обработчиков.
func RegisterRoutes(r *mux.Router) {
    RegisterGoodsHandlers(r) // Регистрация обработчиков товаров.
    r.HandleFunc("/api/goods", CreateGoods).Methods("POST") // Обработчик создания товара.
    r.HandleFunc("/api/goods", GetGoodsList).Methods("GET") // Обработчик получения списка товаров.
}
