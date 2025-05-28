package main

type MinStack struct {
	stack []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		mins:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.mins) == 0 || val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	pop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.mins) > 0 && pop == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		return 0
	}
	return this.mins[len(this.mins)-1]
}
