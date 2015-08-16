package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func single(iters int) float64 {
	var mypi float64 = 0
	ins := 0
	var x, y float64
	for count := 0; count < iters; count++ {
		x = rand.Float64()
		y = rand.Float64()
		if x*x+y*y < 1 {
			ins++
		}
	}
	mypi = (float64(ins) / float64(iters)) * 4.0
	return mypi
}

func multi(iters int) float64 {
	cpus := runtime.NumCPU()
	threadIters := iters / cpus
	results := make(chan float64, cpus)
	var mypi float64

	for j := 0; j < cpus; j++ {
		go func() {
			ins := 0
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for count := 0; count < threadIters; count++ {
				x, y := r.Float64(), r.Float64()
				if x*x+y*y < 1 {
					ins++
				}
			}
			results <- float64(ins)
		}()
	}
	for i := 0; i < cpus; i++ {
		mypi += <-results
	}
	return mypi * 4.0 / float64(cpus*threadIters)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%f\n", multi(1000000000))
}
