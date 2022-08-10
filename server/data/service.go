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
	start := time.Now()
	val := m.Cache[key]
	if val == nil {
		val, err = m.RedStore.Get(ctx, key).Bytes()
		if err == redis.Nil {
			return nil, nil
		}
		err := m.timeLog(key, time.Since(start).Microseconds(), "redis")
		if err != nil {
			return nil, err
		}
		return val, nil
	}
	err = m.timeLog(key, time.Since(start).Microseconds(), "cache")
	if err != nil {
		return nil, err
	}
	return m.Cache[key], nil
}

func (m *Service) clearCache(key string) {
	<-time.After(time.Second * 15)
	delete(m.Cache, key)
}

func (m *Service) timeLog(r string, t int64, s string) error {

	stmt := `insert into responsetimelog (request, timeof, showfrom) values ($1, $2, $3)`
	r = "film/" + r
	_, err := m.DB.Exec(stmt, r, t, s)
	if err != nil {
		return err
	}
	return nil
}
