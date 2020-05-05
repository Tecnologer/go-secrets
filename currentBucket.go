package secrets

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/go-secrets/config"
)

//CurrentBucket is the instance of the bucket for the secrets in the current instance
var (
	CurrentBucket *Bucket
	bucketConfig  *config.Config

	uuidEmpty = uuid.UUID{}
)

//Init secrets bucket for this project using the ID from .secretid
//
//** Don't forget initialize with => go-secret-cli init
func Init() {
	_, err := Get()
	log.WithError(err).Warningf("error initializing the secret bucket")
}

//InitWithConfig inits secrets bucket for this project using the configuration
//
//** Don't forget initialize with => go-secret-cli init
func InitWithConfig(conf *config.Config) {
	bucketConfig = conf
	var err error

	//initialize the bucket with specific ID
	if conf.BucketID != uuidEmpty && (CurrentBucket == nil || CurrentBucket.ID != conf.BucketID) {
		CurrentBucket, err = GetBucketByUUID(conf.BucketID)
		log.WithError(err).Warningf("error initializing the secret bucket")
		if err != nil {
			return
		}
	} else {
		Init()
	}
}

//Get returns the current bucket
func Get() (*Bucket, error) {
	var err error
	if CurrentBucket == nil {
		CurrentBucket, err = GetBucket()
		if err != nil {
			return nil, err
		}
	}

	return CurrentBucket, nil
}

//GetKey returns the value of the key in the current bucket
func GetKey(key string) interface{} {
	b, err := Get()

	if err != nil {
		log.WithError(err).Warningf("error getting the key \"%s\"", key)
		return nil
	}

	return b.Get(key)
}

//GetGroup returns the group in the current bucket
func GetGroup(group string) (Secret, error) {
	b, err := Get()

	if err != nil {
		log.WithError(err).Warningf("error getting the group \"%s\". Err: %q", group)
		return nil, err
	}

	return b.GetGroup(group)
}

//Set adds or updates the value for the specific key in the current bucket
func Set(key string, val interface{}) {
	b, err := Get()

	if err != nil {
		log.WithError(err).Warningf("error when try set the key \"%s\". Err: %q", key)
		return
	}

	b.Set(key, val)
}

//Remove removes the key in the current bucket
func Remove(key string, val interface{}) {
	b, err := Get()

	if err != nil {
		log.WithError(err).Warningf("error when try remove the key \"%s\". Err: %q", key)
		return
	}

	b.Remove(key)
}
