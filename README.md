<p align="center">
  <img src="./res/payload-logo.jpg" height="86">
</p>

<p align="center">
    <a href="https://ainsley.dev">
        <h2 align="center">Go Payload CMS</h2>
		<p align="center">GoLang client library & SDK for Payload CMS</p>
    </a>
</p>

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/ainsleyclark/go-payloadcms)](https://goreportcard.com/report/github.com/ainsleyclark/go-payloadcms)
[![Maintainability](https://api.codeclimate.com/v1/badges/9cb93230fcfd6643dfa5/maintainability)](https://codeclimate.com/github/ainsleyclark/go-payloadcms/maintainability)
[![GoDoc](https://godoc.org/github.com/ainsleyclark/go-payloadcms?status.svg)](https://godoc.org/github.com/ainsleyclark/go-payloadcms)
[![Test](https://github.com/ainsleyclark/go-payloadcms/actions/workflows/test.yaml/badge.svg?query=new)](https://github.com/ainsleyclark/go-payloadcms/actions/workflows/test.yaml)
[![Lint](https://github.com/ainsleyclark/go-payloadcms/actions/workflows/lint.yaml/badge.svg?query=new)](https://github.com/ainsleyclark/go-payloadcms/actions/workflows/lint.yaml)
<br />
[![codecov](https://codecov.io/gh/ainsleyclark/go-payloadcms/graph/badge.svg?token=hhaaiJW6Vq)](https://codecov.io/gh/ainsleyclark/go-payloadcms)
[![ainsley.dev](https://img.shields.io/badge/-ainsley.dev-black?style=flat&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAABmJLR0QA/wD/AP+gvaeTAAAAB3RJTUUH5wEYDzUGL1b35AAABA1JREFUWMPtlttvFVUUxn977ZnZu+W0tLalqRovBAUvQag0xNQbpSIosSSIJC198YknJfHJxDf9A/DBJ0x8MbFACjVqvCASq6FYFLFBvJAaAomkFCmhHGpLO+PDzOmZzpn2nKP4pCs5ycmevb7vW99as/fA//FfD1XO5p1nzuA3NWJHx5T8cVkRBPHHQfRjd0tzyZhOOQIy27bAxET9zCuvvhY0r2kC/OiRABeAN4BL/4oDr9+3lGszPs7UVNfUE23v3Nj5koszR/8N4EXg3XJckFIFuCLUuU7GWNNtTg25cu4syJx0F+gGMuU4UJKAt1Yux1UKV6TVat1qs+OYwQESMwDQCjwKsOv4iZsnwGihwbiuEek2WjJGhMrvv0UujYKa08VFkQvuTXNgz6oVeCIo1CqrZYMRwTiaytERKn44kRQAsAFYDbBrsLgLRQU0GI919TXKiHQaUQ1GBCuCCQKqjg/MqInrM4lZrgc6A1CljHhRAZ4Ip65m77FaOmbJdehC5vzZr1RAf/T6x6NDwb3/uAVfP74GnwCjZasRuXuWXASj9XQme+3t6erqPcB0IvUuYCsUH8YFBRhRNBqvyYpsn0MeOnG6wvc/9x33MPBjSvp24Na/7cDP7Y/gKIURecZoeTBObkSwWg7UNjaOeFfGLgK9KRAPAM8Wc2FeAUaEWtddbEV2WBFtREXkCqvlghE5yOQkvucBHAR+T0BooAtYXLYDI5sewxWFJ/Kk1bI2UTlW5DMFp03+JPwJ+DQFai2wbiEXUgVUas0trmuslm4jUmGi/tuwDVmrpafBuNPVrs7N/wzQA2QTUJbwYLIlOxB0tOGJ4IhqsSJts+T54Rv0lBz1RFh9ZJA385fOAHAshaMNaAF4OcWFQgeUwhMlrlJdnqjaOLkR8Y2WvbWec9VIQeo4sJf8FZ2LmmgWJO1cmm8I7wc2a6XwosGL+v+rFfnYUYplh47Obo5dvZ8Av6TgbSZ8KxYWEGxZn/u7Dbg9t8HNnwF9S2qqzqVUn4vzQF/K+m3AC1A4jGlId0QC8l0BXKVGrahe//okNR99WZAUc6EXuJiC+zxw57wOxKp/DliRAvCFKDUkxS+YIeBwyvryCHuOC0kH6oBOCj/V/gTeA6aK0oefZj3ARGJdRdh1BQ7Eqm8HHk4B/Q7oB1B9acWFEWtDf5STjGbgqbgLcQcqCQ8NL5EUAPuBsRKqz8UVYB+F97QXcSyatSXoWJ8zvB04AFQlkoaBp4HhhaqPR1TdUsLjeVni8TjhVX0odCAkd4AdKeQAHxIwXEb1Odt+Az5IeVQVcTmhgDBWAhtTNl8G9qGAwKfU2N3SnJvi/RFGMjYCD8UFdACNKRsHgZMA6v0j5ZpAlPtNyvqSiJO/AKik60y0ALlUAAAAJXRFWHRkYXRlOmNyZWF0ZQAyMDIzLTAxLTI0VDE1OjUzOjA2KzAwOjAwm5vntAAAACV0RVh0ZGF0ZTptb2RpZnkAMjAyMy0wMS0yNFQxNTo1MzowNiswMDowMOrGXwgAAABXelRYdFJhdyBwcm9maWxlIHR5cGUgaXB0YwAAeJzj8gwIcVYoKMpPy8xJ5VIAAyMLLmMLEyMTS5MUAxMgRIA0w2QDI7NUIMvY1MjEzMQcxAfLgEigSi4A6hcRdPJCNZUAAAAASUVORK5CYII=)](https://ainsley.dev)
[![Twitter Handle](https://img.shields.io/twitter/follow/ainsleydev)](https://twitter.com/ainsleydev)

</div>

## Introduction

Go Payload is a GoLang client library for Payload CMS. It provides a simple and easy-to-use
interface for interacting with the Payload CMS API by saving you the hassle of dealing with
unmarshalling JSON responses into your Payload collections and globals.

## Installation

```bash
go get -u github.com/ainsleyclark/go-payloadcms
```

## Quick Start

```go
import (
	"github.com/ainsleyclark/go-payloadcms"
)

type Entity struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	// ... other fields
}

func main() {
	client, err := payloadcms.New(
		payloadcms.WithBaseURL("http://localhost:8080"),
		payloadcms.WithAPIKey("api-key"),
	)
	
	if err != nil {
		log.Fatalln(err)
	}
	
	var entities payloadcms.ListResponse[Entity]
	resp, err := client.Collections.List(context.Background(), "users", payloadcms.ListParams{}, &entities)
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Printf("Received status: %d, with body: %s\n", resp.StatusCode, string(resp.Content))
}
```

## Docs

Documentation can be found at
the [Go Docs](https://pkg.go.dev/github.com/ainsleyclark/go-payloadcms), but we have included a
kick-start guide below to get you started.

## Services

The client provides the services as defined below. Each

- `Collections`
- `Globals`
- `Media`

### Collections

The collections service provides methods to interact with the collections in Payload CMS.
For more information please visit the docs [here](https://payloadcms.com/docs/api/collections).

#### FindByID

```go
var entity Entity // Any struct that conforms to your collection schema.
resp, err := client.Collections.FindByID(context.Background(), "collection", 1, &entity)
if err != nil {
	fmt.Println(err)
	return
}
```

#### List

```go
var entities payloadcms.ListResponse[Entity] // Must use ListResponse with generic type.
resp, err := client.Collections.List(context.Background(), "collection", payloadcms.ListParams{
	Sort:  "-createdAt",
	Limit: 10,
	Page:  1,
}, entities)
if err != nil {
	fmt.Println(err)
	return
}
```

#### Create

```go
var entity Entity // Any struct representing the entity to be created.
resp, err := client.Collections.Create(context.Background(), "collection", entity)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(string(resp.Content)) // Can unmarshal into response struct if needed.
```

#### UpdateByID

```go
var entity Entity // Any struct representing the updated entity.
resp, err := client.Collections.UpdateByID(context.Background(), "collection", 1, entity)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(string(resp.Content)) // Can unmarshal into response struct if needed.
```

#### DeleteByID

```go
resp, err := client.Collections.DeleteByID(context.Background(), "collection", 1)
if err != nil {
	fmt.Println(err)	
	return
}
// Use response data as needed
```

### Globals

The globals service provides methods to interact with the globals in Payload CMS.
For more information please visit the docs [here](https://payloadcms.com/docs/api/globals).

#### Get

```go
var global Global // Any struct representing a global type.
resp, err := client.Globals.Get(context.Background(), "global", &global)
if err != nil {
	fmt.Println(err)
	return
}
```

#### Update

```go
var global Global // Any struct representing a global type.
resp, err := client.Globals.Update(context.Background(), "global", global)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(string(resp.Content)) // Can unmarshal into response struct if needed.
```

### Media

The media service provides methods to upload media types to Payload CMS.
For more information please visit the docs [here](https://payloadcms.com/docs/upload/overview).

#### Upload

```go
file, err := os.Open("path/to/file")
if err != nil {
	fmt.Println(err)
	return
}

media := &payloadcms.CreateResponse[Media]{}
_, err = m.payload.Media.UploadFromURL(ctx, file, Media{Alt: "alt"}, &media, payloadcms.MediaOptions{
	Collection:       "media",
})

if err != nil {
	fmt.Println(err)
	return
}
```

#### UploadFromURL

```go
media := &payloadcms.CreateResponse[Media]{}
_, err = m.payload.Media.UploadFromURL(ctx, "https://payloadcms.com/picture-of-cat.jpg", Media{Alt: "alt"}, &media, payloadcms.MediaOptions{
	Collection:       "media",
})

if err != nil {
	fmt.Println(err)
	return
}
```

#### Queries

The `Params` allows you to add filters, sort order, pagination, and other query parameters. Here's an example:

```go
import (
	"context"
	"fmt"
	"github.com/ainsleyclark/go-payloadcms"
)

func main() {
	client, err := payloadcms.New(
		payloadcms.WithBaseURL("http://localhost:8080"),
		payloadcms.WithAPIKey("api-key"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	var entities payloadcms.ListResponse[Entity]
	
	params := payloadcms.ListParams{
		Sort: "-createdAt",          
		Limit: 10,                  
		Page: 1,                     
		Where: payloadcms.Query().Equals("status", "active").GreaterThan("age", "18"),
	}

	resp, err := client.Collections.List(context.Background(), "users", params, &entities)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Received status: %d, with body: %s\n", resp.StatusCode, string(resp.Content))
}
```

## Mocks

Mock implementations can be found in `payloadfakes` package located
in [fakes](https://github.com/ainsleyclark/go-payloadcms/tree/main/fakes) directory.

They provide mock implementations of all the services provided by the client for convenience.

**Example:**

```go
func TestPayload(t *testing.T) {
	// Create a new mock collection service
	mockCollectionService := payloadfakes.NewMockCollectionService()
	
	// Define the behavior of the FindByID method
	mockCollectionService.FindByIDFunc = func (ctx context.Context,
		collection payloadcms.Collection,
		id int,
		out any,
	) (payloadcms.Response, error) {
		// Custom logic for the mock implementation
		return payloadcms.Response{}, nil
	}
	
	// Use the mock collection service in your tests
	myFunctionUsingCollectionService(mockCollectionService)
}
```

## Response and Error Types

The library defines custom response and error types to provide a more convenient way to interact with
the API.

* **Response:** This struct wraps the standard `http.Response` returned by the API and provides
  additional fields for easier access to response data and potential errors.
	* `Content`: The response body bytes.
	* `Message`: A user-friendly message extracted from the response (if available).
	* `Errors`: A list of `Error` structs containing details about any API errors encountered.
* **Error:** This struct represents a single API error with a `Message` field containing the error
  description.

These types are used throughout the library to handle successful responses and API errors
consistently.

## Development

### Setup

To get set up with Go Payload simply clone the repo and run the following:

```bash
make setup
```

This will install all dependencies and set up the project for development.

### Payload Dev Env

Within the `./dev` directory, you will find a local instance of Payload CMS that can be used for
testing the client.
To get setup with Payload, simply follow the steps below.

Copy the environment file and replace where necessary. The `postgres-db` adapater is currently being
used for the database.

```bash
cp .env.example .env
```

Then run the Payload instance like you would any other installation.

```bash
pnpm run dev
```

## TODOs

- Auth - https://payloadcms.com/docs/rest-api/overview#auth-operations
- Preferences:    https://payloadcms.com/docs/rest-api/overview#preferences
- Finish E2E Tests under `/tests` directory using `dockertest`

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvement, please open an
issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Trademark

ainsley.dev and the ainsley.dev logo are either registered trademarks or trademarks of ainsley.dev
LTD in the United Kingdom and/or other countries. All other trademarks are the property of their
respective owners.
