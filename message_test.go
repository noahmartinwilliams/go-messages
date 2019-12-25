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

func TestMessageJar(t *testing.T) {
	reqcc := make(chan chan Message)
	addc := make(chan Message)
	MessageJar(48 * time.Hour, addc, reqcc)
	reqc := make(chan Message)
	addc <- Message{Name:"noah", Contents:"hi"}
	reqcc <- reqc
	reply := <-reqc
	if reply.Name != "noah" {
		t.Errorf("Error: MessageJar did not return correct name in first test. Expected: \"noah\". Got: \"" + reply.Name + "\".")
	}

	reqcc2 := make(chan chan Message)
	addc2 := make(chan Message)
	MessageJar( 1 * time.Millisecond, addc2, reqcc2)
	reqc2 := make(chan Message)
	addc2 <- Message{Name:"noah", Contents:"hi"}
	time.Sleep(10 * time.Millisecond)
	reqcc2 <- reqc2
	reply2, ok := <-reqc2

	if ok {
		t.Errorf("Error: MessageJar returned {name:\"" + reply2.Name + "\", contents:\"" + reply2.Contents +"\"} during second test.")
	}
}
