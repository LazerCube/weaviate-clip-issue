package main

import (
	"context"
	"fmt"

	"github.com/lazercube/weaviate-clip-issue/internal/client"
	"github.com/lazercube/weaviate-clip-issue/internal/data"
)

func main() {
	ctx := context.Background()

	client := client.CreatClient()

	fmt.Printf("Delete class\n")
	if err := data.DeleteClass(ctx, client); err != nil {
		panic(err)
	}

	fmt.Printf("Create class\n")
	if err := data.CreateClass(ctx, client); err != nil {
		panic(err)
	}

	fmt.Printf("Import objects\n")
	if err := data.ImportObjects(ctx, client); err != nil {
		panic(err)
	}

	fmt.Printf("Successfully imported data\n")
}
