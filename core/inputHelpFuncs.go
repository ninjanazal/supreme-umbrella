package SupCore

import (
	Logger "SupUmbrela/Utils/Log"
	SMachine "SupUmbrela/Utils/StateMachine"
	"fmt"

	Calculator "SupUmbrela/Calculator"
)

// ON Input actions for the selector state
// Return {string}: Input available actions string
func OnSelectorInputActions() string {
	var outstring string = ""
	outstring += "-> '-calc' : Goes to the calculator\n"
	return outstring
}

// On Process input for the selector state
// @data {*string}: Pointer to the input string
// @fsm {*FSM}: Current state machine pointer
// Retunr {bool}: Was able to process the input string
func OnSelectorProcessInput(fsm *SMachine.FSM, data *string) bool {
	switch *data {
	case "-calc":
		Logger.TraceLog("Changing to the calculator ...", Logger.LOG)
		fsm.GoToNextState("calculator")
		return true
	}
	return false
}

// On input actions for the calculator state
// Return {string}: Input available actions string
func OnCalculatorInputActions() string {
	return "-> Please insert a expression to be calculated\n"
}

// On Process input for the calculator state
func OnCalculatorProcessInput(fsm *SMachine.FSM, data *string) bool {
	var response float32 = Calculator.RequestOperation(*data)
	Logger.TraceLog("= "+fmt.Sprint(response), Logger.LOG)
	return true
}
