# Weaviate CLIP Search Example

This example repo to try to improve CLIP text search in Weaviate.

Currently searching for "airplane" or "sprite fright" does not return any relevant results, even though the embedding text contains these words. It seems like adding more text to the embedding make results worse.

```bash
docker-compose up -d
go run ./cmd/import
go run ./cmd/search <search term>
```