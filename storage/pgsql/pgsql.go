package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"

	_ "github.com/lib/pq"
	"github.com/ribeirosaimon/aergia-utils/logs"
)

var (
	oncePgsql       sync.Once
	aergiaConn      AergiaPgsqlConnection
	pgsqlDefaultUrl = "jdbc:postgresql://localhost:5432/postgres"
)

// Option was a function optional pattern
type Option func(*AergiaPgsqlConnection)

func WithUrl(url string) Option {
	return func(a *AergiaPgsqlConnection) {
		a.url = url
	}
}

type AergiaPgsqlConnection struct {
	url   string
	pgsql *sql.DB
}

type AergiaPgsqlInterface interface {
	GetConnection() *sql.DB
	CreateQuery(v any) string
}

func NewConnPgsql(opts ...Option) AergiaPgsqlInterface {
	aergiaConn = AergiaPgsqlConnection{
		url: pgsqlDefaultUrl,
	}
	for _, opt := range opts {
		opt(&aergiaConn)
	}

	oncePgsql.Do(func() {
		aergiaConn.pgsql = aergiaConn.conn()
	})
	return &aergiaConn
}

func (c *AergiaPgsqlConnection) conn() *sql.DB {
	ctx := context.TODO()
	host, port, dbname, user, password, err := extractDBDetails(c.url)
	if err != nil {
		logs.ERROR.Message(fmt.Sprintf("Error opening database: %q", err))
	}
	sprintf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", sprintf)
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

func extractDBDetails(jdbcURL string) (string, string, string, string, string, error) {
	re := regexp.MustCompile(`^postgres://(.+):(.+)@([^:/?#]+):(\d+)/([^/?#]+)\?`)

	match := re.FindStringSubmatch(jdbcURL)

	if len(match) != 6 {
		logs.ERROR.Message("Error parsing connection string")
		return "", "", "", "", "", errors.New("error parsing connection string")
	}

	user := match[1]
	password := match[2]
	host := match[3]
	port := match[4]
	dbname := match[5]

	return host, port, dbname, user, password, nil
}

func (c *AergiaPgsqlConnection) GetConnection() *sql.DB {
	return c.pgsql
}

func (c *AergiaPgsqlConnection) CreateQuery(v any) string {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		logs.ERROR.Message("CreateQuery: expected a struct")
	}

	tableName := typ.Name()
	var columns []string
	var placeholders []string

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		columns = append(columns, field.Name)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		strings.ToLower(tableName),
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query
}
