package greet

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	if name == "" {
		fmt.Println("Hello, stranger! Welcome to Go.")
	} else {
		fmt.Println("Hello, ", strings.TrimSpace(name)+"! Welcome to Go.")
	}
	return ""
}
