package Operations

import "hash/fnv"

// Get a hash from a string
// @s {*String}: Target string
// @Return {uint32}: Hash value
func Hash(s *string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(*s))
	return h.Sum32()
}

// IsOperator checks if a given character is an arithmetic operator.
// @param char {*byte} - A pointer to a `byte` that represents the character to check.
// @return {bool} - A boolean value indicating whether the given character is an arithmetic operator.
func IsOperator(char *byte) bool {
	return (*char == '/' || *char == '*' || *char == '+' || *char == '-')
}
