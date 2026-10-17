package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "example.com/pz7-redis/internal/cache"
)

func main() {
    c := cache.New("localhost:6380")

    mux := http.NewServeMux()

    // Сохранение значения: /set?key=...&value=...
    mux.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")
        value := r.URL.Query().Get("value")

        if key == "" || value == "" {
            http.Error(w, "key and value required", http.StatusBadRequest)
            return
        }

        // Ключ будет храниться 10 секунд
        err := c.Set(key, value, 10*time.Second)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "OK: %s=%s (TTL 10s)", key, value)
    })

    // Получение значения: /get?key=...
    mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")

        val, err := c.Get(key)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }

        fmt.Fprintf(w, "VALUE: %s", val)
    })

    // Проверка оставшегося времени: /ttl?key=...
    mux.HandleFunc("/ttl", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")

        ttl, err := c.TTL(key)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "TTL: %s", ttl)
    })

    // Удаление значения: /del?key=...
    mux.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Query().Get("key")

        if err := c.Del(key); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "DELETED: %s", key)
    })

    log.Println("Server started on :8081")
    log.Fatal(http.ListenAndServe(":8081", mux))
}