package cayleybench

import (
	"flag"
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/google/cayley"
	"github.com/google/cayley/graph"
	_ "github.com/google/cayley/graph/bolt"
	_ "github.com/google/cayley/graph/mongo"
	_ "github.com/google/cayley/graph/sql"
)

var (
	cm, cb *QuadWriter
	murl   = flag.String("mongo", "localhost:27017", "MongoDB url to connect to")
	pgurl  = flag.String("pg", "postgres://postgres:test@192.168.99.100/postgres?sslmode=disable", "PostGres url to connect to")
)

func TestMain(m *testing.M) {
	flag.Parse()

	m.Run()
}

func BenchmarkMemInsert(b *testing.B) {

	var (
		err  error
		memg *cayley.Handle
	)
	defer runtime.GC()

	if memg, err = cayley.NewMemoryGraph(); err != nil {
		b.Fatal("mem: Can not create MemoryGraph", err)
	}
	defer memg.Close()

	mw := NewQuadWriter(memg)
	mw.InitQuads(b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mw.WriteQuad()
	}
}

func BenchmarkPostgresInsert(b *testing.B) {

	var (
		err error
		pg  *cayley.Handle
	)

	defer runtime.GC()

	if err = graph.InitQuadStore("sql", *pgurl, nil); err != nil {
		b.Fatal("pgsql: Can not init QuadStore", err)
	}

	if pg, err = cayley.NewGraph("sql", *pgurl, nil); err != nil {
		b.Fatal("pgsql: Can not create QuadStore", err)
	}

	defer pg.Close()

	pw := NewQuadWriter(pg)
	pw.InitQuads(b.N)

	b.ResetTimer()

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				pw.WriteQuad()
			}
		})
	runtime.GC()
}

func BenchmarkMongoInsert(b *testing.B) {

	var (
		err error
		mg  *cayley.Handle
	)
	defer runtime.GC()

	if err = graph.InitQuadStore("mongo", *murl, nil); err != nil {
		b.Fatal("mongo: Can not init QuadStore", err)
	}

	mg, err = cayley.NewGraph("mongo", *murl, nil)
	if err != nil {
		b.Fatal("mongo: Can not create QuadStore", err)
	}
	defer mg.Close()

	cm := NewQuadWriter(mg)
	cm.InitQuads(b.N)

	b.ResetTimer()

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				cm.WriteQuad()
			}
		})

}

func BenchmarkBoltInserts(b *testing.B) {

	var (
		err error
		bg  *cayley.Handle
	)

	defer runtime.GC()

	f, err := ioutil.TempFile("", "cayleybench-bolt")
	if err != nil {
		b.Fatal("boltdb: Error creating temporary file for BoltDB", err)
	}

	fn := f.Name()
	f.Close()

	defer os.Remove(fn)

	if err = graph.InitQuadStore("bolt", fn, nil); err != nil {
		b.Fatal("boltdb: Can not init QuadStore", err)
	}

	if bg, err = cayley.NewGraph("bolt", fn, nil); err != nil {
		b.Fatal("boltdb: Can not create QuadStore", err)
	}
	defer bg.Close()

	cb = NewQuadWriter(bg)

	cb.InitQuads(b.N)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cb.WriteQuad()
		}
	})
}
