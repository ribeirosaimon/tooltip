package tcontainer

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type RedisTestContainer struct {
	ctx  context.Context
	host string
}

func NewRedisTestContainer() *RedisTestContainer {
	return &RedisTestContainer{
		ctx: context.Background(),
	}
}
func (r *RedisTestContainer) Redis() error {

	req := testcontainers.ContainerRequest{
		Image:        "redis:alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForListeningPort("6379/tcp"),
	}

	redisContainer, err := testcontainers.GenericContainer(r.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return errors.New(fmt.Sprintf("Could not start container: %s", err))
	}
	host, err := redisContainer.Host(r.ctx)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not get container host: %s", err))
	}

	port, err := redisContainer.MappedPort(r.ctx, "6379")
	if err != nil {
		return errors.New(fmt.Sprintf("Could not get container port: %s", err))
	}
	r.host = fmt.Sprintf("%s:%s", host, port.Port())
	return nil
}

func (r *RedisTestContainer) GetHost() string {
	return r.host
}
