package payload

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	mongoCnt, uri, err := startMongoDB(pool, network)
	if err != nil {
		cleanUp(1, pool, network, mongoCnt)
	}

	payload, err := startPayloadCMS(pool, network, uri)
	if err != nil {
		cleanUp(1, pool, network, mongoCnt)
	}

	println("Starting tests")
	code := m.Run()
	println("Stopping tests")

	cleanUp(code, pool, network, mongoCnt, payload)
}

func startMongoDB(pool *dockertest.Pool, network *dockertest.Network) (*dockertest.Resource, string, error) {
	r, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       "mongodb",
		Repository: "mongo",
		Tag:        "3.6",
		Networks:   []*dockertest.Network{network},
	})
	if err != nil {
		fmt.Printf("Could not start Mongodb: %v \n", err)
		return r, "", err
	}
	mongoPort := r.GetPort("27017/tcp")
	uri := fmt.Sprintf("mongodb://localhost:%s", mongoPort)

	fmt.Printf("mongo-%s - connecting to : %s \n", "3.6", fmt.Sprintf("mongodb://localhost:%s", mongoPort))
	if err := pool.Retry(func() error {
		var err error

		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return err
		}

		err = client.Ping(context.TODO(), nil)
		if err == nil {
			fmt.Println("successfully connected to Mongodb.")
		}
		return err

	}); err != nil {
		fmt.Printf("Could not connect to mongodb container: %v \n", err)
		return r, "", err
	}

	return r, uri, nil
}

func startPayloadCMS(pool *dockertest.Pool, network *dockertest.Network, mongoURI string) (*dockertest.Resource, error) {
	r, err := pool.BuildAndRunWithOptions("../dev/Dockerfile", &dockertest.RunOptions{
		Name:       "payload",
		Repository: "payloadcms",
		Networks:   []*dockertest.Network{network},
		Env:        []string{"DATABASE_URI=" + mongoURI},
	})
	if err != nil {
		fmt.Printf("Could not start Payload CMS: %v \n", err)
		return r, err
	}

	fmt.Printf("Payload CMS is running on: %s \n", r.GetPort("3000/tcp"))

	return r, nil
}

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
