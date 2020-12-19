package main

type Calcer interface{ Calc() float64 }
type Square struct{ X, Y float64 }
type myFloat float64
