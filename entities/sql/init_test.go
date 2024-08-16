package sql

import (
	"fmt"
	"os"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	pass := os.Getenv("DEFAULT_SAIMON_PASSWORD")
	pgsqlUrl := fmt.Sprintf("postgres://localhost:5432/postgres?user=postgres&password=%s", pass)
	initTableDatabase(pgsqlUrl)
}
