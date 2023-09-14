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

	objects := []*models.Object{}
	for _, f := range files {
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", basePath, f.Name()))
		if err != nil {
			return err
		}
		image := b64.StdEncoding.EncodeToString([]byte(data))
		object := &models.Object{
			Class: Class,
			Properties: map[string]interface{}{
				"filename":    f.Name(),
				"description": "A picture of a " + f.Name(),
				"image":       image,
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
