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
	secrets.Init()

    username := secrets.Get("username")

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
	secrets.Init()

	// For security use CLI to add new keys
	// Create a group for SQL authentication
    secrets.Set("SQL.Username", "tecno")
	secrets.Set("SQL.pwd", "123")
	secrets.Set("SQL.host", "localhost")
	secrets.Set("SQL.database", "test")


	sql, err := secrets.GetGroup("SQL")
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
