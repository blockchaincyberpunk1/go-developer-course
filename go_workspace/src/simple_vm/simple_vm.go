package main

import (
	"fmt"
	"os"
)

// Instruction set for the VM
const (
	OpPush  = iota // Push a value onto the stack
	OpAdd          // Pop the top two values from the stack, add them, and push the result
	OpSub          // Pop the top two values from the stack, subtract them, and push the result
	OpPrint        // Pop the top value from the stack and print it
	OpHalt         // Halt execution
)

type VM struct {
	stack []int
	code  []int
	ip    int // Instruction pointer
}

func NewVM() *VM {
	return &VM{
		stack: make([]int, 0),
		code:  []int{},
		ip:    0,
	}
}

func (vm *VM) Push(value int) {
	vm.stack = append(vm.stack, value)
}

func (vm *VM) Pop() int {
	if len(vm.stack) == 0 {
		fmt.Println("Error: Stack underflow")
		os.Exit(1)
	}
	index := len(vm.stack) - 1
	value := vm.stack[index]
	vm.stack = vm.stack[:index]
	return value
}

func (vm *VM) Run() {
	for vm.ip < len(vm.code) {
		instruction := vm.code[vm.ip]
		vm.ip++

		switch instruction {
		case OpPush:
			value := vm.code[vm.ip]
			vm.ip++
			vm.Push(value)
		case OpAdd:
			b := vm.Pop()
			a := vm.Pop()
			vm.Push(a + b)
		case OpSub:
			b := vm.Pop()
			a := vm.Pop()
			vm.Push(a - b)
		case OpPrint:
			fmt.Println(vm.Pop())
		case OpHalt:
			return
		}
	}
}

func main() {
	vm := NewVM()
	// Example program: 5 + 3 - 2
	vm.code = []int{OpPush, 5, OpPush, 3, OpAdd, OpPush, 2, OpSub, OpPrint, OpHalt}
	vm.Run()
}
