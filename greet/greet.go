package greet

import (
	"strings"
)

func Greet(name string) string {
	cleanName := strings.TrimSpace(name)

	if cleanName == "" {
		return "Hello, stranger! Welcome to Go."
	}

	return "Hello, " + cleanName + "! Welcome to Go."
}
