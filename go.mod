module learn/main

go 1.20

replace learn/Calculator => ./Calculator

replace learn/StateMachine => ./Utils/StateMachine

require (
	learn/Calculator v0.0.0-00010101000000-000000000000
	learn/StateMachine v0.0.0-00010101000000-000000000000
)
