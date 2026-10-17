package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

// Генерируем случайный ID из 8 байт
func NewID16() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}

// Записываем основные данные HTTP-запроса
func LogRequest(r *http.Request) {
	fmt.Printf("[%s] %s %s %s\n",
		time.Now().Format(time.RFC3339),
		r.RemoteAddr,
		r.Method,
		r.URL.Path,
	)
}

func LogInfo(msg string) {
	fmt.Printf("[INFO] %s %s\n",
		time.Now().Format(time.RFC3339),
		msg,
	)
}

func LogError(msg string) {
	fmt.Printf("[ERROR] %s %s\n",
		time.Now().Format(time.RFC3339),
		msg,
	)
}