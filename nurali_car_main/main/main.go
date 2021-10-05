package main

import (
	"Go_Projects/nurali_car_main/car"
	"Go_Projects/nurali_car_main/headlights"
	"Go_Projects/nurali_car_main/stereo"
	"Go_Projects/nurali_car_main/wheels"
	"fmt"
)

func main() {
	fmt.Println(car.OpenDoor())
	fmt.Println(headlights.TurnOn())
	fmt.Println(stereo.TurnOn())
	fmt.Println(stereo.BoostBass())
	fmt.Println(wheels.Steer())
	fmt.Println(wheels.Accelerate())
}
