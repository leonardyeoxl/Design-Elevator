# Design-Elevator
System Design and Machine Coding an Elevator system for Tech Interviews

## Requirements
1. The elevator can go up and down in a real-world fashion.
2. Users can send requests to the elevator from both outside and inside the elevator.

### The elevator can go up and down in a real-world fashion
1. When elevator is going up or down, it will stop at all the floors that the users requested.
2. If the elevator received a request of going down while it is going up, the elevator will go to the highest floor in the current requests, and then go down.
3. Users can send requests at anytime.

## Analysis
- The elevator needs to sort the requests by some kind of order. 
- It’s not by timestamp, because if elevator is at floor 1, and customer A wants to go to floor 4, and B wants to go to floor 2, the elevator should not go to floor 4 first just because A sent the request first. 
- Instead, the elevator should stop at floor 2 and let B out, then go to floor 4 to let A out. 
- Thus, the request should be sorted by the distance from the current floor and not by timestamp.

## Assumptions
- going up has more priority than going down, which means that when the elevator is in IDLE state, and has both up and down requests, it will execute up requests first
- max heap to store all down requests and sort them by their desired floor
- min heap to store all up requests and sort them by their desired floor
- When, the requester is outside of the elevator, the elevator needs to stop at the `currentFloor` of the requester, before going to the `desiredFloor` of the requester.

## Run program
Command: `go run ./main/designelevator/.`

## Test
```
Append up request going to floor 5.
Append up request going to floor 3.
Append down request going to floor 4.
Append down request going to floor 0.
Append down request going to floor 1.
Append down request going to floor 2.
Processing up requests. Elevator stopped at floor 3.
Processing up requests. Elevator stopped at floor 5.
Processing down requests. Elevator stopped at floor 4.
Processing down requests. Elevator stopped at floor 2.
Processing down requests. Elevator stopped at floor 1.
Processing down requests. Elevator stopped at floor 0.
Finished all requests.
```

In the example, there are two people inside of the elevator that want to go to floor 5 and 3. A person outside the elevator at floor 4 wants to go to floor 0. And two people inside of the elevator want to go to floor 1 and 2.

We expect the elevator up first and stop at floor 3 and 5. Then, the elevator to go down and stop at floor 4 to pick up the person outside the elevator. Then, the elevator should keep going down and stop at floor 2, 1, and 0 respectively.

## Time and Space Complexity
- Time Complexity: O(nlogn). From the heap.
- Space Complexity: O(n). From the heap.