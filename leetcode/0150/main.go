package main

import "strconv"

func evalRPN1(tokens []string) int {

	funcMap := map[string]func(x, y int) int{
		"*": func(x, y int) int { return x * y },
		"/": func(x, y int) int { return x / y },
		"+": func(x, y int) int { return x + y },
		"-": func(x, y int) int { return x - y },
	}

	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			l := stack[len(stack)-2]
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, funcMap[token](l, r))
		}
	}
	return stack[0]
}

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if token == "+" {
			r := stack[len(stack)-1]
			l := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, l+r)
		} else if token == "-" {
			r := stack[len(stack)-1]
			l := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, l-r)
		} else if token == "*" {
			r := stack[len(stack)-1]
			l := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, l*r)
		} else if token == "/" {
			r := stack[len(stack)-1]
			l := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, l/r)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
