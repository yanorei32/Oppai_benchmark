package main

import (
	"fmt"
)

func main() {
	r := make(chan ScoreReport)
	s := make(chan struct{})
	go benchmark(r, s)

	for {
		select {
		case r := <- r:
			fmt.Printf("\rScore:%f Area:%f", r.Score, r.Area)
		case <-s:
			fmt.Printf("\n")
			return
		}
	}
}
