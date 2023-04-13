// Calculator
package Calculator

import (
	Classes "SupUmbrela/Utils/Classes"
	Logger "SupUmbrela/Utils/Log"
	"fmt"
)

// RequestOperation parses the given operation string, calculates the result,
// and returns it as an integer
// @param operation {string} - The string containing the operation to calculate.
// @return {int} - Returns the calculated result of the operation as an integer.
func RequestOperation(operation string) float32 {
	Logger.TraceLog(fmt.Sprintf("Calculating operation %s", operation), Logger.LOG)
	valid, value := parseInputStrng(&operation)
	if valid {
		return value
	}
	Logger.TraceLog(fmt.Sprintf("Wrong or invalid input - %s", operation), Logger.LOG)
	return 0
}

// Parses and calculate the input string
func parseInputStrng(operation *string) (bool, float32) {
	// If string is empty
	if *operation == "" {
		return false, 0
	}
	var queueOp Classes.Queue[string] = Classes.Queue[string]{}
	QueueOperation(&queueOp, operation)

	// // var start_operation Operation = Operation{}
	// var valid_op bool = true
	// Minimize(operation)
	// if !ValidateOp(operation) {
	// 	return false, 0
	// }
	// postOp := PostFix(operation)

	// Logger.TraceLog("PostFix operation", Logger.LOG, postOp)

	// return valid_op, 1
	return true, 0
}
