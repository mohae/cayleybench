// Copyright 2016 Markus W Mahlberg <markus.mahlberg@me.com>
//
// This file is part of cayleybench.
//
// cayleybench is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cayleybench is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with cayleybench.  If not, see <http://www.gnu.org/licenses/>.

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
