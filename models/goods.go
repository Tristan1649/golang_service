// models/goods.go
package models

import "time"

// Goods представляет товар в системе.
type Goods struct {
	ID          int       `json:"id"`           // Уникальный идентификатор товара
	ProjectID   int       `json:"project_id"`   // Идентификатор проекта, к которому привязан товар
	Name        string    `json:"name"`         // Название товара
	Description string    `json:"description"`  // Описание товара
	Priority    int       `json:"priority"`     // Приоритет товара
	Removed     bool      `json:"removed"`      // Флаг, указывающий на удаление товара
	CreatedAt   time.Time `json:"created_at"`   // Время создания товара
}
