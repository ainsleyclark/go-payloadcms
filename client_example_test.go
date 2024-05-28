package payloadcms_test

import (
	"context"
	"fmt"
	"log"

	"github.com/ainsleydev/go-payloadcms"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func PayloadCMS() {
	client, err := payloadcms.New(
		payloadcms.WithBaseURL("http://localhost:8080"),
		payloadcms.WithAPIKey("api-key"),
	)

	if err != nil {
		log.Fatalln(err)
	}

	var user User
	resp, err := client.Collections.FindByID(context.Background(), "users", 1, &user)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Recieved status: %d, with body: %s\n", resp.StatusCode, string(resp.Body))
}
