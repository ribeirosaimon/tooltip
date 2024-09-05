package sql

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/aergia-utils/logs"
	"github.com/ribeirosaimon/aergia-utils/testutils/aergiatestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestInitDatabase(t *testing.T) {
	pgsqlUrl, err := aergiatestcontainer.Pgsql(context.Background())
	logs.LOG.Message(pgsqlUrl)
	assert.NoError(t, err)

	err = CreateTableDatabase(pgsqlUrl)
	assert.Nil(t, err)
}

func TestInitMockDatabase(t *testing.T) {
	pgsqlUrl, err := aergiatestcontainer.Pgsql(context.Background())
	logs.LOG.Message(pgsqlUrl)
	assert.NoError(t, err)

	err = MockCreateTableDatabase(pgsqlUrl, map[string]bool{
		"user.sql": false,
	})
	assert.Nil(t, err)
}
