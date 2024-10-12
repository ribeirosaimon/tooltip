package tcontainer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PgsqlTestContainer struct {
	ctx  context.Context
	host string
}

func NewPgsqlTestContainer() *PgsqlTestContainer {
	return &PgsqlTestContainer{
		ctx: context.Background(),
	}
}
func (p *PgsqlTestContainer) Pgsql() error {

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	pgsqlC, err := testcontainers.GenericContainer(p.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return errors.New(fmt.Sprintf("Could not start container: %s", err))
	}
	host, err := pgsqlC.Host(p.ctx)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not get container host: %s", err))
	}

	port, err := pgsqlC.MappedPort(p.ctx, "5432")
	if err != nil {
		return errors.New(fmt.Sprintf("Could not get container port: %s", err))
	}
	p.host = fmt.Sprintf("postgres://test:test@%s:%s/testdb?sslmode=disable", host, port.Port())
	return nil
}

func (p *PgsqlTestContainer) CreateScripts(sqlFile string) {
	fileContent, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", sqlFile, err)
	}

	db, err := sql.Open("postgres", p.host)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer db.Close()
	_, err = db.Exec(string(fileContent))
	if err != nil {
		log.Fatalf("failed to execute query %s: %v", sqlFile, err)
	}
}

func (p *PgsqlTestContainer) GetHost() string {
	return p.host
}
