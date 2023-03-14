// Main package for the lerning projects
// This will act as a launcher for the different created projects
package main

import (
	"learn/Calculator"

	Logger "learn/Utils/Log"
	SMachine "learn/Utils/StateMachine"
)

// === === === === === === === === ===
func main() {
	Logger.TraceLog("-- Learning go collection --", Logger.LOG, "")
	Logger.TraceLog("Select a project to launch", Logger.LOG, "")

	Calculator.RequestOperation("test")
	fsm := SMachine.NewMachine("LearnMachine")

	CreateCoreFSM(&fsm)
	fsm.BootMachineOn("selection")
}
