package perlin

import (
  "math"
  "fmt"
)

var _ = fmt.Println

func fromToPercent(c, x, w, pi int) (cfrom, from, cto, to int, percent float32) {
  chunk_per_interval := float32(pi) / float32(w)
  i_chunk_per_interval := 1
  if pi > w {
    i_chunk_per_interval = int(chunk_per_interval)
  }
  chunk_interval_start := c / i_chunk_per_interval * i_chunk_per_interval
  percent_chunk := float32(c % i_chunk_per_interval) / chunk_per_interval

  cell_start := x / pi * pi
  percent = float32(x - cell_start) / (float32(w) * chunk_per_interval) + percent_chunk

  cfrom = chunk_interval_start
  from = x / int(chunk_per_interval * float32(w)) * pi

  // TODO: can we compute 'blah' directly ?
  blah := 0
  if pi < w {
    blah = pi
  }
  cto = int(float32(chunk_interval_start) + chunk_per_interval) + (from + blah) / w
  to = (from + blah) % w

  //fmt.Println(
  //    "x", x,
  //    "pi", pi,
  //    "cell_start", cell_start,
  //    //"x-cell_start", x-cell_start,
  //    //"chunk_per_interval", chunk_per_interval,
  //    //"start_interval", start_interval,
  //    //"chunk_interval_start", chunk_interval_start,
  //    "percent_chunk", percent_chunk,
  //    "cfrom", cfrom,
  //    "from", from,
  //    "cto", cto,
  //    "to", to,
  //    )
  return
}

func noiseAt1d(w, cfrom, from, cto, to, pi int, percent float32, noise NoiseFunc1d) float32 {
  if percent == 0 {
    return noise(w, cfrom, from, pi)
  }
  n1 := noise(w, cfrom, from, pi)
  n2 := noise(w, cto, to, pi)
  return LinearInterpolation(n1, n2 , percent)
}

func noiseAt2d(w, h, cfromX, fromX, ctoX, toX, cfromY, fromY, ctoY, toY, pi int, percentX, percentY float32, noise NoiseFunc2d) float32 {
  var interpolationFromX float32 = 0
  if percentY == 0 {
    interpolationFromX = noise(w, h, cfromX, cfromY, fromX, fromY, pi)
  } else {
    interpolationFromX = LinearInterpolation(
        noise(w, h, cfromX, cfromY, fromX, fromY, pi),
        noise(w, h, fromX, ctoY, fromX, toY, pi),
        percentY)
  }
  if percentX == 0 {
    return interpolationFromX
  }

  var interpolationToX float32 = 0
  if (percentY == 0) {
    interpolationToX = noise(w, h, ctoX, cfromY, toX, fromY, pi)
  } else {
    interpolationToX = LinearInterpolation(
        noise(w, h, ctoX, cfromY, toX, fromY, pi),
        noise(w, h, ctoX, ctoY, toX, toY, pi),
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
      cfrom, from, cto, to, percent := fromToPercent(c, x, w, pi)
      fmt.Printf("from=%d to=%d percent=%f\n", from, to, percent)
      out[x] = (1-p) * out[x] + p * noiseAt1d(w, cfrom, from, cto, to, pi, percent, noise)
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
      cfromX, fromX, ctoX, toX, percentX := fromToPercent(cx, x, w, pi)
      for y := 0; y < h; y++ {
        cfromY, fromY, ctoY, toY, percentY := fromToPercent(cy, y, h, pi)
        out[x][y] = (1-p) * out[x][y] + p * noiseAt2d(
            w, h, cfromX, fromX, ctoX, toX, cfromY, fromY, ctoY, toY, pi,
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
      _, fromX, _, toX, percentX := fromToPercent(cx, x, w, pi)
      for y := 0; y < h; y++ {
        _, fromY, _, toY, percentY := fromToPercent(cy, y, h, pi)
        for z := 0; z < d; z++ {
          _, fromZ, _, toZ, percentZ := fromToPercent(cz, z, d, pi)
          out[x][y][z] = (1-p) * out[x][y][z] + p * noiseAt3d(
              w, h, d, cx, cy, cz, fromX, toX, fromY, toY, fromZ, toZ, pi,
              percentX, percentY, percentZ, noise)
        }
      }
    }
  }
  return out
}
