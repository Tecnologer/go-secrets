package secrets

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func getDecryptedData(id uuid.UUID, path string) ([]byte, error) {
	encrKey := getEncrypKey(id)
	fileContent, err := ioutil.ReadFile(secretFilePath)

	if err != nil {
		return nil, errors.Wrap(err, "error reading secret file")
	}

	c, err := aes.NewCipher(encrKey)
	if err != nil {
		return nil, errors.Wrap(err, "error creating cipher")
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Galois Counter Mode")
	}

	nonceSize := gcm.NonceSize()
	if len(fileContent) < nonceSize {
		return nil, fmt.Errorf("invalid NonceSize")
	}

	nonce, ciphertext := fileContent[:nonceSize], fileContent[nonceSize:]
	json, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error decrypting secret file")
	}

	return json, nil
}

func encriptData(id uuid.UUID, content []byte) ([]byte, error) {
	encrKey := getEncrypKey(id)
	c, err := aes.NewCipher(encrKey)
	// if there are any errors, handle them
	if err != nil {
		return nil, errors.Wrap(err, "error creating cipher")
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Galois Counter Mode")
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.Wrap(err, "error creating random sequence")
	}

	return gcm.Seal(nonce, nonce, content, nil), nil
}

func getEncrypKey(id uuid.UUID) []byte {
	return []byte(id.String() + currentPath)
}
