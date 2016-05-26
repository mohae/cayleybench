package cayleybench

import (
	"sync"

	"github.com/google/cayley"
	"github.com/google/cayley/quad"
	"github.com/pborman/uuid"
)

type QuadWriter struct {
	sync.Mutex
	cayley  *cayley.Handle
	quads   []quad.Quad
	counter int
}

func (w *QuadWriter) WriteQuad() {
	var c int
	w.Lock()
	c = w.counter
	w.counter++
	w.Unlock()

	w.cayley.AddQuad(w.quads[c])
}

func (w *QuadWriter) InitQuads(nums int) {

	w.Lock()
	w.counter = 0
	w.quads = make([]quad.Quad, nums, nums)

	for i := 0; i < nums; i++ {
		w.quads[i] = cayley.Quad(uuid.NewRandom().String(), "follows", uuid.NewRandom().String(), "")
	}
	w.Unlock()
}
func NewQuadWriter(handle *cayley.Handle) *QuadWriter {
	return &QuadWriter{cayley: handle}
}
