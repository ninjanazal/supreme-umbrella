package Operations

import (
	Logger "SupUmbrela/Utils/Log"
	"fmt"
	"hash/fnv"
	"strconv"
)

// Get a hash from a string
// @s {*String}: Target string
// @Return {uint32}: Hash value
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// Isparentheses checks if the input byte is a left or right parentheses.
// char (*byte): A pointer to the byte to check.
// Return {bool}: True if the byte is a left or right parentheses, false otherwise.
func Isparentheses(char *byte) bool {
	return *char == '(' || *char == ')'
}

// IsSemicolon checks if the input byte is a semicolon or a comma.
// char (*byte): A pointer to the byte to check.
// Return {bool}: True if the byte is a semicolon or comma, false otherwise.
func IsSemicolum(char *byte) bool {
	return *char == '.' || *char == ','
}

// ToFloat32 converts the given string to a float32 value.
// value (*string): A pointer to the string to be converted.
// Returns {float32}: The converted float32 value.

func ToFloat32(value *string) float32 {
	v, err := strconv.ParseFloat(*value, 32)
	if err != nil {
		Logger.TraceLog("Failed to parse value", Logger.ERRO, fmt.Sprintf("Value: %s", *value))
	}
	return float32(v)
}
