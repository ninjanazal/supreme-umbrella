package main

import (
	"fmt"
	Logger "learn/Utils/Log"
	SMachine "learn/Utils/StateMachine"
)

// Creates the base finite state machine
// @fsm {*SMachine.FSM} Target state machine object
func CreateCoreFSM(fsm *SMachine.FSM) {
	selectionState := SMachine.NewState("selection")
	calculatorMain := SMachine.NewState("calculator_main")

	selectionState.AddNextState(&calculatorMain)

	fsm.AddState(selectionState)
	fsm.AddState(calculatorMain)
}

// Prints the current options
func InputOptions(state_name *string) bool {
	if state_name == nil {
		Logger.TraceLog("Not a valid state name", Logger.ERRO)
		return false
	}

	var outstring string = "- - - - - - - - - -\n :: Instructions ::\n"
	switch *state_name {
	case "selection":
		outstring += "-> '-e' : Closes the application\n"

	}
	fmt.Println(outstring)
	return true
}

// Process the user input data
// @data {*string}: Input data
func ProcessInput(data *string) bool {
	switch *data {
	case "-e":
		return false
	}

	Logger.TraceLog("Invalid input, please attend to the instructions", Logger.LOG)
	return true
}
