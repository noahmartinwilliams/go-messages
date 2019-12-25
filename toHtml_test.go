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

	inputc <- Message{Name:"noah", Contents:"<html>"}
	out2 := <-outc
	if out2.Contents != "&lt;html&gt;" {
		t.Errorf("Error: MsgEscape did not return correct contents. Expected: \"&lt;html&gt;\". Got: \"" + out2.Contents + "\"")
	}

	inputc <- Message{Name:"noah", Contents:"pb&j" }
	out3 := <-outc
	if out3.Contents != "pb&amp;j" {
		t.Errorf("Error: MsgEscape did not return correct contents. Expected: \"pb&amp;j\". Got: \"" + out3.Contents + "\".")
	}

}

func TestMsgToHTML(t *testing.T) {
	inputc := make(chan Message)
	outc := MsgToHTML(inputc)

	inputc <- Message{Name:"noah", Contents:"hello, world."}
	out1 := <-outc
	if out1 != "<div style=\"background-color:green\"><h3>noah</h3><p>hello, world.</p></div>" {
		t.Errorf("Error: MsgToHTML did not return correct output on first test. Got: \"" + out1 + "\"")
	}
}
