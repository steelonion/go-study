package main

import (
	"math/rand"
)

type RollPoint interface {
	GetPoint() float64
}

type FooRandom struct{}

func (c FooRandom) GetPoint() float64 {
	randomFloat := rand.Float64()
	return randomFloat
}

type FooRandom2 struct{}

func (c FooRandom2) GetPoint() float64 {
	return 1
}
