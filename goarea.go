package goarea

import "math"

const PI = 3.1416

func Circ(raio float64) float64 {
  return math.Pow(raio, 2) * PI
}
