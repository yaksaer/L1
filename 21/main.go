package main

import "fmt"

// car struct for car, which can drive somewhere
type car struct {
}

func (c *car) driveCar(road roadType) {
	fmt.Println("Car is on its way")
	road.carRides()
}

// roadType interface which car use for driving
type roadType interface {
	carRides()
}

// asphalt type of road which car can drive on
type asphalt struct {
}

func (a *asphalt) carRides() {
	fmt.Println("Car is driving on asphalt")
}

// railway type of road which car can drive on, but only with adapter
type railway struct {
}

func (r *railway) carRides() {
	fmt.Println("Car is driving on railway")
}

// carAdapter struct which help car driving on railway
type carAdapter struct {
	railway *railway
}

func (adap *carAdapter) carRides() {
	fmt.Println("Adapter puts the car on the railway platform")
	adap.railway.carRides()
}

func main() {
	//create car and asphalts types
	car := &car{}
	asphalt := &asphalt{}
	//let car drives on asphalt
	car.driveCar(asphalt)
	//create railway and adapter types
	rail := &railway{}
	carAdapter := &carAdapter{
		railway: rail,
	}
	//let car drives on railway by adapter
	car.driveCar(carAdapter)
}
