package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lazercube/weaviate-clip-search/internal/client"
	"github.com/lazercube/weaviate-clip-search/internal/data"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <search term>\n", os.Args[0])
		os.Exit(1)
	}

	ctx := context.Background()
	client := client.CreatClient()

	result, err := data.SearchObjects(ctx, client, os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d results\n", len(result))
	for _, r := range result {
		fmt.Printf("Filename: %s\n", r.Filename)
		fmt.Printf("Certainty: %f\n", r.Certainty)
		fmt.Printf("Distance: %f\n", r.Distance)
		fmt.Printf("ID: %s\n", r.ID)
		fmt.Printf("Image: %s\n", "[redacted]")
		fmt.Println()
	}
}
