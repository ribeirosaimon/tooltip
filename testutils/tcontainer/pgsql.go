package tcontainer

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Pgsql(ctx context.Context) (string, error) {

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

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not start container: %s", err))
	}
	host, err := mongoC.Host(ctx)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not get container host: %s", err))
	}

	port, err := mongoC.MappedPort(ctx, "5432")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not get container port: %s", err))
	}
	return fmt.Sprintf("postgres://test:test@%s:%s/testdb?sslmode=disable", host, port.Port()), nil
}
