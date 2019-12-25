package main

import "testing"

func TestNumGen(t *testing.T) {
	numc := NumGen()
	num := <-numc
	if num != 0 {
		t.Errorf("Error: NumGen() did not send out 0 as first number.")
	}
}
