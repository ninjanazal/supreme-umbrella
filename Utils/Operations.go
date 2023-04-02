package Operations

import (
	"hash/fnv"
)

// Get a hash from a string
// @s {*String}: Target string
// @Return {uint32}: Hash value
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

var mapOperators []byte = []byte{
	'+', '-', '*', '/', '%',
}

// isOperator checks if the input byte is a supported operator.
// char (*byte): A pointer to the byte to check.
// Return:
// - {bool}: True if the byte is a supported operator, false otherwise.
func IsOperator(char *byte) bool {
	for _, elm := range mapOperators {
		if elm == *char {
			return true
		}
	}
	return false
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
