package cayleybench

import (
	"flag"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/google/cayley"
	"github.com/google/cayley/graph"
	_ "github.com/google/cayley/graph/bolt"
	_ "github.com/google/cayley/graph/mongo"
	_ "github.com/google/cayley/graph/sql"
)

var (
	murl  = flag.String("mongo", "mongodb:27017", "MongoDB url to connect to")
	pgurl = flag.String("pg", "postgres://postgres@postgres/postgres?sslmode=disable&connect_timeout=0", "PostGres url to connect to")
	sleep = flag.Int64("sleep", 0, "waits this number of seconds before executing the tests")
	once  sync.Once
)

func TestMain(m *testing.M) {
	flag.Parse()
	if *sleep > 0 {
		time.Sleep(time.Duration(*sleep) * time.Second)
	}
	m.Run()
}

func BenchmarkMemInsert(b *testing.B) {

	var (
		err  error
		memg *cayley.Handle
		memw *QuadWriter
	)

	if memg, err = cayley.NewMemoryGraph(); err != nil {
		b.Fatal("mem: Can not create MemoryGraph", err)
	}
	defer memg.Close()

	memw = NewQuadWriter(memg)
	memw.InitQuads(b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		memw.WriteQuad()
	}
}

func BenchmarkPostgresInsert(b *testing.B) {

	var (
		err error
		pg  *cayley.Handle
		pw  *QuadWriter
	)

	once.Do(
		func() {
			if err = graph.InitQuadStore("sql", *pgurl, nil); err != nil {
				b.Fatal("pgsql: Can not init QuadStore", err)
			}
		})

	if pg, err = cayley.NewGraph("sql", *pgurl, nil); err != nil {
		b.Fatal("pgsql: Can not create QuadStore", err)
	}

	defer pg.Close()

	pw = NewQuadWriter(pg)
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
		mw  *QuadWriter
	)

	if err = graph.InitQuadStore("mongo", *murl, nil); err != nil {
		b.Fatal("mongo: Can not init QuadStore", err)
	}

	mg, err = cayley.NewGraph("mongo", *murl, nil)
	if err != nil {
		b.Fatal("mongo: Can not create QuadStore", err)
	}
	defer mg.Close()

	mw = NewQuadWriter(mg)
	mw.InitQuads(b.N)

	b.ResetTimer()

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				mw.WriteQuad()
			}
		})

}

func BenchmarkBoltInserts(b *testing.B) {

	var (
		err error
		bg  *cayley.Handle
		bw  *QuadWriter
	)

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

	bw = NewQuadWriter(bg)

	bw.InitQuads(b.N)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bw.WriteQuad()
		}
	})
}
