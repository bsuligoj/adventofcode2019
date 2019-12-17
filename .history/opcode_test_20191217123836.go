package main

import "testing"

func TestOpcodeRun(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	wantedOutput := []int{2, 0, 0, 0, 99}
	opcode := NewOpcode(input)
	opcode.Run()
	output := opcode.Data

	if opcode.Data != wantedOutput {
		t.Errorf("for mass %v got %v; want %v", input, output, wantedOutput)
	}
}