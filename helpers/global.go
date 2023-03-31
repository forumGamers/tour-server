package helpers

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateInput(data map[string]string) error{
	for key, value := range data {
		if !strings.Contains(strings.ToLower(key),"password") && ValidateInvalidCharacter(value){
			return errors.New("Invalid data")
		}
	}

	return nil
}

func ValidateInvalidCharacter(data string) bool {
	return regexp.MustCompile(`[^a-zA-Z0-9.,\-\s@]`).MatchString(data)
}