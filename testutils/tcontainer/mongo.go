package tcontainer

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Mongo(ctx context.Context) (string, error) {

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
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

	port, err := mongoC.MappedPort(ctx, "27017")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not get container port: %s", err))
	}
	return fmt.Sprintf("mongodb://%s:%s", host, port.Port()), nil
}
