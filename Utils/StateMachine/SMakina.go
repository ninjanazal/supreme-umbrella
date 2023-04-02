package FSM

import (
	Operations "SupUmbrela/Utils"
	Logger "SupUmbrela/Utils/Log"
	"fmt"
)

type FSM struct {
	name      string
	active    bool
	current   *State
	states    map[uint32]State
	on_change func(*string)
}

// Get the Machine name
// @Return {String}: Machine name
func (f FSM) Name() string { return f.name }

// Machine constructor
// @name {String}: Machine name
// @Return {FSM}: FSM object
func NewMachine(name string) FSM {
	var machine FSM = FSM{
		name:   name,
		active: false}
	machine.states = make(map[uint32]State)
	Logger.TraceLog(fmt.Sprintf("Created machine named %s", name), Logger.LOG)
	return machine
}

// Sets the on state change callback
// @f_callback {*string}: Callback function for on state change
func (f *FSM) SetOnChangeCallback(f_callback func(*string)) {
	f.on_change = f_callback
}

// Creates a new state on a state machine
// @stateName {String}
func (f *FSM) AddState(stateObj State) bool {
	_, ok := f.states[Operations.Hash(stateObj.name)]

	if ok {
		return false
	}

	f.states[Operations.Hash(stateObj.name)] = stateObj
	return true
}

// Boots the machine on a target state
// @stateName {*String}: Target state name
// Return {bool}: Was able to boot
func (f *FSM) BootMachineOn(stateName string) bool {
	t_state, ok := f.states[Operations.Hash(stateName)]
	if ok {
		f.active = true
		f.current = &t_state
		Logger.TraceLog("Booting machine on "+stateName, Logger.LOG)

		if f.on_change != nil {
			f.on_change(&f.current.name)
		}
		if f.current.on_enter != nil {
			f.current.on_enter()
		}
		return true
	}
	Logger.TraceLog("Failed to boot on "+stateName, Logger.WARN)
	return false
}

// Request a state change by name
// @stateName {String}: Target state name
// Return {bool}:
func (f *FSM) GoToNextState(stateName string) bool {
	if !f.active {
		Logger.TraceLog("Innactive stateMachine, please boot", Logger.WARN)
		return false
	}

	t_state, ok := f.states[Operations.Hash(stateName)]
	if ok {
		if f.current.ValidNext(&stateName) {
			Logger.TraceLog("Change state to "+stateName, Logger.LOG)
			f.current = &t_state
			if f.on_change != nil {
				f.on_change(&f.current.name)
			}
			if f.current.on_enter != nil {
				f.current.on_enter()
			}
			return true
		}
		Logger.TraceLog("Failed to change to state "+stateName+" from "+f.current.name, Logger.WARN)
		return false
	}
	Logger.TraceLog("Failed to change to state "+stateName+", not found ", Logger.WARN)
	return false
}

// Get the current state name if active, else return empty string
// @Return {*String} Pointer to the current state name if defined else nill
func (f *FSM) GetCurrentName() *string {
	if f.active {
		return &f.current.name
	}
	return nil
}

// Gets a pointer to the current state if defined
// @Return {*State} Pointer to the current state f defined else nill
func (f *FSM) GetCurrentState() *State {
	if f.active {
		return f.current
	}
	return nil
}

// Gets the input help string for the current state
// @Return {string} On Input help string
func (f *FSM) OnInputOptions() string {
	if f.current.on_input != nil {
		return f.current.on_input()
	}
	return ""
}

// Process the input for the current state
// @data {*String}: Input string
// @fsm {*FSM}: Pointer to the current machine

func (f *FSM) OnInputProcess(fsm *FSM, data *string) bool {
	if f.current.on_process != nil {
		return f.current.on_process(fsm, data)
	}
	return false
}
