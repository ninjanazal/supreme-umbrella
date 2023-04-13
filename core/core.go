package SupCore

import (
	"SupUmbrela/Calculator"

	"bufio"
	"fmt"
	"os"

	Logger "SupUmbrela/Utils/Log"
	SMachine "SupUmbrela/Utils/StateMachine"
)

// Executes the core main loop
func Run() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	Logger.TraceLog("-- Learning go collection --", Logger.LOG, "")
	Logger.TraceLog("Select a project to launch", Logger.LOG, "")

	Calculator.RequestOperation("test")
	fsm := SMachine.NewMachine("LearnMachine")

	CreateCoreFSM(&fsm)
	fsm.BootMachineOn("selection")

	// Application main loop
	for {
		InputOptions(&fsm)
		fmt.Print("what next >> ")
		scanner.Scan()
		var input_data string = scanner.Text()

		Logger.TraceLog("Inserted data", Logger.LOG, "data : { "+input_data+" }")
		if !ProcessInput(&fsm, &input_data) {
			break
		}
	}

	Logger.TraceLog("Closing application.", Logger.WARN)
}
