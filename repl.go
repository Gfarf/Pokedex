package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	res := strings.Fields(text)
	return res
}
