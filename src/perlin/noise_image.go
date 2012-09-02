package perlin

import (
  "image"
  "image/color"
)

type Noise1dImage struct {
  w, h int
  noise []float32
}

func NewNoise1dImage(w, h int, noise []float32) *Noise1dImage {
  return &Noise1dImage{ w: w, h: h, noise: noise }
}

func (n *Noise1dImage) ColorModel() color.Model {
  return color.RGBAModel
}

func (n *Noise1dImage) Bounds() image.Rectangle {
  return image.Rect(0, 0, n.w, n.h)
}

func (n *Noise1dImage) At(x, y int) color.Color {
  v := n.noise[x] * float32(n.h)
  yf := float32(y)
  if v >= yf && v < yf + 1 {
    return color.Gray{0}
  }
  return color.Gray{255}
}

// -------------------------------------------------

type Noise2dImage struct {
  w, h int
  noise [][]float32
}

func NewNoise2dImage(w, h int, noise [][]float32) *Noise2dImage {
  return &Noise2dImage{ w: w, h: h, noise: noise }
}

func (n *Noise2dImage) ColorModel() color.Model {
  return color.RGBAModel
}

func (n *Noise2dImage) Bounds() image.Rectangle {
  return image.Rect(0, 0, n.w, n.h)
}

func (n *Noise2dImage) At(x, y int) color.Color {
  return color.Gray{uint8(n.noise[x][y] * 255)}
}
