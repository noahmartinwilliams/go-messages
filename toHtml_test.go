package main

import "testing"

func TestMsgEscape(t *testing.T) {
	inputc := make(chan Message)
	outc := MsgEscape(inputc)

	inputc <- Message{Name:"noah", Contents:"hello, world."}
	out1 := <-outc
	if out1.Name != "noah" {
		t.Errorf("Error: MsgEscape did not return correct Name. Expected: \"noah\". Got: \"" + out1.Name + "\".")
	}

	if out1.Contents != "hello, world." {
		t.Errorf("Error: MsgEscape did not return correct Contents. Expected: \"hello, world.\". Got: \"" + out1.Contents + "\".")
	}
}
