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
	mongoConn       MongoConnection
	mongoDefaultUrl = "mongodb://localhost:27017"
)

type MConnInterface interface {
	GetConnection() *mongo.Database
}

// Option was a function optional pattern
type Option func(*MongoConnection)

func WithUrl(url string) Option {
	return func(a *MongoConnection) {
		a.url = url
	}
}

func WithDatabase(db string) Option {
	return func(a *MongoConnection) {
		a.database = db
	}
}

func NewConnMongo(ctx context.Context, opts ...Option) MConnInterface {
	mongoConn = MongoConnection{
		url: mongoDefaultUrl,
	}
	mongoConn.conn(ctx)

	for _, opt := range opts {
		opt(&mongoConn)
	}
	if mongoConn.database == "" {
		panic("Need to set DATABASE")
	}
	onceMongo.Do(func() {
		mongoConn.mongo = mongoConn.conn(ctx)
	})
	return &mongoConn
}

type MongoConnection struct {
	url      string
	database string
	mongo    *mongo.Database
}

func (c *MongoConnection) GetConnection() *mongo.Database {
	return c.mongo
}

func (c *MongoConnection) conn(ctx context.Context) *mongo.Database {
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
