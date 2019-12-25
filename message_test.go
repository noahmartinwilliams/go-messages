package main

import "testing"
import "time"

func TestMsg(t *testing.T) {
	outputc := make(chan int)
	Msg(0, 10 * time.Millisecond, outputc)
	out := <-outputc
	if out != 0 {
		t.Errorf("Error: Msg did not send out 0 for first test.")
	}
}
