package perlin

import (
  "math"
  "fmt"
)

var _ = fmt.Println

func fromToPercent(x, pi int) (from, to int, percent float32) {
  from = x / pi
  percent = float32(x) / float32(pi) - float32(from)
  to = (from + 1) * pi
  from = from * pi
  return
}

func noiseAt1d(w, c, from, to, pi int, percent float32, noise NoiseFunc1d) float32 {
  if percent == 0 {
    return noise(w, c, from, pi)
  }
  n1 := noise(w, c, from, pi)
  n2 := noise(w, c, to, pi)
  return LinearInterpolation(n1, n2 , percent)
}

func noiseAt2d(w, h, cx, cy, fromX, toX, fromY, toY, pi int, percentX, percentY float32, noise NoiseFunc2d) float32 {
  var interpolationFromX float32 = 0
  if percentY == 0 {
    interpolationFromX = noise(w, h, cx, cy, fromX, fromY, pi)
  } else {
    interpolationFromX = LinearInterpolation(
        noise(w, h, cx, cy, fromX, fromY, pi),
        noise(w, h, cx, cy, fromX, toY, pi),
        percentY)
  }
  if percentX == 0 {
    return interpolationFromX
  }

  var interpolationToX float32 = 0
  if (percentY == 0) {
    interpolationToX = noise(w, h, cx, cy, toX, fromY, pi)
  } else {
    interpolationToX = LinearInterpolation(
        noise(w, h, cx, cy, toX, fromY, pi),
        noise(w, h, cx, cy, toX, toY, pi),
        percentY)
  }
  return LinearInterpolation(interpolationFromX, interpolationToX, percentX)
}

func noiseAt3d(w, h, d, cx, cy, cz, fromX, toX, fromY, toY, fromZ, toZ, pi int, percentX, percentY, percentZ float32, noise NoiseFunc3d) float32 {
  var interpolationFromXFromY float32 = 0
  if percentZ == 0 {
    interpolationFromXFromY = noise(w, h, d, cx, cy, cz, fromX, fromY, fromZ, pi)
  } else {
    interpolationFromXFromY = LinearInterpolation(
        noise(w, h, d, cx, cy, cz, fromX, fromY, fromZ, pi),
        noise(w, h, d, cx, cy, cz, fromX, fromY, toZ, pi),
        percentZ)
  }

  var interpolationFromX float32 = 0
  if percentY == 0 {
    interpolationFromX = interpolationFromXFromY
  } else {
    var interpolationFromXToY float32 = 0
    if percentZ == 0 {
      interpolationFromXFromY = noise(w, h, d, cx, cy, cz, fromX, toY, fromZ, pi)
    } else {
      interpolationFromXFromY = LinearInterpolation(
          noise(w, h, d, cx, cy, cz, fromX, toY, fromZ, pi),
          noise(w, h, d, cx, cy, cz, fromX, toY, toZ, pi),
          percentZ)
    }
    interpolationFromX = LinearInterpolation(
        interpolationFromXFromY, interpolationFromXToY, percentY)
  }

  if (percentX == 0) {
    return interpolationFromX
  }

  var interpolationToXFromY float32 = 0
  if percentZ == 0 {
    interpolationToXFromY = noise(w, h, d, cx, cy, cz, toX, fromY, fromZ, pi)
  } else {
    interpolationToXFromY = LinearInterpolation(
        noise(w, h, d, cx, cy, cz, toX, fromY, fromZ, pi),
        noise(w, h, d, cx, cy, cz, toX, fromY, toZ, pi),
        percentZ)
  }

  var interpolationToX float32 = 0
  if percentY == 0 {
    interpolationToX = interpolationToXFromY
  } else {
    var interpolationToXToY float32 = 0
    if percentZ == 0 {
      interpolationToXFromY = noise(w, h, d, cx, cy, cz, toX, toY, fromZ, pi)
    } else {
      interpolationToXFromY = LinearInterpolation(
          noise(w, h, d, cx, cy, cz, toX, toY, fromZ, pi),
          noise(w, h, d, cx, cy, cz, toX, toY, toZ, pi),
          percentZ)
    }
    interpolationToX = LinearInterpolation(
        interpolationToXFromY, interpolationToXToY, percentY)
  }
  return LinearInterpolation(interpolationFromX, interpolationToX, percentX)
}

func Noise1d(w, c int, persist float32, f, t int, noise NoiseFunc1d) []float32 {
  // Start clean and generate a random array.
  out := make([]float32, w)

  for i := f; i < t; i++ {
    pi := int(math.Pow(2, float64(i)))
    p := persist
    if i == f {
      // The first time, we take 100% of the noise.
      p = 1
    }
    for x := range out {
      from, to, percent := fromToPercent(x, pi)
      out[x] = (1-p) * out[x] + p * noiseAt1d(w, c, from, to, pi, percent, noise)
    }
  }
  return out
}

func Noise2d(w, h, cx, cy int, persist float32, f, t int, noise NoiseFunc2d) [][]float32 {
  // Generate a random array and prepare output array.
  out := make([][]float32, w)
  for x := 0; x < w; x++ {
    out[x] = make([]float32, h)
  }

  for i := f; i < t; i++ {
    pi := int(math.Pow(2, float64(i)))
    p := persist
    if i == f {
      p = 1
    }
    for x := 0; x < w; x++ {
      fromX, toX, percentX := fromToPercent(x, pi)
      for y := 0; y < h; y++ {
        fromY, toY, percentY := fromToPercent(y, pi)
        out[x][y] = (1-p) * out[x][y] + p * noiseAt2d(
            w, h, cx, cy, fromX, toX, fromY, toY, pi,
            percentX, percentY, noise)
      }
    }
  }
  return out
}

func Noise3d(w, h, d, cx, cy, cz int, persist float32, f, t int, noise NoiseFunc3d) [][][]float32 {
  // Generate a random array and prepare output array.
  out := make([][][]float32, w)
  for x := 0; x < w; x++ {
    out[x] = make([][]float32, h)
    for y := 0; y < h; y++ {
      out[x][y] = make([]float32, d)
    }
  }

  for i := f; i < t; i++ {
    pi := int(math.Pow(2, float64(i)))
    p := persist
    if i == f {
      p = 1
    }
    for x := 0; x < w; x++ {
      fromX, toX, percentX := fromToPercent(x, pi)
      for y := 0; y < h; y++ {
        fromY, toY, percentY := fromToPercent(y, pi)
        for z := 0; z < d; z++ {
          fromZ, toZ, percentZ := fromToPercent(z, pi)
          out[x][y][z] = (1-p) * out[x][y][z] + p * noiseAt3d(
              w, h, d, cx, cy, cz, fromX, toX, fromY, toY, fromZ, toZ, pi,
              percentX, percentY, percentZ, noise)
        }
      }
    }
  }
  return out
}
