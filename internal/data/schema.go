package data

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

func CreateClass(ctx context.Context, client *weaviate.Client) error {
	multiModal := &models.Class{
		Class:       Class,
		Description: "Sample class holding all the images",
		ModuleConfig: map[string]interface{}{
			"multi2vec-clip": map[string]interface{}{
				"imageFields": []string{"image"},
				"textFields":  []string{"description"},
				// "weights": map[string]interface{}{
				// 	"textFields":  0.7,
				// 	"imageFields": 0.3,
				// },
			},
		},
		VectorIndexType: "hnsw",
		Vectorizer:      "multi2vec-clip",
		Properties: []*models.Property{
			{
				DataType:    []string{"string"},
				Description: "The name of the file",
				Name:        "filename",
			},
			{
				DataType:    []string{"string"},
				Description: "The description of the image",
				Name:        "description",
			},
			{
				DataType:    []string{"blob"},
				Description: "Base64 encoded image",
				Name:        "image",
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
