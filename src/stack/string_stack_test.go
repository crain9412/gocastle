package stack

import "testing"

func Test_StringStack_Push(t *testing.T) {
	want := "[world] -> [hello]"
	stack := new(StringStack)
	stack.Push("hello")
	stack.Push("world")
	if got := stack.Print(); got != want {
		t.Errorf("Stack contained %q, want %q", got, want)
	}
}

func Test_StringStack_Pop(t *testing.T) {
	want := "world"
	stack := new(StringStack)
	stack.Push("hello")
	stack.Push("world")
	if got, ok := stack.Pop(); got != want {
		if !ok {
			t.Errorf("Stack didn't contain any elements")
		}
		t.Errorf("Stack's head was %q, want %q", got, want)
	}
}

func Test_StringStack_Pop_Too_Many(t *testing.T) {
	want := "∅"
	stack := new(StringStack)
	stack.Push("hello")
	stack.Push("world")
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	if got, ok := stack.Pop(); got != want {
		if !ok {
			t.Errorf("Stack didn't contain any elements")
		}
		t.Errorf("Stack's head was %q, want %q", got, want)
	}
}
