package main

import (
	"testing"
)

/*
Vehicle and Engines with Single Responsibility Principle, Open-Closed Principle, and Liskov Substitution Principle
Checked and reviewed by ChatGPT
*/

func TestCarGasEngine_StartEngine(t *testing.T) {
	vehicle := NewVehicle(CAR, &GasEngine{})
	t.Log(vehicle.StartEngine())
}

func TestCarGasEngine_StopEngine(t *testing.T) {
	vehicle := NewVehicle(CAR, &GasEngine{})
	t.Log(vehicle.StopEngine())
}

func TestCarElectricEngine_StartEngine(t *testing.T) {
	vehicle := NewVehicle(CAR, &ElectricEngine{})
	t.Log(vehicle.StartEngine())
}

func TestCarElectricEngine_StopEngine(t *testing.T) {
	vehicle := NewVehicle(CAR, &ElectricEngine{})
	t.Log(vehicle.StopEngine())
}

func TestMotorcycleGasEngine_StartEngine(t *testing.T) {
	vehicle := NewVehicle(MOTORCYCLE, &GasEngine{})
	t.Log(vehicle.StartEngine())
}

func TestMotorcycleGasEngine_StopEngine(t *testing.T) {
	vehicle := NewVehicle(MOTORCYCLE, &GasEngine{})
	t.Log(vehicle.StopEngine())
}

func TestDrive(t *testing.T) {
	vehicle := NewVehicle(CAR, &GasEngine{})
	t.Log(Drive(vehicle))
}
