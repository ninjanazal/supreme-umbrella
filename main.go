// Main package for the lerning projects
// This will act as a launcher for the different created projects
package main

import (
	"fmt"
	"learn/Calculator"
	"learn/StateMachine"
)

func main() {
	fmt.Println("-- Learning go collection --")

	fmt.Print("Select a project to launch\n")
	Calculator.RequestOperation("test")

	fsm := StateMachine.NewMachine("LearnMachine")
	fmt.Println(fsm.Name())
}
