package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é: %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2) //(c.raio * c.raio)
}

func main() {
	r := retangulo{altura: 10, largura: 5}
	escreverArea(r)

	c := circulo{raio: 5}
	escreverArea(c)
}
