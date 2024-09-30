package main

import "fmt"

type Vehicle struct {
	make  string
	model string
}

func (v Vehicle) Start() {
	fmt.Println("Vehicle is about to start.")
}

type Car struct {
	Vehicle
	seats int
}

func (c Car) Drive() {
	fmt.Printf("Driving a %s %s car with %d seats.\n", c.make, c.model, c.seats)
}

type BatteryPower interface {
	Charge()
}

type ElectricVehicle struct {
	Vehicle
	batteryCapacity int
}

func (ev ElectricVehicle) Charge() {
	fmt.Printf("Charging the electric vehicle with %d kWh battery.\n", ev.batteryCapacity)
}

func (ev ElectricVehicle) Drive() {
	fmt.Printf("Driving an electric %s %s with a battery capacity of %d kWh.\n", ev.make, ev.model, ev.batteryCapacity)
}