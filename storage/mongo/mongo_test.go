package mongo

import (
	"context"
	"log"
	"testing"

	"github.com/ribeirosaimon/aergia-utils/testutils/aergiatestcontainer"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

func (t *Test) GetId() primitive.ObjectID {
	return t.Id
}
func (t *Test) SetId(p primitive.ObjectID) {
	t.Id = p
}

func newTest() *Test {
	return &Test{
		Id: primitive.NewObjectID(),
	}
}

func TestMongo(t *testing.T) {
	ctx := context.Background()
	url, err := aergiatestcontainer.Mongo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	mongo := NewConnMongo(ctx, WithUrl(url), WithDatabase("test"))

	t.Run("need to create a insertId", func(t *testing.T) {
		test := newTest()
		one, insertError := mongo.GetConnection().Collection("structTest").
			InsertOne(ctx, test)
		assert.NoError(t, insertError)
		assert.NotNil(t, one)
		assert.Equal(t, one.InsertedID, test.GetId())
	})

	t.Run("need to find a insertId", func(t *testing.T) {
		// before
		test := newTest()
		one, insertError := mongo.GetConnection().Collection("structTest").
			InsertOne(ctx, test)

		// when
		log.Printf(test.GetId().String())
		var testInDatabase Test
		findError := mongo.GetConnection().
			Collection("structTest").
			FindOne(ctx, bson.M{"_id": test.GetId()}).
			Decode(&testInDatabase)

		assert.NoError(t, insertError)
		assert.NoError(t, findError)
		assert.NotNil(t, one)
		assert.NotNil(t, testInDatabase)
		assert.Equal(t, one.InsertedID, test.GetId())
		assert.Equal(t, testInDatabase.GetId(), test.GetId())
	})
}
