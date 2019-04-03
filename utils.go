package check

import (
	"errors"
	"strings"
	"unicode"
)

func requiredErr(required bool, message string) error {
	if !required {
		return nil
	}
	if message = strings.TrimSpace(message); message != "" {
		return errors.New(message)
	}

	return errEmpty
}

func stripSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}

		return r
	}, s)
}

func isEmptyStr(field string) bool {
	return strings.TrimSpace(field) == ""
}
