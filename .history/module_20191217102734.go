package main

// Module to launch
type Module struct {
	Mass int
}

// RequiredFuel return how much fuled does module need to lift off
func (module *Module) RequiredFuel() int {
	return int(module.Mass/3) - 2
}

func requiredFuelForMass() int {
	return int(module.Mass/3) - 2
}

func requiredFuelForFuel(fuelMass int) int {

}

// NewModule return new instance of Module
func NewModule(mass int) *Module {
	return &Module{
		Mass: mass,
	}
}