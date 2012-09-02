package main

import (
  "perlin"
  "fmt"
  "flag"
  "math/rand"
  "os"
  "image/png"
)

var persist = flag.Float64("p", 0.95, "Persistance")
var from = flag.Int("f", 1, "From this harmonic")
var to = flag.Int("t", 4, "To this harmonic")
var d = flag.Int("d", 1, "Dimensions")
var seed = flag.Int64("seed", 0, "Seed")

var x = flag.Int("x", 128, "Width")
var y = flag.Int("y", 128, "Height (only for 2d and 3d)")
var z = flag.Int("z", 128, "Depth (only for 3d)")

var cx = flag.Int("cx", 0, "Chunk X")
var cy = flag.Int("cy", 0, "Chunk Y (only for 2d and 3d)")
var cz = flag.Int("cz", 0, "Chunk Z (only for 3d)")

var dx = flag.Int("dx", 1, "Number of tiles to generate in X, starts at -cx.")
var dy = flag.Int("dy", 1, "Number of tiles to generate in y, starts at -cy.")
var dz = flag.Int("dz", 1, "Number of tiles to generate in z, starts at -cz.")

var img = flag.String("img", "", "Output image prefix. " +
                                 "Will be built with _x_y_z.png suffix. " +
                                 "Only for 1 and 2 d")

// To plot with gnuplot: pipe output to
// 1d:
//   gnuplot -p -e "plot '-' with lines"
// 2d:
//   gnuplot -p -e "set pm3d map;splot '-'"
// 3d:
//   gnuplot -p -e "splot '-' with points palette"
func main() {
  flag.Parse()
  rand.Seed(*seed)

  html := ""

  if *d == 1 {
    for ccx := *cx; ccx < *cx + *dx; ccx++ {
      out := perlin.Noise1d(*x, ccx, float32(*persist), *from, 1 + *to, perlin.TileNoise)
      if len(*img) > 0 {
        filename := fmt.Sprintf("%s_%d.png", *img, ccx)
        html += fmt.Sprintf("<img src=\"%s\">", filename)
        f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0600)
        defer f.Close()
        if err == nil {
          png.Encode(f, perlin.NewNoise1dImage(*x, *y, out))
        } else {
          fmt.Println(err.Error())
        }
      }
    }
  } else if *d == 2 {
    for ccy := *cy; ccy < *cy + *dy; ccy++ {
      for ccx := *cx; ccx < *cx + *dx; ccx++ {
        out := perlin.Noise2d(*x, *y, ccx, ccy, float32(*persist), *from, 1 + *to, perlin.TileNoise2d)
        if len(*img) > 0 {
          filename := fmt.Sprintf("%s_%d_%d.png", *img, ccx, ccy)
          html += fmt.Sprintf("<img src=\"%s\">", filename)
          f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0600)
          defer f.Close()
          if err == nil {
            png.Encode(f, perlin.NewNoise2dImage(*x, *y, out))
          } else {
            fmt.Println(err.Error())
          }
        }
      }
      html += "<br>"
    }
  } else if *d == 3 {
    out := perlin.Noise3d(*x, *y, *z, *cx, *cy, *cz, float32(*persist), *from, 1 + *to, perlin.TileNoise3d)
    for x := range out {
      for y := range out[x] {
        for z, v := range out[x][y] {
          if v > 0.4 {
            fmt.Printf("%d %d %d %f\n", x, y, z, v)
          }
        }
      }
    }
  }

  if len(*img) > 0 && len(html) > 0 {
    filename := fmt.Sprintf("%s.html", *img)
    f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0600)
    defer f.Close()
    if err == nil {
      f.Write([]byte(html))
    }
  }
}
