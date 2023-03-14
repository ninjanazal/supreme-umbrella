package Utils

import (
	Operations "learn/Utils"
)

type State struct {
	name     string
	next     map[uint32]*State
	on_enter func()
}

// State Constructor
// @name {String}: State name
// @Return {State}: State object
func NewState(name string) State {
	var state State = State{name: name}
	state.next = make(map[uint32]*State)
	return state
}

// Adds a next valide state to a state
// @next {*State}: Next state
func (s *State) AddNextState(nextS *State) bool {
	_, ok := s.next[Operations.Hash(&nextS.name)]
	if !ok {
		s.next[Operations.Hash(&nextS.name)] = nextS
		return true
	}
	return false
}

// Defines the on_enter callback fucntion
// @cBack {func} Callback function
// @@Parameters
func (s *State) SetOnEnterCallback(cBack func()) {
	s.on_enter = cBack
}

// Gets if a state by name is a valid next state
func (s *State) ValidNext(next_state *string) bool {
	_, ok := s.next[Operations.Hash(next_state)]
	return bool(ok)
}

// Get the state name
// @Return {String}: State name
func (s State) Name() *string { return &s.name }
