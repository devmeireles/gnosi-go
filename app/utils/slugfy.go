package utils

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile("[^a-z0-9]+")

// Slugfy returns a slugfy string from a text
func Slugfy(s string) string {
	return strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
