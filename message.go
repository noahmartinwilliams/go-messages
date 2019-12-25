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

func MessageJar(dur time.Duration, adder chan Message, requestc chan chan Message) {
	go func() {
	jar := make(map[int]Message)
	numgen := NumGen()
	delc := make(chan int)
	for {
		select {
			case toAdd := <-adder:
				num := <-numgen
				jar[num]=toAdd
				Msg(num, dur, delc)
			case del := <-delc:
				delete(jar, del)
			case req := <-requestc:
				for _, value := range jar {
					req <- value
				}
		}
	}
	} ()
}
