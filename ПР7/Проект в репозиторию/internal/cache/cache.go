package cache

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
)

type Cache struct {
    rdb *redis.Client
}

func New(addr string) *Cache {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: "", // Подключаемся без пароля
        DB:       0,
    })

    return &Cache{rdb: rdb}
}

func (c *Cache) Set(key string, value string, ttl time.Duration) error {
    return c.rdb.Set(context.Background(), key, value, ttl).Err()
}

func (c *Cache) Get(key string) (string, error) {
    return c.rdb.Get(context.Background(), key).Result()
}

func (c *Cache) TTL(key string) (time.Duration, error) {
    return c.rdb.TTL(context.Background(), key).Result()
}

func (c *Cache) Del(key string) error {
    return c.rdb.Del(context.Background(), key).Err()
}