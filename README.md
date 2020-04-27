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

> For security use [go-secrets-cli][2] to add new keys:

`go-secrets-cli set -key username -val tecnologer`

- Create a group of keys and use it

```golang
package main

import (
    "fmt"

	"github.com/tecnologer/go-secrets"
)

func main() {
	bucket, err := secrets.GetBucket()

	if err != nil {
		panic(err)
	}

	// For security use CLI to add new keys
	// Create a group for SQL authentication
    bucket.Set("SQL.Username", "tecno")
	bucket.Set("SQL.pwd", "123")
	bucket.Set("SQL.host", "localhost")
	bucket.Set("SQL.database", "test")


	sql, err := bucket.GetGroup("SQL")
	if err == nil {
		fmt.Println("SQL connection string:")
		fmt.Printf("Server=%v;Database=%v;User Id=%v;Password=%v;\n", sql.Get("host"), sql.Get("database"), sql.Get("Username"), sql.Get("pwd"))
	}
}
```

### Dependencies

- [Google UUID][1]

[1]: https://pkg.go.dev/github.com/google/uuid
[2]: https://github.com/Tecnologer/go-secrets-cli
