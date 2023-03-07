package StateMachine

import "fmt"

type FSM struct {
	name    string
	current *State
	states  []*State
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
	fmt.Printf(fmt.Sprintf("Created a machine names %s\n", machine.name))
	return &machine
}

// Creates a new state on a state machine
// @stateName {String}
func (f *FSM) NewState(stateName string) bool {
	for _, elm := range f.states {
		if elm.name == stateName {
			return false
		}
	}
	var state *State = &State{name: stateName}
	f.states = append(f.states, state)
	return true
}

// Get the Machine name
// @Return {String}: Machine name
func (f FSM) Name() string { return f.name }

// Get the state name
// @Return {String}: State name
func (s State) Name() string { return s.name }
