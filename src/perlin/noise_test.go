package perlin

import "testing"

func check(t *testing.T, w, c, x, expected_c, expected_x int) {
  cc, cx := canonical(w, c, x)
  if cc != expected_c || cx != expected_x {
    t.Errorf("Canonical(%d, %d, %d) = %d, %d. Want %d, %d",
        w, c, x, cc, cx, expected_c, expected_x)
  }
}

func TestCanonical(t *testing.T) {
  check(t, 1, 0, 0, 0, 0)
  check(t, 1, 0, 1, 1, 0)
  check(t, 1, 0, 2, 2, 0)
  check(t, 1, 0, 3, 3, 0)

  for c := 0; c < 5; c++ {
    check(t, 4, c, 0, c, 0)
    check(t, 4, c, 1, c, 1)
    check(t, 4, c, 2, c, 2)
    check(t, 4, c, 3, c, 3)
    check(t, 4, c, 4, c + 1, 0)
    check(t, 4, c, 5, c + 1, 1)
    check(t, 4, c, 6, c + 1, 2)
    check(t, 4, c, 7, c + 1, 3)
    check(t, 4, c, 8, c + 2, 0)
    check(t, 4, c, 9, c + 2, 1)
  }
}
