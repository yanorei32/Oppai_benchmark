package main

import (
	"fmt"
	"time"
)

type chan_t struct {
	t     float64
	S     float64
	score float64
}

var benchmark_running bool
var chan_data chan chan_t

func main() {
	chan_data = make(chan chan_t, 4096)
	go benchmark()
	var temp_chan_data chan_t
	for benchmark_running {
	L2:
		for {
			select {
			case temp_chan_data = <-chan_data:

			default:
				break L2
			}
		}
		fmt.Printf("\r Score:%f Area:%f", temp_chan_data.score, temp_chan_data.S)
		time.Sleep(time.Millisecond * 250)
	}
}
