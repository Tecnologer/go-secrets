package secrets

import (
	"fmt"
	"regexp"
)

//Secret type for secrets bucket
type Secret map[string]interface{}

//NewSecret returns new instance of Secrets
func NewSecret() Secret {
	return Secret{}
}

//Set adds or updates the value for the specific key
func (s Secret) Set(key string, value interface{}) {
	s[key] = value
}

//Get returns the value for the specific key
func (s Secret) Get(key string) interface{} {
	if val, exists := s[key]; exists {
		return val
	}

	return nil
}

//GetString returns the value as string for the specific key
func (s Secret) GetString(key string) string {
	if val, exists := s[key]; exists {
		return fmt.Sprintf("%v", val)
	}

	return ""
}

//GetInt returns the value as int for the specific key
func (s Secret) GetInt(key string) int {
	if val, exists := s[key]; exists {
		return val.(int)
	}

	return 0
}

//Remove removes the key
func (s Secret) Remove(key string) {
	delete(s, key)
}

//GetGroup gets the key with the same group
//
//  - "<group>:<key>", I.e.: "SQL:User"
func (s Secret) GetGroup(group string) (Secret, error) {
	regGroup, err := regexp.Compile(fmt.Sprintf(`(?m)^%s\.`, group))
	if err != nil {
		return nil, err
	}

	secret := NewSecret()

	for key, val := range s {
		if !regGroup.Match([]byte(key)) {
			continue
		}
		key = regGroup.ReplaceAllString(key, "")
		secret.Set(key, val)
	}

	return secret, nil
}
