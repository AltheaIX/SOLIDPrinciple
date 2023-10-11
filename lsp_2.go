package main

/*
Vehicle and Engines with Single Responsibility Principle, Open-Closed Principle, and Liskov Substitution Principle
Checked and reviewed by ChatGPT
*/

const (
	CAR        string = "car"
	MOTORCYCLE        = "motorcycle"
)

type Engine interface {
	Start() string
	Stop() string
}

type GasEngine struct{}

func (ge *GasEngine) Start() string {
	return "[GasEngineStart] Gas engine started."
}

func (ge *GasEngine) Stop() string {
	return "[GasEngineStop] Gas engine stopped."
}

type ElectricEngine struct{}

func (ge *ElectricEngine) Start() string {
	return "[ElectricEngineStart] Electric engine started."
}

func (ge *ElectricEngine) Stop() string {
	return "[ElectricEngineStop] Electric engine stopped."
}

type Vehicle interface {
	StartEngine() string
	StopEngine() string
}

type Car struct {
	engine Engine
}

func (c *Car) StartEngine() string {
	return c.engine.Start() + "\n[CarStartEngine] Engine started and vehicle ready to go."
}

func (c *Car) StopEngine() string {
	return c.engine.Stop() + "\n[CarStopEngine] Engine stopped and vehicle is steady now."
}

type Motorcycle struct {
	engine Engine
}

func (m *Motorcycle) StartEngine() string {
	return m.engine.Start() + "\n[MotorcycleStartEngine] Engine started and vehicle ready to go."
}

func (m *Motorcycle) StopEngine() string {
	return m.engine.Stop() + "\n[MotorcycleStopEngine] Engine started and vehicle is steady now."
}

func Drive(v Vehicle) string {
	return v.StartEngine() + "\n[Drive] You are driving now."
}

func NewVehicle(vehicleType string, engine Engine) Vehicle {
	switch vehicleType {
	case CAR:
		return &Car{engine: engine}
	case MOTORCYCLE:
		return &Motorcycle{engine: engine}
	default:
		return nil
	}
}
