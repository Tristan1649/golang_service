// utils/postgres.go
package utils

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pgPool *pgxpool.Pool

// ConnectToPostgres подключается к серверу PostgreSQL
func ConnectToPostgres() (*pgxpool.Pool, error) {
	// Эта функция устанавливает соединение с сервером PostgreSQL
	config, err := pgxpool.ParseConfig("postgresql://user:password@localhost:5432/database")
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	pgPool = pool

	return pgPool, nil
}

// ClosePostgres закрывает соединение с PostgreSQL
func ClosePostgres() {
	// Эта функция закрывает соединение с сервером PostgreSQL
	if pgPool != nil {
		pgPool.Close()
	}
}

// ExecuteSQL выполняет SQL-запрос в PostgreSQL
func ExecuteSQL(query string, args ...interface{}) (pgx.Rows, error) {
	// Эта функция выполняет SQL-запрос в PostgreSQL и возвращает результаты
	return pgPool.Query(context.Background(), query, args...)
}
