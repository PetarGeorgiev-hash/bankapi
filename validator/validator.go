package validator

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func ValidateString(value string, min int, max int) error {
	length := len(value)
	if length < min || length > max {
		return fmt.Errorf("must contain between %d and %d characters", min, max)
	}
	return nil
}

func ValidateUsername(username string) error {
	err := ValidateString(username, 3, 100)
	if err != nil {
		return err
	}

	if ok := isValidUsername(username); !ok {
		return fmt.Errorf("username can only contain alphanumeric characters and underscores")

	}

	return nil
}

func ValidateFullName(full_name string) error {
	err := ValidateString(full_name, 2, 100)
	if err != nil {
		return err
	}

	if ok := isValidFullName(full_name); !ok {
		return fmt.Errorf("full_name can only contain alphanumeric characters and spaces")

	}

	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 5, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email format")
	}
	return nil
}
