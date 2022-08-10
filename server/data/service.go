package data

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v9"
	"time"
)

var ctx = context.Background()

type Service struct {
	DB       *sql.DB
	Cache    map[string][]byte
	RedStore *redis.Client
}

func (m *Service) AddToCache(value []byte, key string) {
	m.Cache[key] = value
	m.RedStore.Set(ctx, key, value, 30*time.Second)
	go m.clearCache(key)
}

func (m *Service) ShowFromCache(key string) (value []byte, err error) {
	if m.Cache[key] == nil {
		return m.RedStore.Get(ctx, key).Bytes()
	}
	return m.Cache[key], nil
}

func (m *Service) clearCache(key string) {
	<-time.After(time.Second * 15)
	delete(m.Cache, key)
}
