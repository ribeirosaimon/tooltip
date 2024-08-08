package mongo

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	onceMongo       sync.Once
	aergiaConn      AergiaMongoConnection
	mongoDefaultUrl = "mongodb://localhost:27017"
)

type AergiaMongoInterface interface {
	GetConnection() *mongo.Database
}

// Option was a function optional pattern
type Option func(*AergiaMongoConnection)

func WithUrl(url string) Option {
	return func(a *AergiaMongoConnection) {
		a.url = url
	}
}

func WithDatabase(db string) Option {
	return func(a *AergiaMongoConnection) {
		a.database = db
	}
}

func NewConnMongo(ctx context.Context, opts ...Option) AergiaMongoInterface {
	aergiaConn = AergiaMongoConnection{
		url: mongoDefaultUrl,
	}
	aergiaConn.conn(ctx)

	for _, opt := range opts {
		opt(&aergiaConn)
	}
	if aergiaConn.database == "" {
		panic("Need to set DATABASE")
	}
	onceMongo.Do(func() {
		aergiaConn.mongo = aergiaConn.conn(ctx)
	})
	return &aergiaConn
}

type AergiaMongoConnection struct {
	url      string
	database string
	mongo    *mongo.Database
}

func (c *AergiaMongoConnection) GetConnection() *mongo.Database {
	return c.mongo
}

func (c *AergiaMongoConnection) conn(ctx context.Context) *mongo.Database {
	clientOptions := options.Client().ApplyURI(c.url)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(c.database)
}
