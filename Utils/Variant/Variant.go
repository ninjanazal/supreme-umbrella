package Variant

import "reflect"

// integerTypes is a slice of reflect.Type values that represents all
// the integer types in Go. This includes signed and unsigned integers
// with different sizes (8-bit, 16-bit, 32-bit, and 64-bit).
var integerTypes = []reflect.Type{
	reflect.TypeOf(int(0)),    // Signed integer, size depends on platform
	reflect.TypeOf(int8(0)),   // 8-bit signed integer
	reflect.TypeOf(int16(0)),  // 16-bit signed integer
	reflect.TypeOf(int32(0)),  // 32-bit signed integer
	reflect.TypeOf(int64(0)),  // 64-bit signed integer
	reflect.TypeOf(uint(0)),   // Unsigned integer, size depends on platform
	reflect.TypeOf(uint8(0)),  // 8-bit unsigned integer
	reflect.TypeOf(uint16(0)), // 16-bit unsigned integer
	reflect.TypeOf(uint32(0)), // 32-bit unsigned integer
	reflect.TypeOf(uint64(0)), // 64-bit unsigned integer
}

// Variant object type
type Variant struct {
	data   string
	i_data interface{}
	o_type reflect.Type
}

// IsInteger checks if a value is an integer by comparing its type to
// the types in the integerTypes slice. If the value's type matches
// one of the integer types, it returns true. Otherwise, it returns false.
// This function uses reflection to determine the type of the value.
// @value {interface}: Test value
// @Return {bool}: Is an integer
func IsInteger(value interface{}) bool {
	for _, tp := range integerTypes {
		if reflect.TypeOf(value) == tp {
			return true
		}
	}
	return false
}

// Adds two variant type
// @right {Variant} Right operation value
// @Return {Variant}: Result variant object
func (v *Variant) Add(right Variant) Variant {
	var result Variant = Variant{}
	result.o_type = v.o_type

	// reflect.TypeOf()

	return result
}
