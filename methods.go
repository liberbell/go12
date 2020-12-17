package main

import "math"

type Root struct {
	A, B float64
}

func (r *Root) Hyp() float64 {
	return math.Sqrt(r.A*r.A + r.B*r.B)
}
