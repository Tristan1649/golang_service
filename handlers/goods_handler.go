// handlers/goods_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test_go/models"
	"test_go/utils"
)

// CreateGoods обработчик создания записи в таблице Goods.
// Принимает HTTP-запрос с данными о товаре, добавляет запись в базу данных и отправляет ответ с созданным товаром.
func CreateGoods(w http.ResponseWriter, r *http.Request) {
	var goods models.Goods
	err := json.NewDecoder(r.Body).Decode(&goods)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// Логика создания записи в базе данных

	// Пример использования функции для выполнения SQL-запроса
	_, err = utils.ExecuteSQL("INSERT INTO Goods (Name, Description, Priority) VALUES ($1, $2, $3) RETURNING Id, Project_id, Name, Description, Priority, Removed, Created_at",
		goods.Name, goods.Description, goods.Priority)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating goods: %v", err), http.StatusInternalServerError)
		return
	}
	
	// Логика отправки записи в ClickHouse через Nats
	utils.LogToClickhouse("clickhouse_log", "insert", goods)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(goods)
}

// GetGoodsList обработчик получения списка записей из таблицы Goods.
// Возвращает список всех товаров из базы данных.
func GetGoodsList(w http.ResponseWriter, r *http.Request) {
	// Логика получения списка записей из базы данных

	// Пример использования функции для выполнения SQL-запроса
	rows, err := utils.ExecuteSQL("SELECT Id, Project_id, Name, Description, Priority, Removed, Created_at FROM Goods")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting goods list: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var goodsList []models.Goods
	for rows.Next() {
		var goods models.Goods
		err := rows.Scan(&goods.ID, &goods.ProjectID, &goods.Name, &goods.Description, &goods.Priority, &goods.Removed, &goods.CreatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		goodsList = append(goodsList, goods)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(goodsList)
}
