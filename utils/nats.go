// utils/logging.go
package utils

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

// LogToClickhouse отправляет лог в ClickHouse через Nats
func LogToClickhouse(subject string, action string, data interface{}) {
	// Эта функция отправляет лог в ClickHouse через Nats.
	// Она принимает тему (subject), действие (action) и данные (data),
	// создает строку лога и отправляет ее по Nats.

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS:", err)
		return
	}
	defer nc.Close()

	logData := fmt.Sprintf("Subject: %s, Action: %s, Data: %+v", subject, action, data)
	nc.Publish(subject, []byte(logData))
}
