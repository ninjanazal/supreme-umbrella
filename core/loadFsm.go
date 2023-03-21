package SupCore

import (
	Logger "SupUmbrela/Utils/Log"
	SMachine "SupUmbrela/Utils/StateMachine"
	"fmt"
	"strings"
)

var outstring strings.Builder

// Creates the base finite state machine
// @fsm {*SMachine.FSM} Target state machine object
func CreateCoreFSM(fsm *SMachine.FSM) {
	selectionState := SMachine.NewState("selection")
	calculatorState := SMachine.NewState("calculator")

	selectionState.SetonInputActions(OnSelectorInputActions)
	selectionState.SetonInputProcess(OnSelectorProcessInput)

	calculatorState.SetonInputActions(OnCalculatorInputActions)

	fsm.AddState(selectionState)
	fsm.AddState(calculatorState)

	selectionState.AddNextState(&calculatorState)
}

// Prints the current options
func InputOptions(fsm *SMachine.FSM) bool {
	if fsm.GetCurrentName() == nil {
		Logger.TraceLog("Not a valid state name", Logger.ERRO)
		return false
	}

	outstring.Reset()
	outstring.WriteString("- - - - - - - - - -\n :: Instructions ::")
	outstring.WriteString("-> '-e' : Closes the application\n")
	outstring.WriteString(fsm.OnInputOptions())

	fmt.Println(outstring.String())
	return true
}

// Process the user input data
// @data {*string}: Input data
func ProcessInput(fsm *SMachine.FSM, data *string) bool {
	if fsm.OnInputProcess(fsm, data) {
		return true
	}

	// Default process input
	switch *data {
	case "-e":
		return false
	}

	Logger.TraceLog("Invalid input, please attend to the instructions", Logger.LOG)
	return true
}
