package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	concurrent := 2
	done := make(chan bool, concurrent)

	for i := 1; i <= 4; i++ {
		done <- true
		go Calculate(2*time.Second, i, done)

	}

	for i := 0; i < concurrent; i++ {
		done <- true
	}

	fmt.Println("finish")
	fmt.Println(time.Since(start))
}

func Calculate(d time.Duration, num int, done <-chan bool) {
	defer func() {
		<-done
	}()
	start := time.Now()
	fmt.Println("job:", num, "start")
	time.Sleep(d)

	fmt.Println("job:", num, "done")
	fmt.Println(time.Since(start))
	fmt.Println()
}
