package main

import (
	"testing"
)

// Tests:
// [ ] Jump moves to correct position
// [ ] All previous instructions increment after a jump

func TestJumpMovesToCorrectPosition(t *testing.T) {
	instructions := []int{0, 3, 0, 1, -3}
	tests := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 4},
		{2, 2},
		{3, 4},
		{4, 1},
	}

	for _, tt := range tests {
		_, actual := jump(instructions, tt.in) //execIndex + mov

		if tt.out != actual {
			t.Errorf("jump did not move to correct position")
		}
	}
}

func TestJumpIncrementsExecutedInstruction(t *testing.T) {
	instructions := []int{0, 3, 0, 1, -3}
	tests := []struct {
		in  int
		out int
	}{
		{0, 1},
		{1, 4},
		{2, 1},
		{3, 2},
		{4, -2},
	}

	for _, tt := range tests {
		newInstructions, _ := jump(instructions, tt.in)
		actual := newInstructions[tt.in]

		if tt.out != actual {
			t.Errorf("jump did not increment executed instruction")
		}
	}
}

func TestStrangeJumpIncrementsExecutedInstruction(t *testing.T) {
	instructions := []int{0, 3, 0, 1, -3}
	tests := []struct {
		in  int
		out int
	}{
		{0, 1},
		{1, 2},
		{2, 1},
		{3, 2},
		{4, -2},
	}

	for _, tt := range tests {
		newInstructions, _ := strangeJump(instructions, tt.in)
		actual := newInstructions[tt.in]

		if tt.out != actual {
			t.Errorf("jump did not increment executed instruction")
		}
	}
}
