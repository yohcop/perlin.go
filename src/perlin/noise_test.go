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

  check(t, 4, 0, 0, 0, 0)
  check(t, 4, 0, 1, 0, 1)
  check(t, 4, 0, 2, 0, 2)
  check(t, 4, 0, 3, 0, 3)

  check(t, 4, 0, 4, 1, 0)
  check(t, 4, 0, 5, 1, 1)
  check(t, 4, 0, 6, 1, 2)
  check(t, 4, 0, 7, 1, 3)

  check(t, 4, 0, 8, 2, 0)
  check(t, 4, 0, 9, 2, 1)
}
