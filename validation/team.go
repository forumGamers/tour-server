package validation

import (
	"errors"

	h "github.com/forumGamers/tour-service/helpers"
)

func ValidateCreateTeam(name string,description string,players []string) error {
	if name == "" {
		return errors.New("Invalid data")
	}

	if len(players) < 1 {
		return errors.New("Invalid data")
	}

	if h.ValidateInvalidCharacter(name) {
		return errors.New("name do not allow contains symbol")
	}

	if h.ValidateInvalidCharacter(description) {
		return errors.New("description do not allow contains symbol")
	}

	return nil
}