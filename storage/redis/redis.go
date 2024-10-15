package redis

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	oncePgsql       sync.Once
	redisConn       Connection
	pgsqlDefaultUrl = "localhost:6379"
)

type Option func(*Connection)

func WithUrl(url string) Option {
	return func(a *Connection) {
		a.url = url
	}
}

type Connection struct {
	url      string
	password string
	database int
	redis    *redis.Client
}

func NewRedisConnection(opts ...Option) *Connection {
	redisConn = Connection{
		url: pgsqlDefaultUrl,
	}
	for _, opt := range opts {
		opt(&redisConn)
	}

	oncePgsql.Do(func() {
		redisConn.redis = redisConn.conn()
	})

	return &redisConn
}

func (c *Connection) conn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.url,
		Password: c.password,
		DB:       c.database,
	})
}
