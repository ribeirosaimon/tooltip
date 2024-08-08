package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/ribeirosaimon/aergia-utils/logs"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	oncePgsql       sync.Once
	aergiaConn      AergiaPgsqlConnection
	pgsqlDefaultUrl = "jdbc:postgresql://localhost:5432/postgres"
)

type AergiaMongoInterface interface {
	GetConnection() *mongo.Database
}

// Option was a function optional pattern
type Option func(*AergiaPgsqlConnection)

func WithUrl(url string) Option {
	return func(a *AergiaPgsqlConnection) {
		a.url = url
	}
}

func WithDatabase(db string) Option {
	return func(a *AergiaPgsqlConnection) {
		a.database = db
	}
}

type AergiaPgsqlConnection struct {
	database string
	url      string
	pgsql    *sql.DB
}

type AergiaPgsqlConnectionInterface interface {
	GetConnection() *sql.DB
}

func NewConnPgsql(opts ...Option) AergiaPgsqlConnectionInterface {
	aergiaConn = AergiaPgsqlConnection{
		url: pgsqlDefaultUrl,
	}
	for _, opt := range opts {
		opt(&aergiaConn)
	}

	if aergiaConn.database == "" {
		panic("Need to set DATABASE")
	}
	oncePgsql.Do(func() {
		aergiaConn.pgsql = aergiaConn.conn()
	})
	return &aergiaConn
}

func (c *AergiaPgsqlConnection) conn() *sql.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := sql.Open("postgres", c.url)
	if err != nil {
		logs.ERROR.Message(fmt.Sprintf("Error opening database: %q", err))
	}
	err = db.PingContext(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			logs.ERROR.Message("Connection attempt timed out")
		} else {
			logs.ERROR.Message(fmt.Sprintf("Error connecting to the database: %v", err))
		}
	}

	return db
}

func (c *AergiaPgsqlConnection) GetConnection() *sql.DB {
	return c.pgsql
}
