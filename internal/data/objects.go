package data

import (
	"context"
	"errors"
	"fmt"
	"os"

	b64 "encoding/base64"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

const Class = "MultiModal"

func ImportObjects(ctx context.Context, client *weaviate.Client) error {
	basePath := "./images"
	files, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}

	// Removing these descriptions seems to improve the search
	// results even when text is used as a search term.
	descriptions := map[string]string{
		"airplane.jpg":      "Airplane flying through a blue sky with clouds",
		"sprite-fright.png": "A poster for sprite fright",
	}

	objects := []*models.Object{}
	for _, f := range files {
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", basePath, f.Name()))
		if err != nil {
			return err
		}

		description, ok := descriptions[f.Name()]
		if !ok {
			description = ""
		}

		image := b64.StdEncoding.EncodeToString(data)
		object := &models.Object{
			Class: Class,
			Properties: map[string]interface{}{
				"filename": f.Name(),
				"text":     description,
				"image":    image,
			},
		}
		objects = append(objects, object)
	}

	batcher := client.Batch().ObjectsBatcher()
	for i := range objects {
		batcher.WithObjects(objects[i])
	}

	resp, err := batcher.Do(ctx)
	if err != nil {
		return err
	}

	if len(resp) != len(objects) {
		return errors.New("not all objects imported")
	}

	return nil
}
