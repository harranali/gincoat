package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/gincoat/gincoat/core/env"
	"github.com/go-redis/redis/v8"
)

//redisCTX context for redis
var redisCTX context.Context

// redisTTL expiry time for redis records
var redisTTL time.Duration

// CacheEngine handles the caching operations
type CacheEngine struct {
	redisDB *redis.Client
}

var cEngine *CacheEngine

// New initiates a new caching engine
func New() *CacheEngine {
	redisCTX = context.Background()
	ttl, _ := strconv.ParseUint(env.Get("REDIS_TTL_SECONDS"), 10, 64)
	redisTTL = time.Duration(ttl)

	cEngine = &CacheEngine{
		redisDB: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
	return cEngine
}

// Resolve resolves initiated caching engine
func Resolve() *CacheEngine {
	return cEngine
}

// Set set a key, val pair in the cache
func (c *CacheEngine) Set(key string, val string) (bool, error) {
	status := c.redisDB.Set(redisCTX, key, val, redisTTL)
	if status.Err() != nil {
		return false, status.Err()
	}

	return true, nil
}

// Get retrieves a val from cache by a given key
func (c *CacheEngine) Get(key string) (interface{}, error) {
	val, err := c.redisDB.Get(redisCTX, key).Result()
	if err != nil {
		return false, err
	}
	return val, nil
}
