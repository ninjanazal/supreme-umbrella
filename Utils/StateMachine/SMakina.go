package Utils

import (
	"fmt"
	Logger "learn/Utils/Log"
)

type FSM struct {
	name    string
	current *State
	states  map[*string]State
}

type State struct {
	name string
	next []*State
}

// Machine constructor
// @name {String}: Machine name
// @Return {*FSM}: Pointer to a new FSM
func NewMachine(name string) *FSM {
	var machine FSM = FSM{name: name}
	fmt.Println("Created a machine names " + machine.name)
	Logger.TraceLog("Test log", Logger.LOG, "")
	return &machine
}

// Creates a new state on a state machine
// @stateName {String}
func (f *FSM) NewState(stateName string) bool {

	_, ok := f.states[&stateName]

	if ok {
		return false
	}

	var state *State = &State{name: stateName}
	f.states[&stateName] = *state
	return true
}

// Get the Machine name
// @Return {String}: Machine name
func (f FSM) Name() string { return f.name }

// Get the state name
// @Return {String}: State name
func (s State) Name() *string { return &s.name }
