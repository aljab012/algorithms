package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}

	for i := range temperatures {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result[pop] = i - pop
		}
		stack = append(stack, i)
	}
	return result
}
