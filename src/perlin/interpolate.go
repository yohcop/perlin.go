package perlin

import (
  "math"
)

func LinearInterpolation(a, b, t float32) float32 {
  return  a * (1-t) + b * t
}

func CosineInterpolation(a, b, t float32) float32 {
  return LinearInterpolation(a, b,
    (1 - float32(math.Cos(float64(t * math.Pi)))) * 0.5)
}

