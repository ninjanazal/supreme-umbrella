module learn/main

go 1.20

replace learn/Utils => ./Utils

replace learn/Calculator => ./Calculator

require (
	learn/Calculator v0.0.0-00010101000000-000000000000
	learn/Utils v0.0.0-00010101000000-000000000000
)
