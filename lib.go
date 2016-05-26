package cayleybench

import (
	"sync"

	"github.com/google/cayley"
	"github.com/google/cayley/quad"
	"github.com/pborman/uuid"
)

// QuadWriter is a helper struct for the benchmarks
type QuadWriter struct {
	sync.Mutex
	cayley  *cayley.Handle
	quads   []quad.Quad
	counter int
}

// WriteQuad is called in the Benchmark* functions to
// perform the actual write operation
func (w *QuadWriter) WriteQuad() {
	var c int
	w.Lock()
	c = w.counter
	w.counter++
	w.Unlock()

	w.cayley.AddQuad(w.quads[c])
}

// InitQuads is used to pregenerate Quads
func (w *QuadWriter) InitQuads(nums int) {

	w.Lock()
	w.counter = 0
	w.quads = make([]quad.Quad, nums, nums)

	for i := 0; i < nums; i++ {
		w.quads[i] = cayley.Quad(uuid.NewRandom().String(), "follows", uuid.NewRandom().String(), "")
	}
	w.Unlock()
}

// NewQuadWriter creates a new Quadwriter.
// Each time WriteQuad is called, a Quad is written to the underlying handle
func NewQuadWriter(handle *cayley.Handle) *QuadWriter {
	return &QuadWriter{cayley: handle}
}
