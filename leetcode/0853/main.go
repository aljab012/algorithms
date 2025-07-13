package main

import "sort"

type Car struct {
	position    int
	arrivalTime float64
}

func carFleet(target int, position []int, speed []int) int {
	cars := []Car{}
	for i := 0; i < len(position); i++ {
		cars = append(cars, Car{
			position:    position[i],
			arrivalTime: float64(target-position[i]) / float64(speed[i]),
		})
	}

	// sort by position
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	// merge cars in a stack
	stack := []Car{}
	for _, car := range cars {
		if len(stack) == 0 || car.arrivalTime > stack[len(stack)-1].arrivalTime {
			stack = append(stack, car)
		}
	}
	return len(stack)
}
