package main

import (
    "context"
    "database/sql"
    "log"
    "time"

    // Регистрируем драйвер pgx для database/sql
    _ "github.com/jackc/pgx/v5/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        return nil, err
    }

    // Настраиваем пул соединений
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(30 * time.Minute)

    // Проверяем подключение с ограничением по времени
    ctx, cancel := context.WithTimeout(
        context.Background(),
        2*time.Second,
    )
    defer cancel()

    if err := db.PingContext(ctx); err != nil {
        return nil, err
    }

    log.Println("Connected to PostgreSQL")
    return db, nil
}