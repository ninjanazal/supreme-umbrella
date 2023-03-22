// Calculator
package Calculator

import (
	Logger "SupUmbrela/Utils/Log"
	"fmt"
)

// Process a requested operation as string
// @operation {String}: Requested operation
func RequestOperation(operation string) int {
	Logger.TraceLog(fmt.Sprintf("Calculating operation %s", operation), Logger.LOG)
	return -1
}
