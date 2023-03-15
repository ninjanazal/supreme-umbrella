// Main package for the lerning projects
// This will act as a launcher for the different created projects
package main

import (
	"bufio"
	"fmt"
	"learn/Calculator"
	"os"

	Logger "learn/Utils/Log"
	SMachine "learn/Utils/StateMachine"
)

// === === === === === === === === ===
func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	Logger.TraceLog("-- Learning go collection --", Logger.LOG, "")
	Logger.TraceLog("Select a project to launch", Logger.LOG, "")

	Calculator.RequestOperation("test")
	fsm := SMachine.NewMachine("LearnMachine")

	CreateCoreFSM(&fsm)
	fsm.BootMachineOn("selection")

	for {
		InputOptions(fsm.GetCurrentName())
		fmt.Print("what next >> ")
		scanner.Scan()
		var input_data string = scanner.Text()

		Logger.TraceLog("Inserted data", Logger.LOG, "data : { "+input_data+" }")
		if !ProcessInput(&input_data) {
			break
		}
	}

	Logger.TraceLog("Closing application.", Logger.WARN)
}
