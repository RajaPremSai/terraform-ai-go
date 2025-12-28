package utils

import "strings"

func ToPtr[T any](i T) *T {
	return &i
}

func RemoveBlankLinesFromString(input string) string {
	return strings.TrimLeft(input, "\n\r \t")
}
