package cayleybench

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/cayley"
	"github.com/google/cayley/graph"
	_ "github.com/google/cayley/graph/bolt"
	_ "github.com/google/cayley/graph/mongo"
)

var (
	cm, cb *QuadWriter
)

func TestMain(m *testing.M) {
	var err error

	if err = graph.InitQuadStore("mongo", "mongodb://127.0.0.1:27017", nil); err != nil {
		panic(err)
	}

	mg, err := cayley.NewGraph("mongo", "mongodb://127.0.0.1:27017", nil)

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to MongoDB")
	}

	cm = NewQuadWriter(mg)

	f, err := ioutil.TempFile("", "cayleybench-bolt")
	if err != nil {
		panic("Error creating temporary file for BoltDB")
	}

	fn := f.Name()
	f.Close()
	defer os.Remove(fn)

	if err = graph.InitQuadStore("bolt", fn, nil); err != nil {
		panic(err)
	}

	bg, err := cayley.NewGraph("bolt", fn, nil)
	if err != nil {
		panic(err)
	}

	cb = NewQuadWriter(bg)
	m.Run()
}

func BenchmarkMongoInsert(b *testing.B) {
	cm.InitQuads(b.N * 2)
	b.ResetTimer()
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				cm.WriteQuad()
			}
		})
}

func BenchmarkBoltInserts(b *testing.B) {
	cb.InitQuads(b.N)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cb.WriteQuad()
		}
	})
}
