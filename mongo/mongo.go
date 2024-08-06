package mongo

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	onceMongo       sync.Once
	aergiaConn      AergiaConnection
	mongoDefaultUrl = "mongodb://localhost:27017"
)

type Entity interface {
	GetId() primitive.ObjectID
	SetId(id primitive.ObjectID)
}

type AergiaMongoInterface interface {
	GetConnection() *mongo.Database
}

// Option was a function optional pattern
type Option func(*AergiaConnection)

func WithUrl(url string) Option {
	return func(a *AergiaConnection) {
		a.url = url
	}
}

func WithDatabase(db string) Option {
	return func(a *AergiaConnection) {
		a.database = db
	}
}

func NewConnMongo(ctx context.Context, opts ...Option) AergiaMongoInterface {
	aergiaConn = AergiaConnection{
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

type AergiaConnection struct {
	url      string
	database string
	mongo    *mongo.Database
}

func (c *AergiaConnection) GetConnection() *mongo.Database {
	return c.mongo
}

func (c *AergiaConnection) conn(ctx context.Context) *mongo.Database {
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
