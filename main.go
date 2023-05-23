package main

import (
	"fmt"
	"syscall/js"
)

// VirtualMachine represents a simple virtual machine with registers.
type VirtualMachine struct {
	registers map[string]int
}

// NewVirtualMachine creates a new instance of VirtualMachine.
func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		registers: make(map[string]int),
	}
}

// SetRegister sets the value of a register.
func (vm *VirtualMachine) SetRegister(name string, value int) {
	vm.registers[name] = value
}

// GetRegister retrieves the value of a register.
func (vm *VirtualMachine) GetRegister(name string) int {
	return vm.registers[name]
}

// Expose the virtual machine methods to JavaScript.
func setupVM(vm *VirtualMachine) {
	js.Global().Set("setRegister", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			fmt.Println("setRegister expects 2 arguments")
			return nil
		}

		name := args[0].String()
		value := args[1].Int()

		vm.SetRegister(name, value)

		return nil
	}))

	js.Global().Set("getRegister", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			fmt.Println("getRegister expects 1 argument")
			return nil
		}

		name := args[0].String()
		value := vm.GetRegister(name)

		return js.ValueOf(value)
	}))
}

func main() {
	vm := NewVirtualMachine()

	// Expose the virtual machine to JavaScript.
	setupVM(vm)

	// Wait for JavaScript calls.
	select {}
}
