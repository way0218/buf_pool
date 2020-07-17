package buf_pool

import (
	"testing"
)

func TestBufPool(t *testing.T) {
	_ = pool.New()
	b := Get()

	Put(b)
}
