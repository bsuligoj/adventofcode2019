package main

// Module to launch
type Module struct {
	Mass int
}

// RequiredFuel return how much fuled does module need to lift off
func (module *Module) RequiredFuel() int {
	return 0
}

//
func NewModule(mass int) *Module {
	return &Module{
		Mass: mass,
	}
}
