package main

func NumGen() chan int {
	retc := make(chan int, 100)
	num := 0
	go func() {
		for {
			retc <- num
			num = num + 1
		}
	} ()
	return retc
}
