package main

import "math"

type hitung interface {
	luas() float64
	keliling() float64
}

type Persegi struct {
	sisi float64
}

type Lingkaran struct {
	diameter, jariJari float64
}

func (l Lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari, 2);
}

func (p Persegi) luas() float64 {
	return math.Pow(p.sisi, 2);
}

func (l Lingkaran) keliling() float64 {
	return math.Pi * l.diameter;
}

func (p Persegi) keliling() float64 {
	return p.sisi * 4;
}