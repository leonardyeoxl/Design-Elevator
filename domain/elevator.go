package domain

import (
	"container/heap"
	"fmt"
)

type IElevator interface {
	SendUpRequest(upRequest Request)
	SendDownRequest(downRequest Request)
	Run()
	processRequests()
	processUpRequests()
	processDownRequests()
	GetCurrentFloor() int
}

type Elevator struct {
	currentFloor int
	direction    Direction
	upQueue      heap.Interface
	downQueue    heap.Interface
}

func NewElevator(currentFloor int) Elevator {
	minHeap := make(MinHeap, 0)
	maxHeap := make(MaxHeap, 0)
	heap.Init(&minHeap)
	heap.Init(&maxHeap)
	return Elevator{
		currentFloor: currentFloor,
		direction:    IDLE,
		upQueue:      &minHeap,
		downQueue:    &maxHeap,
	}
}

func (e Elevator) SendUpRequest(request Request) {
	if request.location == OUTSIDE_ELEVATOR {
		heap.Push(e.upQueue, &Request{
			request.currentFloor,
			request.currentFloor,
			UP,
			OUTSIDE_ELEVATOR,
		})

		fmt.Printf("Append up request going to floor %d.\n", request.currentFloor)
	}

	heap.Push(e.upQueue, &request)

	fmt.Printf("Append up request going to floor %d.\n", request.desiredFloor)
}

func (e Elevator) SendDownRequest(request Request) {
	if request.location == OUTSIDE_ELEVATOR {
		heap.Push(e.downQueue, &Request{
			request.currentFloor,
			request.currentFloor,
			DOWN,
			OUTSIDE_ELEVATOR,
		})

		fmt.Printf("Append down request going to floor %d.\n", request.currentFloor)
	}

	heap.Push(e.downQueue, &request)

	fmt.Printf("Append down request going to floor %d.\n", request.desiredFloor)
}

func (e Elevator) Run() {
	for {
		if e.upQueue.Len() == 0 || e.downQueue.Len() == 0 {
			break
		}

		e.processRequests()
	}

	fmt.Println("Finished all requests.")
	e.direction = IDLE
}

func (e Elevator) processRequests() {
	if e.direction == UP || e.direction == IDLE {
		e.processUpRequests()
		e.processDownRequests()
	} else {
		e.processDownRequests()
		e.processUpRequests()
	}
}

func (e Elevator) processUpRequests() {
	for {
		if e.upQueue.Len() == 0 {
			break
		}

		upRequest := heap.Pop(e.upQueue).(*Request)
		e.currentFloor = upRequest.desiredFloor
		fmt.Printf("Processing up requests. Elevator stopped at floor %d.\n", e.currentFloor)
	}

	if e.downQueue.Len() > 0 {
		e.direction = DOWN
	} else {
		e.direction = IDLE
	}
}

func (e Elevator) processDownRequests() {
	for {
		if e.downQueue.Len() == 0 {
			break
		}

		downRequest := heap.Pop(e.downQueue).(*Request)
		e.currentFloor = downRequest.desiredFloor
		fmt.Printf("Processing down requests. Elevator stopped at floor %d.\n", e.currentFloor)
	}

	if e.upQueue.Len() > 0 {
		e.direction = UP
	} else {
		e.direction = IDLE
	}
}

func (e Elevator) GetCurrentFloor() int {
	return e.currentFloor
}
