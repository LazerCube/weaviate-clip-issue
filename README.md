# Weaviate CLIP Search Example

This example repo to try to improve CLIP text search in Weaviate.

Currently searching for "airplane" does not return any relevant results, even though the embedding text contains the word "airplane".

```shell
$ go run ./cmd/search "airplane"

Found 3 results
Filename: dog.jpg
Description: 
Certainty: 0.868234
Distance: 0.263533
ID: 35ff35af-1415-4414-a963-8cf65555c881
Image: [redacted]

Filename: city.jpg
Description: 
Certainty: 0.864373
Distance: 0.271254
ID: 9f3b602f-1303-4eba-bfc5-9fd355a72c0a
Image: [redacted]

Filename: sprite-fright.png
Description: A poster for sprite fright
Certainty: 0.817505
Distance: 0.364990
ID: ec88328e-54b0-4e0a-a6b8-b1bdba530e19
Image: [redacted]
```

You'd expect to see `airplane.jpg` in the results, but it's not there. Searching for "sprite fright" on the other hand does return the result but it's not the top result.

```shell
$ go run ./cmd/search "sprite fright"

Found 3 results
Filename: dog.jpg
Description: 
Certainty: 0.876298
Distance: 0.247405
ID: 35ff35af-1415-4414-a963-8cf65555c881
Image: [redacted]

Filename: sprite-fright.png
Description: A poster for sprite fright
Certainty: 0.870450
Distance: 0.259100
ID: ec88328e-54b0-4e0a-a6b8-b1bdba530e19
Image: [redacted]

Filename: city.jpg
Description: 
Certainty: 0.868882
Distance: 0.262237
ID: 9f3b602f-1303-4eba-bfc5-9fd355a72c0a
Image: [redacted]
```

It seems like adding more text to the embedding makes the results worse.

As an example `city.jpg` and `dog.jpg` don't have any extra descriptions added and searching for "animals" or "city" gives what you would expect.

```shell
$ go run ./cmd/search "city"

Found 3 results
Filename: city.jpg
Description: 
Certainty: 0.883115
Distance: 0.233770
ID: 9f3b602f-1303-4eba-bfc5-9fd355a72c0a
Image: [redacted]

Filename: dog.jpg
Description: 
Certainty: 0.867975
Distance: 0.264051
ID: 35ff35af-1415-4414-a963-8cf65555c881
Image: [redacted]

Filename: sprite-fright.png
Description: A poster for sprite fright
Certainty: 0.807907
Distance: 0.384187
ID: ec88328e-54b0-4e0a-a6b8-b1bdba530e19
Image: [redacted]
```

## How to run

```bash
docker-compose up -d
go run ./cmd/import
go run ./cmd/search <search term>
```