package goarea

const PI = 3.1416

import "math"

func Circ(raio float64) float64 {
  return math.Pow(raio, 2) * PI
}
