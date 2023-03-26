// Calculator
package Calculator

import (
	Logger "SupUmbrela/Utils/Log"
	"fmt"
)

// validOperators is an array of strings containing the valid arithmetic operators for an expression.
var validOperators [4]string = [4]string{"+", "-", "*", "/"}

// Process a requested operation as string
// @operation {String}: Requested operation
func RequestOperation(operation string) int {
	Logger.TraceLog(fmt.Sprintf("Calculating operation %s", operation), Logger.LOG)
	valid, value := parseInputStrng(&operation)
	if valid {
		return value
	}
	Logger.TraceLog(fmt.Sprintf("Wrong or invalid input - %s", operation), Logger.LOG)
	return 0

}

type Operation struct {
	l_op int
	r_op int
	op   string
	next *Operation
}

// Parses and calculate the input string
func parseInputStrng(operation *string) (bool, int) {
	// var start_operation Operation = Operation{}
	var valid_op bool = true
	// var op_array []string = strings.Split(*operation, "")

	return valid_op, 1
}

func convert_values(inputData *[]string) {
	
}
