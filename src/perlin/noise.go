package perlin

import (
  "fmt"
)

var _ = fmt.Println

type NoiseFunc1d func(w, c, x, f int) float32
type NoiseFunc2d func(w, h, cx, cy, x, y, f int) float32
type NoiseFunc3d func(w, h, d, cx, cy, cz, x, y, z, f int) float32

func combine(numbers ... int) int {
  res := 0
  for _, n := range numbers {
    res = res * 31 + n
  }
  return res
}

func randFrom(seed int) float32 {
  seed = (seed << 13) ^ seed
  return (float32(1.0 - float64(
             (seed*(seed*seed*15731+789221)+1376312589)&0x7fffffff) /
          float64(0x40000000)) + 1) / 2.0
}

func TileNoise(w, c, x, step int) float32 {
  //fmt.Printf("w=%d c=%d x=%d step=%d ---> c'=%d x'=%d\n", w, c, x, step, cp, xp)
  return randFrom(combine(w, c, x, step))
}

func TileNoise2d(w, h, cx, cy, x, y, step int) float32 {
  return randFrom(combine(w, h, cx, cy, x, y, step))
}

func TileNoise3d(w, h, d, cx, cy, cz, x, y, z, step int) float32 {
  return randFrom(combine(w, h, z, cx, cy, cz, x, y, z, step))
}
