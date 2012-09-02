package perlin

import (
  "testing"
)

func verifyChunkIntervalStart(t *testing.T, c, w, step, expected int) {
  if v := chunkIntervalStart(c, w, step); v != expected {
    t.Errorf("chunkIntervalStart(%d, %d, %d) = %d. Want %d\n", c, w, step, v, expected)
  }
}

func TestChunkIntervalStart(t *testing.T) {
  //                          c  w step  expected
  verifyChunkIntervalStart(t, 0, 8,   1, 0)
  verifyChunkIntervalStart(t, 0, 8,   2, 0)
  verifyChunkIntervalStart(t, 0, 8,   4, 0)
  verifyChunkIntervalStart(t, 0, 8,   8, 0)

  verifyChunkIntervalStart(t, 1, 8,   1, 1)
  verifyChunkIntervalStart(t, 1, 8,   2, 1)
  verifyChunkIntervalStart(t, 1, 8,   4, 1)
  verifyChunkIntervalStart(t, 1, 8,   8, 1)
  verifyChunkIntervalStart(t, 1, 8,  16, 0)
  verifyChunkIntervalStart(t, 1, 8,  32, 0)

  verifyChunkIntervalStart(t, 2, 8,   1, 2)
  verifyChunkIntervalStart(t, 2, 8,   2, 2)
  verifyChunkIntervalStart(t, 2, 8,   4, 2)
  verifyChunkIntervalStart(t, 2, 8,   8, 2)
  verifyChunkIntervalStart(t, 2, 8,  16, 2)
  verifyChunkIntervalStart(t, 2, 8,  32, 0)
}

func TestChunkIntervalStartNegativeChunk(t *testing.T) {
  //                           c  w step  expected
  verifyChunkIntervalStart(t, -1, 8,   1, -1)
  verifyChunkIntervalStart(t, -1, 8,   2, -1)
  verifyChunkIntervalStart(t, -1, 8,   4, -1)
  verifyChunkIntervalStart(t, -1, 8,   8, -1)
  verifyChunkIntervalStart(t, -1, 8,  16, -2)
  verifyChunkIntervalStart(t, -1, 8,  32, -4)
  verifyChunkIntervalStart(t, -4, 8,  64, -8)

  verifyChunkIntervalStart(t, -2, 8,   1, -2)
  verifyChunkIntervalStart(t, -2, 8,   2, -2)
  verifyChunkIntervalStart(t, -2, 8,   4, -2)
  verifyChunkIntervalStart(t, -2, 8,   8, -2)
  verifyChunkIntervalStart(t, -2, 8,  16, -2)
  verifyChunkIntervalStart(t, -2, 8,  32, -4)
  verifyChunkIntervalStart(t, -4, 8,  64, -8)

  verifyChunkIntervalStart(t, -3, 8,   1, -3)
  verifyChunkIntervalStart(t, -3, 8,   2, -3)
  verifyChunkIntervalStart(t, -3, 8,   4, -3)
  verifyChunkIntervalStart(t, -3, 8,   8, -3)
  verifyChunkIntervalStart(t, -3, 8,  16, -4)
  verifyChunkIntervalStart(t, -3, 8,  32, -4)
  verifyChunkIntervalStart(t, -4, 8,  64, -8)

  verifyChunkIntervalStart(t, -4, 8,   1, -4)
  verifyChunkIntervalStart(t, -4, 8,   2, -4)
  verifyChunkIntervalStart(t, -4, 8,   4, -4)
  verifyChunkIntervalStart(t, -4, 8,   8, -4)
  verifyChunkIntervalStart(t, -4, 8,  16, -4)
  verifyChunkIntervalStart(t, -4, 8,  32, -4)
  verifyChunkIntervalStart(t, -4, 8,  64, -8)

  verifyChunkIntervalStart(t, -5, 8,   1, -5)
  verifyChunkIntervalStart(t, -5, 8,   2, -5)
  verifyChunkIntervalStart(t, -5, 8,   4, -5)
  verifyChunkIntervalStart(t, -5, 8,   8, -5)
  verifyChunkIntervalStart(t, -5, 8,  16, -6)
  verifyChunkIntervalStart(t, -5, 8,  32, -8)
  verifyChunkIntervalStart(t, -5, 8,  64, -8)

  verifyChunkIntervalStart(t, -9, 8,   1, -9)
  verifyChunkIntervalStart(t, -9, 8,   2, -9)
  verifyChunkIntervalStart(t, -9, 8,   4, -9)
  verifyChunkIntervalStart(t, -9, 8,   8, -9)
  verifyChunkIntervalStart(t, -9, 8,  16, -10)
  verifyChunkIntervalStart(t, -9, 8,  32, -12)
  verifyChunkIntervalStart(t, -9, 8,  64, -16)
}
