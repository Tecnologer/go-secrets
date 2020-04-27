# Go Secrets

This tool can help you to manage the secrets key for a Golang project, keeping in secrets all sensitive data. I.e.: username and password for a database server.

## How to use it

`go get -u github.com/tecnologer/go-secrets`

To initialize the bucket, use `go-secrets-cli init`. go-secrets-cli is [here][2]. If you don't want use CLI, you can get the bucket with `secrets.GetBucketByUUID(<uuid.UUID>)` or `secrets.GetBucketByID(<string uuid>)`

Code:

```golang
package main

import (
    "fmt"

    "github.com/google/uuid"
	"github.com/tecnologer/go-secrets"
)

func main() {
	bucket, err := secrets.GetBucket()

	if err != nil {
		panic(err)
    }

    username := bucket.Get("username")

	fmt.Printf("Secret username: %s", username)
}
```

To add new keys you can use [go-secrets-cli][2]:

`go-secrets-cli set -id "906e7526-f379-42e3-a6e4-8299488d90b1" -key username -val tecnologer`

### Dependencies

- [Google UUID][1]

[1]: https://pkg.go.dev/github.com/google/uuid
[2]: https://github.com/Tecnologer/go-secrets-cli
