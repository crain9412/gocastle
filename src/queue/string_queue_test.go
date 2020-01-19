package queue

import "testing"

func Test_StringStack_Add(t *testing.T) {
	want := "[world] <-> [hello]"
	queue := new(StringQueue)
	queue.Add("hello")
	queue.Add("world")
	if got := queue.Print(); got != want {
		t.Errorf("Queue contained %q, want %q", got, want)
	}
}

func Test_StringStack_Poll(t *testing.T) {
	want := "hello"
	queue := new(StringQueue)
	queue.Add("hello")
	queue.Add("world")
	if got, ok := queue.Poll(); got != want {
		if !ok {
			t.Errorf("Queue didn't contain any elements")
		}

		t.Errorf("Queue first element was %q, want %q", got, want)
	}
}

func Test_StringStack_PollTooMany(t *testing.T) {
	want := "∅"
	queue := new(StringQueue)
	queue.Add("hello")
	queue.Add("world")
	queue.Poll()
	queue.Poll()
	queue.Poll()
	queue.Poll()
	if got, ok := queue.Poll(); got != want {
		if !ok {
			t.Errorf("Queue didn't contain any elements")
		}

		t.Errorf("Queue first element was %q, want %q", got, want)
	}
}
