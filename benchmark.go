package main

import (
	"math"
	"runtime"
	"sync"
	"time"
)

const OppaiFuncTStop = 32.0
const OppaiFuncDeltaT = 0.5
const MinimumDuration = 30

var thread_n int

func oppai_func(y float64, t float64) float64 {
	y = 0.02 * (y - 100)

	a1 := (1.5 * math.Exp((0.12*math.Sin(t)-0.5)*math.Pow((y+0.16*math.Sin(t)), 2))) / (1 + math.Exp(-20*(5*y+math.Sin(t))))
	a2 := ((1.5 + 0.8*math.Pow((y+0.2*math.Sin(t)), 3)) * math.Pow(1+math.Exp(20*(5*y+math.Sin(t))), -1)) / (1 + math.Exp(-(100*(y+1) + 16*math.Sin(t))))
	a3 := (0.2 * (math.Exp(-math.Pow(y+1, 2)) + 1)) / (1 + math.Exp(100*(y+1)+16*math.Sin(t)))
	a4 := 0.1 / math.Exp(2*math.Pow((10*y+1.2*(2+math.Sin(t))*math.Sin(t)), 4))

	return 65 * (a1 + a2 + a3 + a4)
}

func integral_f_p(alpha, beta float64, f func(float64) float64) float64 {
	wg := &sync.WaitGroup{}

	n := 1000000

	delta := (beta - alpha) / float64(n)
	data := make([]float64, n+1)

	for i := 0; i < thread_n; i++ {
		wg.Add(1)
		go func(i2 int) {
			temp_s := (beta - alpha) / float64(thread_n) * float64(i2)
			temp_e := (beta - alpha) / float64(thread_n) * float64(i2+1)
			j := 0
			for j_d := temp_s; j_d < temp_e; j_d += delta {
				f_a := f(j_d)
				f_b := f(j_d + delta)
				f_m := f((j_d + j_d + delta) / 2)

				data[j+i2*(n/thread_n)] = delta / 6 * (f_a + 4*f_m + f_b)
				j++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	a := 0.0
	for _, e := range data {
		a += e
	}
	return a
}

func dtime2score(v time.Duration) float32 {
	return 1 / (float32(v) / 1000000.0) * 1000000
}

func average(a []float32) float32 {
	m := float32(0.0)

	for i := 0; i < len(a); i++ {
		m += a[i]
	}

	m /= float32(len(a))

	return m
}

func benchmark(r chan ScoreReport, done chan struct{}) {
	thread_n = runtime.NumCPU()
	scores := []float32{}

	benchmark_begin := time.Now().Unix()
	for func_t := 0.0; time.Now().Unix() <= benchmark_begin+MinimumDuration || func_t < OppaiFuncTStop; func_t += OppaiFuncDeltaT {
		epoch_start_at := time.Now()

		area := integral_f_p(
			-1000,
			1000,
			func(v float64) float64 {
				return oppai_func(v, func_t)
			},
		)

		scores = append(scores, dtime2score(time.Now().Sub(epoch_start_at)))

		r <- ScoreReport{
			Area:  area,
			Score: average(scores),
		}
	}

	done <- struct{}{}
}
