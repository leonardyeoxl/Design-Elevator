package main

import (
	domain "github.com/leonardyeoxl/Design-Elevator/domain"
)

func main() {
	elevator := domain.NewElevator(0)
	upRequest1 := domain.NewRequest(
		elevator.GetCurrentFloor(),
		5,
		domain.UP,
		domain.INSIDE_ELEVATOR,
	)
	upRequest2 := domain.NewRequest(
		elevator.GetCurrentFloor(),
		3,
		domain.UP,
		domain.INSIDE_ELEVATOR,
	)

	downRequest1 := domain.NewRequest(
		elevator.GetCurrentFloor(),
		1,
		domain.DOWN,
		domain.INSIDE_ELEVATOR,
	)
	downRequest2 := domain.NewRequest(
		elevator.GetCurrentFloor(),
		2,
		domain.DOWN,
		domain.INSIDE_ELEVATOR,
	)
	downRequest3 := domain.NewRequest(
		4,
		0,
		domain.DOWN,
		domain.OUTSIDE_ELEVATOR,
	)

	elevator.SendUpRequest(upRequest1)
	elevator.SendUpRequest(upRequest2)
	elevator.SendDownRequest(downRequest3)
	elevator.SendDownRequest(downRequest1)
	elevator.SendDownRequest(downRequest2)
	elevator.Run()
}
