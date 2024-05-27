package payload_test

import "testing"

func TestSomething(t *testing.T) {
	// db.Query()
}

//var db *sql.DB
//
//func TestMain(m *testing.M) {
//	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
//	pool, err := dockertest.NewPool("")
//	if err != nil {
//		log.Fatalf("Could not construct pool: %s", err)
//	}
//	//pool.BuildAndRunWithBuildOptions("mysql", []string{"MYSQL_ROOT_PASSWORD=secret"}, &dockertest.BuildOptions{
//	//	Dockerfile: "../Dockerfile",
//	//})
//	//
//
//	// uses pool to try to connect to Docker
//	err = pool.Client.Ping()
//	if err != nil {
//		log.Fatalf("Could not connect to Docker: %s", err)
//	}
//
//	// pulls an image, creates a container based on it and runs it
//	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
//	if err != nil {
//		log.Fatalf("Could not start resource: %s", err)
//	}
//
//	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
//	if err := pool.Retry(func() error {
//		var err error
//		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
//		if err != nil {
//			return err
//		}
//		return db.Ping()
//	}); err != nil {
//		log.Fatalf("Could not connect to database: %s", err)
//	}
//
//	code := m.Run()
//
//	// You can't defer this because os.Exit doesn't care for defer
//	if err := pool.Purge(resource); err != nil {
//		log.Fatalf("Could not purge resource: %s", err)
//	}
//
//	os.Exit(code)
//}
