package data

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

func CreateClass(ctx context.Context, client *weaviate.Client) error {
	multiModal := &models.Class{
		Class:       Class,
		Description: "Sample class holding information about a file",
		ModuleConfig: map[string]interface{}{
			"multi2vec-bind": map[string]interface{}{
				"imageFields": []string{"image"},
				"textFields":  []string{"text"},
				"audioFields": []string{"audio"},
				"videoFields": []string{"video"},
				"weights": map[string]interface{}{
					"textFields":  0.4,
					"imageFields": 0.2,
					"audioFields": 0.2,
					"videoFields": 0.2,
				},
			},
		},
		VectorIndexType: "hnsw",
		Vectorizer:      "multi2vec-bind",
		Properties: []*models.Property{
			{
				DataType:    []string{"string"},
				Description: "The name of the file",
				Name:        "filename",
			},
			{
				DataType:    []string{"string"},
				Description: "The description of the image",
				Name:        "text",
			},
			{
				DataType:    []string{"blob"},
				Description: "Base64 encoded image",
				Name:        "image",
			},
			{
				DataType:    []string{"blob"},
				Description: "audio",
				Name:        "audio",
			},
			{
				DataType:    []string{"blob"},
				Description: "video",
				Name:        "video",
			},
		},
	}
	return client.Schema().ClassCreator().WithClass(multiModal).Do(ctx)
}

func DeleteClass(ctx context.Context, client *weaviate.Client) error {
	exists, err := client.Schema().ClassExistenceChecker().WithClassName(Class).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return nil
	}

	return client.Schema().ClassDeleter().WithClassName(Class).Do(ctx)
}
