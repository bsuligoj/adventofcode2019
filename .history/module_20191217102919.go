package main

// Module to launch
type Module struct {
	Mass         int
	RequiredFuel int
}

// CalculateRequiredFuel return how much fuled does module need to lift off
func (module *Module) calculateRequiredFuel() int {
	return int(module.Mass/3) - 2
}

func requiredFuelForMass(mass int) int {
	return int(mass/3) - 2
}

func requiredFuelForFuel(fuelMass int) int {

}

// NewModule return new instance of Module
func NewModule(mass int) *Module {
	module := &Module{
		Mass: mass,
	}
	return module
}
