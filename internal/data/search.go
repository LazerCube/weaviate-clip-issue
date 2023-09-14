package data

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
)

type (
	Result struct {
		Filename    string
		Image       string
		Description string
		Certainty   float64
		Distance    float64
		ID          string
	}
)

func SearchObjects(ctx context.Context, client *weaviate.Client, text string) ([]*Result, error) {
	nearText := client.GraphQL().
		NearTextArgBuilder().
		WithConcepts([]string{text}).
		WithAutocorrect(true)

	response, err := client.GraphQL().Get().
		WithClassName("MultiModal").
		WithFields([]graphql.Field{
			{Name: "filename"},
			{Name: "image"},
			{Name: "description"},
			{Name: "_additional", Fields: []graphql.Field{
				{Name: "certainty"},
				{Name: "distance"},
				{Name: "id"},
			}},
		}...).
		WithNearText(nearText).
		WithAutocut(1).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	data := response.Data["Get"].(map[string]interface{})["MultiModal"].([]interface{})
	results := []*Result{}

	for _, row := range data {
		rowMap := row.(map[string]interface{})
		additionalMap := rowMap["_additional"].(map[string]interface{})
		results = append(results, &Result{
			Filename:    rowMap["filename"].(string),
			Image:       rowMap["image"].(string),
			Description: rowMap["description"].(string),
			Certainty:   additionalMap["certainty"].(float64),
			Distance:    additionalMap["distance"].(float64),
			ID:          additionalMap["id"].(string),
		})
	}

	return results, nil
}
