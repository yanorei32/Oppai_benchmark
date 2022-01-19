package main

import (
	"math"
)

func Oppai_func(y float64, t float64) float64 {
	y = 0.02 * (y - 100)

	a1 := (1.5 * math.Exp((0.12*math.Sin(t)-0.5)*math.Pow((y+0.16*math.Sin(t)), 2))) / (1 + math.Exp(-20*(5*y+math.Sin(t))))
	a2 := ((1.5 + 0.8*math.Pow((y+0.2*math.Sin(t)), 3)) * math.Pow(1+math.Exp(20*(5*y+math.Sin(t))), -1)) / (1 + math.Exp(-(100*(y+1) + 16*math.Sin(t))))
	a3 := (0.2 * (math.Exp(-math.Pow(y+1, 2)) + 1)) / (1 + math.Exp(100*(y+1)+16*math.Sin(t)))
	a4 := 0.1 / math.Exp(2*math.Pow((10*y+1.2*(2+math.Sin(t))*math.Sin(t)), 4))

	return 65 * (a1 + a2 + a3 + a4)

}

type chan_t struct {
	t     float64
	S     float64
	score float64
}

var chan_data chan chan_t

var benchmark_running bool

type Game struct {
	x              []float64
	y              []float64
	temp_chan_data chan_t
}

