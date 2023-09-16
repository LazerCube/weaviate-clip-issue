package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lazercube/weaviate-clip-issue/internal/client"
	"github.com/lazercube/weaviate-clip-issue/internal/data"
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
		fmt.Printf("ID: %s\n", r.ID)
		fmt.Printf("Filename: %s\n", r.Filename)
		fmt.Printf("Text: %s\n", r.Text)
		fmt.Println("Image: [redacted]")
		fmt.Printf("Video: %s\n", "[redacted]")
		fmt.Printf("Audio: %s\n", "[redacted]")
		fmt.Printf("Certainty: %f\n", r.Certainty)
		fmt.Printf("Distance: %f\n", r.Distance)
		fmt.Println()
	}
}
