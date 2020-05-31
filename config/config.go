package config

import "github.com/google/uuid"

//Config contains configuration for secrets
type Config struct {
	EncryptionEnabled bool
	BucketID          uuid.UUID
	currentPath       string
	// IDSource          strings
}
