package pgsql

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/aergia-utils/testutils/aergiatestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestMongo(t *testing.T) {
	ctx := context.Background()
	url, err := aergiatestcontainer.Pgsql(ctx)
	if err != nil {
		t.Fatal(err)
	}
	pgsql := NewConnPgsql(WithUrl(url), WithDatabase("test"))

	t.Run("need insert test", func(t *testing.T) {
		one, insertError := pgsql.GetConnection().Exec(`
        CREATE TABLE IF NOT EXISTS test (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL
        );
    `)
		assert.NoError(t, err)

		_, err = pgsql.GetConnection().Exec("INSERT INTO test (name) VALUES ($1)", "test")
		assert.NoError(t, err)

		assert.NoError(t, insertError)
		assert.NotNil(t, one)

		var name string
		err = pgsql.GetConnection().QueryRow("SELECT name FROM test WHERE name = $1", "test").Scan(&name)
		assert.NoError(t, err)
		assert.Equal(t, "test", name)

	})

}
