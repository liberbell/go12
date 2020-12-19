package main

type Calcer interface{ Calc() float64 }
type Square struct{ X, Y float64 }
type myFloat float64

func (f myFloat) Calc() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
