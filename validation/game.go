package validation

import (
	"errors"

	h "github.com/forumGamers/tour-service/helpers"
)

func ValidateCreateGame(name string, types string, description string) error {
	if name == "" || types == "" || description == "" {
		return errors.New("Invalid data")
	}

	if err := h.ValidateInput(map[string]string{
		"name":name,
		"types":types,
		"description":description,
	}) ; err != nil {
		return err
	}

	return nil
}