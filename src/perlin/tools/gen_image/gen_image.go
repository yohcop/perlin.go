package main

import (
  "perlin"
  "fmt"
  "flag"
  "math/rand"
  "os"
  "image/png"
)

var persist = flag.Float64("p", 0.5, "Persistance")
var from = flag.Int("f", 1, "From this harmonic")
var to = flag.Int("t", 4, "To this harmonic")
var at = flag.Int("at", -1, "If set, -f and -t are set to -at value")

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

  html := `
  <script>
  function ImgError(source) {
    var s = source.src;
    source.src = "";
    window.setTimeout(function() {
        source.src = s;
    }, 500 + Math.random(200) - 100);
    return true;
  }
  </script>
  `

  if *at != -1 {
    *from = *at
    *to = *at
  }
  if *d == 1 {
    for ccx := *cx; ccx < *cx + *dx; ccx++ {
      out := perlin.Noise1d(*x, ccx, float32(*persist), *from, 1 + *to, perlin.TileNoise)
      if len(*img) > 0 {
        filename := fmt.Sprintf("%s_%d.png", *img, ccx)
        html += fmt.Sprintf("<img src=\"%s\" onerror='ImgError(this)'>", filename)
        f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0600)
        if err == nil {
          png.Encode(f, perlin.NewNoise1dImage(*x, *y, out))
        } else {
          fmt.Println(err.Error())
        }
        f.Close()
      }
    }
  } else if *d == 2 {
    for ccy := *cy; ccy < *cy + *dy; ccy++ {
      for ccx := *cx; ccx < *cx + *dx; ccx++ {
        out := perlin.Noise2d(*x, *y, ccx, ccy, float32(*persist), *from, 1 + *to, perlin.TileNoise2d)
        if len(*img) > 0 {
          filename := fmt.Sprintf("%s_%d_%d.png", *img, ccx, ccy)
          html += fmt.Sprintf("<img src=\"%s\" onerror='ImgError(this)'>", filename)
          f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0600)
          if err == nil {
            png.Encode(f, perlin.NewNoise2dImage(*x, *y, out))
          } else {
            fmt.Println(err.Error())
          }
          f.Close()
        }
      }
      html += "<br>"
    }
  } else if *d == 3 {
    for ccy := *cy; ccy < *cy + *dy; ccy++ {
      for ccx := *cx; ccx < *cx + *dx; ccx++ {
        for ccz := *cz; ccz < *cz + *dz; ccz++ {
          out := perlin.Noise3d(*x, *y, *z, ccx, ccy, ccz,
                   float32(*persist), *from, 1 + *to, perlin.TileNoise3d)
          if len(*img) > 0 {
            for altitude := 0; altitude < *z; altitude++ {
              filename := fmt.Sprintf("%s_%d_%d_%d_%d.png", *img, ccx, ccy, ccz, altitude)
              html += fmt.Sprintf(
                  "<img class=\"a_%d_%d\" src=\"%s\" onerror='ImgError(this)'>",
                  ccz, altitude, filename)
              f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0600)
              if err == nil {
                png.Encode(f, perlin.NewNoise3dImage(*x, *y, *z, altitude, out))
              } else {
                fmt.Println(err.Error())
              }
              f.Close()
            }
          }
        }
      }
      html += "<br>"
    }
    html += fmt.Sprintf(`
    <div id="at"></div>
    <script>
    var startcz = %d;
    var dz = %d;
    var z = %d;

    var at_el = document.getElementById("at")
    var at_cz = startcz;
    var at_z = 0;
    window.setInterval(function() {
      for(var i = 0; i < document.images.length; ++i) {
        document.images[i].style.display = 'none';
      }
      at_z ++
      if (at_z >= z) {
        at_z = 0;
        at_cz++;
        if (at_cz >= startcz + dz) {
          at_cz = startcz;
        }
      }
      var cls = 'a_' + at_cz + '_' + at_z;
      for(var i = 0; i < document.images.length; ++i) {
        if (document.images[i].className === cls) {
          document.images[i].style.display = '';
        }
      }
      at_el.innerText = cls;
    }, 200)
    </script>
    `, *cz, *dz, *z)
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
