package main

import "time"
type Message struct {
	Name string
	Contents string
}

func Msg(i int, t time.Duration, outputc chan int) {
	go func() {
		time.Sleep(t)
		outputc <- i
	} ()
}
