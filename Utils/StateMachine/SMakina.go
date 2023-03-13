package Utils

import (
	"fmt"
	Operations "learn/Utils"
	Logger "learn/Utils/Log"
)

type FSM struct {
	name    string
	current *State
	states  map[uint32]State
}

// Machine constructor
// @name {String}: Machine name
// @Return {FSM}: FSM object
func NewMachine(name string) FSM {
	var machine FSM = FSM{name: name}
	fmt.Println("Created a machine names " + machine.name)
	Logger.TraceLog("Test log", Logger.LOG, "")
	return machine
}

// Creates a new state on a state machine
// @stateName {String}
func (f *FSM) AddState(stateObj State) bool {
	_, ok := f.states[Operations.Hash(&stateObj.name)]

	if ok {
		return false
	}

	f.states[Operations.Hash(&stateObj.name)] = stateObj
	return true
}

// Get the Machine name
// @Return {String}: Machine name
func (f FSM) Name() string { return f.name }

// - - - - - - - - - - - - - - - - - - - -
// - - - - - - - - - - - - - - - - - - - -

type State struct {
	name string
	next map[uint32]*State
}

// State Constructor
// @name {String}: State name
// @Return {State}: State object
func NewState(name string) State {
	var state State = State{name: name}
	return state
}

// Adds a next valide state to a state
// @next {*State}: Next state
func (s State) AddNextState(nextS *State) bool {
	_, ok := s.next[Operations.Hash(&nextS.name)]
	if !ok {
		s.next[Operations.Hash(&nextS.name)] = nextS
		return true
	}
	return false
}

// Get the state name
// @Return {String}: State name
func (s State) Name() *string { return &s.name }
