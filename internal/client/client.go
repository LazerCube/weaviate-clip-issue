package client

import "github.com/weaviate/weaviate-go-client/v4/weaviate"

func CreatClient() *weaviate.Client {
	cfg := weaviate.Config{
		Host:   "localhost:8080",
		Scheme: "http",
	}

	return weaviate.New(cfg)
}
