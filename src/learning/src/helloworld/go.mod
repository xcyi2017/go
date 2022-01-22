module helloworld

go 1.14

require (
	github.com/myuser/calculator v0.0.0 //ct
	github.com/rs/zerolog v1.26.1 // indirect
)

replace github.com/myuser/calculator => ../calculator
