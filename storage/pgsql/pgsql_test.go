package pgsql

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/aergia-utils/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/testutils/aergiatestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestPgsql(t *testing.T) {
	var err error
	ctx := context.Background()
	url, err := aergiatestcontainer.Pgsql(ctx)
	if err != nil {
		t.Fatal(err)
	}
	pgsql := NewConnPgsql(WithUrl(url))

	t.Run("need insert test", func(t *testing.T) {
		one, insertError := pgsql.GetConnection().Exec(`
	    CREATE TABLE IF NOT EXISTS test (
	        id SERIAL PRIMARY KEY,
	        name TEXT NOT NULL
	    );
	`)
		assert.NoError(t, err)

		v, err := pgsql.GetConnection().Exec("INSERT INTO test (name) VALUES ($1)", "test")
		assert.NoError(t, err)

		assert.NoError(t, insertError)
		assert.NotNil(t, one)
		assert.NotNil(t, v)

		var name string
		err = pgsql.GetConnection().QueryRow("SELECT name FROM test WHERE name = $1", "test").Scan(&name)
		assert.NoError(t, err)
		assert.Equal(t, "test", name)

	})

	t.Run("create string query", func(t *testing.T) {
		u := sql.User{
			Email:    "test@test.com",
			Username: "test",
			Password: "test",
		}
		query := pgsql.CreateQuery(u)

		assert.Equal(t, "INSERT INTO user (ID, Username, Password, Email, FirstName, LastName, LoginAtempt, Role, Audit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);", query)
	})

	for _, v := range []struct {
		host      string
		port      string
		database  string
		username  string
		password  string
		urlString string
	}{
		{
			host: "localhost", port: "1233", database: "testdb", username: "username", password: "password",
			urlString: "postgres://username:password@localhost:1233/testdb?sslmode=disable",
		},
	} {
		t.Run("extract url values", func(t *testing.T) {
			host, port, database, username, password, err := extractDBDetails(v.urlString)

			assert.NoError(t, err)
			assert.Equal(t, v.host, host)
			assert.Equal(t, v.port, port)
			assert.Equal(t, v.database, database)
			assert.Equal(t, v.username, username)
			assert.Equal(t, v.password, password)

		})
	}

}
