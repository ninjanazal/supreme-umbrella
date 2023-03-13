package utils

import "hash/fnv"

// Get a hash from a string
// @s {*String}: Target string
func Hash(s *string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(*s))
	return h.Sum32()
}
