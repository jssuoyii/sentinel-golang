package util

import (
	"encoding/json"
	"strings"
)

// IsBlank checks whether the given string is blank.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func ToString(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}
