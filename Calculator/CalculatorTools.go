package Calculator

import "strings"

// Minimize removes all spaces from the given string by replacing all spaces with an empty string.
// @param data {*string} - A pointer to a string that needs to be minimized.
func Minimize(data *string) {
	*data = strings.ReplaceAll(*data, " ", "")
}

func PostFix(data *string) string {
}
