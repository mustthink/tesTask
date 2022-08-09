package data

import (
	"database/sql"
	"time"
)

type Service struct {
	DB    *sql.DB
	Cache map[string][]byte
}

func (m *Service) AddToCache(value []byte, key string) {
	m.Cache[key] = value
	go m.clearCache(key)
}

func (m *Service) ShowFromCache(key string) (value []byte) {
	if m.Cache[key] != nil {
		print("use cache")
	}
	return m.Cache[key]
}

func (m *Service) clearCache(key string) {
	<-time.After(time.Second * 15)
	delete(m.Cache, key)
}
