package secrets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	//name of the secret file, where all secrets key are in
	secretFile = ".secret"
	secretsDir = ".go-secrets"

	//initSecretFilePath is the path where is the local file that contains the id of the bucket
	initSecretFilePath = ".secretid"
)

var (
	secretFilePath      string
	currentPath         string
	localSecretFilePath string
)

//Bucket struct for collection of secrets
type Bucket struct {
	ID      uuid.UUID
	Secrets Secret
}

func init() {
	currentPath, _ = os.Getwd()

	if currentPath == "" {
		currentPath = "."
	}
	localSecretFilePath = fmt.Sprintf("%s/%s", currentPath, initSecretFilePath)
	fmt.Println(localSecretFilePath)
}

//GetBucket creates or return the bucket with the specific UUID in the local secret file
func GetBucket() (*Bucket, error) {
	if !secretExists(localSecretFilePath) {
		return nil, fmt.Errorf("Cannot get the bucket, the secret file is not initialized, use \"go-secrets-cli init\" to create it")
	}

	var secretContent []byte
	secretContent, err := ioutil.ReadFile(localSecretFilePath)

	if err != nil {
		return nil, errors.Wrap(err, "Cannot get the bucket, the secret file is corrupted, use \"go-secrets-cli init\" to repair it")
	}

	return GetBucketByID(string(secretContent))
}

//GetBucketByID creates or return the bucket with the specific UUID as string
func GetBucketByID(id string) (*Bucket, error) {
	bucketID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot get the bucket, the secret file is corrupted, use \"go-secrets-cli init\" to repair it")
	}

	return GetBucketByUUID(bucketID)
}

//GetBucketByUUID creates or return the bucket with the specific UUID
func GetBucketByUUID(id uuid.UUID) (*Bucket, error) {
	data, err := readData(id)
	if err != nil {
		logrus.WithError(err).Warning("GetBucketByUUID: error reading stored data")
	}

	return &Bucket{
		ID:      id,
		Secrets: data,
	}, nil
}

func readData(id uuid.UUID) (Secret, error) {
	var err error
	if secretFilePath == "" {
		secretFilePath, err = getSecretPath(id)
		if err != nil {
			return NewSecret(), errors.Wrap(err, "readData: error getting secret path")
		}

		// fmt.Println(secretFilePath)
		//if the secret file doesn't exists
		if !secretExists(secretFilePath) {
			newSecret := NewSecret()
			writeSecrets(newSecret)
			return newSecret, nil
		}
	}
	file, err := ioutil.ReadFile(secretFilePath)

	if err != nil {
		logrus.WithError(err).Warning("error getting data from secret bucket file")
		file = []byte{}
	}

	data := NewSecret()
	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		return data, errors.Wrap(err, "readata: error parsing file content to json")
	}

	return data, nil
}

func getSecretPath(id uuid.UUID) (string, error) {
	homeDir, err := getUserDir()
	if err != nil {
		return "", err
	}
	secretDir := fmt.Sprintf("%s/%s/%s", homeDir, secretsDir, id)
	if _, err := os.Stat(secretDir); os.IsNotExist(err) {
		errDir := os.MkdirAll(secretDir, 0755)
		if errDir != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s/%s", secretDir, secretFile), nil
}

func getUserDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func secretExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func writeSecrets(secret Secret) error {
	json, err := json.MarshalIndent(secret, "", "")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(secretFilePath, json, 0644)
}

//IO

//Set adds or updates the value for the specific key
func (b *Bucket) Set(key string, value interface{}) {
	b.Secrets.Set(key, value)
	writeSecrets(b.Secrets)
}

//Get returns the value for the specific key
func (b *Bucket) Get(key string) interface{} {
	return b.Secrets.Get(key)
}

//Remove removes the key
func (b *Bucket) Remove(key string) {
	b.Secrets.Remove(key)
	writeSecrets(b.Secrets)
}

//GetGroup gets the key with the same group
//
//  - "<group>:<key>", I.e.: "SQL:User"
func (b *Bucket) GetGroup(group string) (Secret, error) {
	return b.Secrets.GetGroup(group)
}
