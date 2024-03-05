// utils/redis.go
package utils

import (
	"log"

	"github.com/go-redis/redis/v8"
)

// ConfigRedis описывает конфигурацию подключения к Redis
type ConfigRedis struct {
	Addr     string
	Password string
	DB       int
}

// InitRedis создает соединение с сервером Redis
func InitRedis(config ConfigRedis) *redis.Client {
	// Эта функция инициализирует подключение к серверу Redis и возвращает клиент
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	return rdb
}

// CloseRedis закрывает соединение с сервером Redis
func CloseRedis(rdb *redis.Client) {
	// Эта функция закрывает соединение с сервером Redis и логирует ошибку, если она есть
	err := rdb.Close()
	if err != nil {
		log.Println("Error closing Redis connection:", err)
	}
}
