package cache

import (
	"time"

	"github.com/muesli/cache2go"
)

type AppCache struct {
	client *cache2go.CacheTable
}

var MyCache *AppCache

func InitCache() {
	MyCache = &AppCache{
		client: cache2go.Cache("cache"),
	}
}

func (r *AppCache) GetCacheItem(key string) *cache2go.CacheItem {

	res, err := r.client.Value(key)
	if err != nil {
		return nil
	}
	return res
}

func (r *AppCache) AddCacheItem(key string, value interface{}) {

	r.client.Add(key, 24*time.Hour, value)
}
