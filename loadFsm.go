package main

import (
	SMachine "learn/Utils/StateMachine"
)

func CreateCoreFSM(fsm *SMachine.FSM) {
	selectionState := SMachine.NewState("selection")
	calculatorMain := SMachine.NewState("calculator_main")

	selectionState.AddNextState(&calculatorMain)

	fsm.AddState(selectionState)
	fsm.AddState(calculatorMain)
}
