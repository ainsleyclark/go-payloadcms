package payload

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// See: https://github.com/mahjadan/go-integration-test/blob/main/integration/main_test.go
// For example

var db *sql.DB
var postgresR *dockertest.Resource

func TestMain(m *testing.M) {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	network, err := pool.CreateNetwork("backend")
	if err != nil {
		log.Fatalf("Could not create Network to docker: %s \n", err)
	}

	postgresCnt, uri, err := startPostgresDB(pool, network)
	if err != nil {
		cleanUp(1, pool, network, postgresCnt)
	}

	payload, err := startPayloadCMS(pool, network, uri)
	if err != nil {
		fmt.Printf("Could not start Payload CMS: %v \n", err)
		cleanUp(1, pool, network, postgresCnt, payload)
	}

	println("Starting tests")
	code := m.Run()
	println("Stopping tests")

	cleanUp(code, pool, network, postgresCnt, payload)
}

func startPostgresDB(pool *dockertest.Pool, network *dockertest.Network) (*dockertest.Resource, string, error) {
	fmt.Println("Starting Postgres")

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Name:       "postgres",
		Tag:        "13",
		Networks:   []*dockertest.Network{network},
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user",
			"POSTGRES_DB=payload",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		fmt.Printf("Could not start Mongodb: %v \n", err)
		return resource, "", err
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user:secret@%s/payload?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	if err := pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		fmt.Printf("Could not connect to mongodb container: %v \n", err)
		return resource, "", err
	}

	postgresR = resource

	return resource, databaseUrl, nil
}

// CONTAINER ID   IMAGE         COMMAND                  CREATED          STATUS          PORTS                     NAMES
//1a1e69ed5a2a   postgres:13   "docker-entrypoint.s…"   42 seconds ago   Up 41 seconds   0.0.0.0:32830->5432/tcp   postgres

func startPayloadCMS(pool *dockertest.Pool, network *dockertest.Network, dbURL string) (*dockertest.Resource, error) {
	fmt.Println("Starting Payload CMS")

	resource, err := pool.BuildAndRunWithOptions("../dev/Dockerfile", &dockertest.RunOptions{
		Hostname: "payload",
		Name:     "payload",
		Networks: []*dockertest.Network{network},
		Env: []string{
			"DATABASE_URI=" + fmt.Sprintf("postgres://user:secret@%s:5432/payload?sslmode=disable",
				postgresR.GetIPInNetwork(network),
			),
			"PAYLOAD_SECRET=secret",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		fmt.Printf("Could not start Payload CMS: %v \n", err)
		return resource, err
	}

	fmt.Printf("Moving to retry: " + "http://localhost:" + resource.GetPort("3000/tcp") + "/admin")

	if err := pool.Retry(func() error {
		url := resource.GetHostPort("3000/tcp") + "/admin"
		fmt.Println(url)
		post, err := http.DefaultClient.Get(url)
		if err != nil {
			return err
		}
		buf, err := io.ReadAll(post.Body)
		fmt.Println(string(buf))

		if post.StatusCode != 200 {
			return errors.New(fmt.Sprintf("Expected status 200, got %d", post.StatusCode))
		}
		return nil
	}); err != nil {
		return resource, err
	}

	fmt.Printf("Payload CMS is running on: %s \n", resource.GetPort("3000/tcp"))

	return resource, nil
}

//2024-05-27 21:30:42 [20:30:42] ERROR (payload): error: relation "users" does not exist

func cleanUp(code int, pool *dockertest.Pool, network *dockertest.Network, resources ...*dockertest.Resource) {
	cleanUpResources(pool, resources)
	if err := network.Close(); err != nil {
		log.Fatalf("Could not close network: %s\n", err)
	}
	os.Exit(code)
}

func cleanUpResources(pool *dockertest.Pool, resources []*dockertest.Resource) {
	fmt.Println("removing resources.")
	for _, resource := range resources {
		if resource != nil {
			if err := pool.Purge(resource); err != nil {
				log.Fatalf("Could not purge resource: %s\n", err)
			}
		}
	}
}
