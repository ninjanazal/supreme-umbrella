package Calculator

import (
	Operations "SupUmbrela/Utils"
	Classes "SupUmbrela/Utils/Classes"
	"bytes"
	"strings"
	"unicode"
)

// Minimize removes all spaces from the given string by replacing all spaces with an empty string.
// @param data {*string} - A pointer to a string that needs to be minimized.
func Minimize(data *string) {
	*data = strings.ReplaceAll(*data, " ", "")
}

// QueueOperation takes a string input and a pointer to a Queue of strings and
// separates the input into individual numbers and operators, adding each to
// the queue. Any character that is not a digit or a semicolon is treated as
// an operator. If the input ends with a number, it will be added to the queue
// after processing the input.
// stkOut (*Classes.Queue[string]): A pointer to the output queue that will
// hold the separated input.
// data (*string): A pointer to the input string to separate.
func QueueOperation(stkOut *Classes.Queue[string], data *string) {
	var dataP bytes.Buffer = bytes.Buffer{}

	for i := 0; i < len(*data); i++ {
		character := (*data)[i]
		isSemicolum := Operations.IsSemicolum(&character)

		if !unicode.IsDigit(rune(character)) && !isSemicolum {
			if dataP.Len() > 0 {
				stkOut.Add(dataP.String())
				dataP.Reset()
			}
			stkOut.Add(string(character))
		} else {
			if isSemicolum {
				dataP.WriteByte('.')
			} else {
				dataP.WriteByte(character)
			}
		}
	}
	if dataP.Len() > 0 {
		stkOut.Add(dataP.String())
		dataP.Reset()
	}
}
