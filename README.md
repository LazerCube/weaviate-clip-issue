# Weaviate CLIP Text Search: Adverse Impact of Additional Descriptions

## Overview

This repository underlines a specific behavior observed with the CLIP text search in Weaviate. There's a noticeable degradation in search results when detailed descriptions are added to the CLIP model's embeddings.

## Problem Details

- **Primary Concern**: Adding detailed descriptions to the CLIP model's embeddings seems to interfere with the accuracy of the search results.
- **Evidence**: The following code snippet adds detailed descriptions to certain images:
    ```go
    descriptions := map[string]string{
        "airplane.jpg":      "Airplane flying through a blue sky with clouds",
        "sprite-fright.png": "A poster for sprite fright",
    }
    ```

When these descriptions are present:
- A search for "airplane" fails to produce the relevant image at all.
- "sprite fright" returns the expected image but not as the top result.

**Before** removing these descriptions:

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

...
```

**After** commenting out these descriptions:

```shell
$ go run ./cmd/search "sprite fright"
Found 1 results
Filename: sprite-fright.png
Description: 
Certainty: 0.904607
Distance: 0.190787
ID: 7078501e-0db7-4822-a497-daa6efafe951
Image: [redacted]
```

## Reproduction Steps

To replicate this behavior and contrast the search outcomes:

1. **Setup Environment:**

    ```shell
    $ docker-compose up -d
    ```

2. **Import Sample Data with Descriptions:**

    ```shell
    $ go run ./cmd/import
    ```

3. **Run a Search with Descriptions:**

    Use a search term like "sprite fright".
    ```shell
    $ go run ./cmd/search "sprite fright"
    ```

4. **Modify Descriptions:**
    
    Comment out the descriptions in the code. Specifically, go to [./internal/data/objects.go on line 26](https://github.com/lazercube/weaviate-clip-issue/blob/master/internal/data/objects.go#L26).