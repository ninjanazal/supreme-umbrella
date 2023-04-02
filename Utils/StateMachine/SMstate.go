package FSM

import (
	Operations "SupUmbrela/Utils"
)

// State represents a state in a finite state machine. It contains a name, a map of next states,
// and optional functions to execute when the state is entered, when an input is received, and
// when the state is processed. The next states are indexed by their hash value.
type State struct {
	name       string                   // The name of the state.
	next       map[uint32]*State        // A map of the next states, indexed by their hash value.
	on_enter   func()                   // An optional function to execute when the state is entered.
	on_input   func() string            // An optional function to execute when an input is received.
	on_process func(*FSM, *string) bool // An optional function to execute when the state is processed.
}

// NewState creates a new state with the given name and initializes its "next" map.
// @param name {string} - The name of the new state.
// @return {State} - Returns a new State object with the given name and an empty "next" map.
func NewState(name string) State {
	var state State = State{name: name}
	state.next = make(map[uint32]*State)
	return state
}

// AddNextState adds the given next state to the map of next states for this state.
// It uses the hash of the next state's name to index the map.
// @param nextS {*State} - A pointer to the next state to add to the map.
// If the next state is not already in the map, it adds it and returns true.
// Otherwise, it does nothing and returns false.
// @return {bool} - Returns true if the next state was added to the map, false if it was already present
func (s *State) AddNextState(nextS *State) bool {
	_, ok := s.next[Operations.Hash(nextS.name)]
	if !ok {
		s.next[Operations.Hash(nextS.name)] = nextS
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
	_, ok := s.next[Operations.Hash(*next_state)]
	return bool(ok)
}

// Get the state name
// @Return {String}: State name
func (s *State) Name() *string { return &s.name }

// Sets the Callback funtion on state should display available optons
// @onInput {func() string}: On state request actions
func (s *State) SetonInputActions(onInput func() string) { s.on_input = onInput }

// Set the on input process callback function
// @onProcess {func() bool}: On state input process action
func (s *State) SetonInputProcess(onProcess func(*FSM, *string) bool) { s.on_process = onProcess }
