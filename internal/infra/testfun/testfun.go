package testfun

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

func CreateTestDB(ctx context.Context) (string, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return "", fmt.Errorf("failed to initialize a pool: %w", err)
	}

	pg, err := pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "16",
			Name:       "integration-tests",
			Env: []string{
				"POSTGRES_USER=test",
				"POSTGRES_PASSWORD=test",
				"POSTGRES_DB=test",
			},
			ExposedPorts: []string{"5432"},
		},
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to run the postgres container: %w", err)
	}

	go func() {
		<-ctx.Done()
		if err := pool.Purge(pg); err != nil {
			log.Printf("failed to purge the postgres container: %v", err)
		}
	}()
	hostPort := pg.GetHostPort("5432/tcp")

	// wait db
	pool.MaxWait = 10 * time.Second
	host, port, err := getHostPort(hostPort)
	if err != nil {
		return "", fmt.Errorf("failed to extract the host and port parts from the string %s: %w", hostPort, err)
	}
	var getSUConnection = func() (*pgx.Conn, error) {
		conn, err := pgx.Connect(pgx.ConnConfig{
			Host:     host,
			Port:     port,
			Database: "test",
			User:     "test",
			Password: "test",
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get a user connection: %w", err)
		}
		return conn, nil
	}
	var conn *pgx.Conn
	if err := pool.Retry(func() error {
		conn, err = getSUConnection()
		if err != nil {
			return fmt.Errorf("failed to connect to the DB: %w", err)
		}
		return nil
	}); err != nil {
		return "", err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to correctly close the connection: %v", err)
		}
	}()

	return fmt.Sprintf("postgres://test:test@%s:%d/test?sslmode=disable", host, port), nil
}

func getHostPort(hostPort string) (string, uint16, error) {
	hostPortParts := strings.Split(hostPort, ":")
	if len(hostPortParts) != 2 {
		return "", 0, fmt.Errorf("got an invalid host-port string: %s", hostPort)
	}

	portStr := hostPortParts[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to cast the port %s to an int: %w", portStr, err)
	}
	return hostPortParts[0], uint16(port), nil
}
