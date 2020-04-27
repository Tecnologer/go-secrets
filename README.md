# Go Secrets

This tool can help you to manage the secrets key for a Golang project, keeping in secrets all sensitive data. I.e.: username and password for a database server.

## How to use it

`go get github.com/tecnologer/go-secrets`

Code:

```golang
package main

import (
    "fmt"

    "github.com/google/uuid"
	"github.com/tecnologer/go-secrets"
)

func main() {
    bucketID, err := uuid.Parse("906e7526-f379-42e3-a6e4-8299488d90b1")
	if err != nil {
		log.Fatalf("Invalid bucket id. Error: %v", err)
    }

	bucket, err := secrets.GetBucket(bucketID)

	if err != nil {
		panic(err)
    }

    username := bucket.Get("username")

	fmt.Printf("Secret username: %s", username)
}
```

To add new keys you can use [go-secrets-cli][2]:

`go-secrets-cli set -id "906e7526-f379-42e3-a6e4-8299488d90b1" -key pwd -val 123`

### Dependencies

- [Google UUID][1]

[1]: https://pkg.go.dev/github.com/google/uuid
[2]: https://github.com/Tecnologer/go-secrets-cli
