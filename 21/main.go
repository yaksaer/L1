package main

import "fmt"

type car struct {
}

func (c *car) driveCar(road roadType) {
	fmt.Println("Car is on its way")
	road.carRides()
}

type roadType interface {
	carRides()
}

type asphalt struct {
}

func (a *asphalt) carRides() {
	fmt.Println("Car is driving on asphalt")
}

type railway struct {
}

func (r *railway) carRides() {
	fmt.Println("Car is driving on railway")
}

type carAdapter struct {
	railway *railway
}

func (adap *carAdapter) carRides() {
	fmt.Println("Adapter puts the car on the railway platform")
	adap.railway.carRides()
}

func main() {
	car := &car{}
	asphalt := &asphalt{}
	car.driveCar(asphalt)

	rail := &railway{}
	carAdapter := &carAdapter{
		railway: rail,
	}
	car.driveCar(carAdapter)
}
