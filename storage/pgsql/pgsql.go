package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"

	_ "github.com/lib/pq"
	"github.com/ribeirosaimon/tooltip/tlog"
	"github.com/ribeirosaimon/tooltip/tserver"
)

var (
	oncePgsql       sync.Once
	pgConn          PostgresConnection
	pgsqlDefaultUrl = "jdbc:postgresql://localhost:5432/postgres"
)

// Option was a function optional pattern
type Option func(*PostgresConnection)

func WithUrl(url string) Option {
	return func(a *PostgresConnection) {
		a.url = url
	}
}

type PostgresConnection struct {
	url   string
	pgsql *sql.DB
}

type PConnInterface interface {
	GetConnection() *sql.DB
}

func NewConnPgsql(opts ...Option) *PostgresConnection {
	pgConn = PostgresConnection{
		url: pgsqlDefaultUrl,
	}
	for _, opt := range opts {
		opt(&pgConn)
	}

	oncePgsql.Do(func() {
		pgConn.pgsql = pgConn.conn()
		pgConn.CreateScripts(tserver.GetPgsqlConfig().EntryPoint)
	})
	return &pgConn
}

func (c *PostgresConnection) conn() *sql.DB {
	ctx := context.TODO()
	host, port, dbname, user, password, err := extractDBDetails(c.url)
	if err != nil {
		tlog.Error("NewConnPgsql", fmt.Sprintf("Error opening database: %q", err))
	}
	sprintf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", sprintf)
	if err != nil {
		tlog.Error("NewConnPgsql", "Error opening database: %q", err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			tlog.Error("NewConnPgsql", "Connection attempt timed out")
		} else {
			tlog.Error("NewConnPgsql", fmt.Sprintf("Error connecting to the database: %v", err))
		}
	}

	return db
}

func extractDBDetails(jdbcURL string) (string, string, string, string, string, error) {
	re := regexp.MustCompile(`^postgres://(.+):(.+)@([^:/?#]+):(\d+)/([^/?#]+)`)

	match := re.FindStringSubmatch(jdbcURL)

	if len(match) != 6 {
		tlog.Error("extractDBDetails", "Error parsing connection string")
		return "", "", "", "", "", errors.New("error parsing connection string")
	}

	user := match[1]
	password := match[2]
	host := match[3]
	port := match[4]
	dbname := match[5]

	return host, port, dbname, user, password, nil
}

func (c *PostgresConnection) CreateScripts(sqlFile string) {
	fileContent, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", sqlFile, err)
	}

	db, err := sql.Open("postgres", c.url)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer db.Close()
	_, err = db.Exec(string(fileContent))
	if err != nil {
		log.Fatalf("failed to execute query %s: %v", sqlFile, err)
	}
}

func (c *PostgresConnection) GetConnection() *sql.DB {
	return c.pgsql
}
