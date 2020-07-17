package buf_pool

import (
	"bytes"
	"sync"
)

var pool sync.Pool

func init() {
	pool.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

// Get .
func Get() *bytes.Buffer {
	return pool.Get().(*bytes.Buffer)
}

// Put .
func Put(buf *bytes.Buffer) {
	buf.Truncate(0)
	pool.Put(buf)
}
