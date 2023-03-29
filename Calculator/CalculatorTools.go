package Calculator

import (
	Opr "SupUmbrela/Utils"
	Classes "SupUmbrela/Utils/Classes"
	"strings"
)

// Minimize removes all spaces from the given string by replacing all spaces with an empty string.
// @param data {*string} - A pointer to a string that needs to be minimized.
func Minimize(data *string) {
	*data = strings.ReplaceAll(*data, " ", "")
}

// PostFix converts an infix expression to postfix notation.
// @param data {*string} - A pointer to a string that represents the infix expression to convert.
// @return {string} - A string that represents the postfix notation of the input expression.
func PostFix(data *string) string {
	var dataStack *Classes.Stack[byte] = Classes.NewStack[byte]()
	var result strings.Builder = strings.Builder{}

	for _, lookInChar := range *data {
		switch lookInChar {
		case '(':
			dataStack.Push(byte(lookInChar))
		case ')':
			for dataStack.Peek() != '(' {
				result.WriteByte(dataStack.Pop())
				result.WriteByte(' ')
			}
			dataStack.Pop()
		case '+', '-', '*', '/', '^':
			result.WriteByte(byte(lookInChar))
			result.WriteByte(' ')

			for i := 1; i < len(*data); i++ {
				nextChar := (*data)[1]
				if nextChar == '(' {
					dataStack.Push(nextChar)
				} else if Opr.IsOperator(&nextChar) || nextChar == '-' {
					result.WriteByte(byte(nextChar))
					result.WriteByte(' ')
					i++
				} else {
					break
				}
			}
		default:
			result.WriteString(string(lookInChar))
			result.WriteByte(' ')
		}
	}
	for !dataStack.IsEmpty() {
		result.WriteByte(dataStack.Pop())
		result.WriteByte(' ')
	}
	return result.String()
}
