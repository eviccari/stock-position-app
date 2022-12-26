package utils

import "strings"

func IsEmptyString(v string) bool {
	return v == "" || strings.TrimSpace(v) == ""
}
